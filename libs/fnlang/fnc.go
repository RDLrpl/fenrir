package fnlang

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type fnc struct {
	Target string
	Params string
}

type tgparams struct {
	ApisPath    string `json:"apis"`
	ProxiesPath string `json:"proxies"`
	MessagePath string `json:"message"`
	Flags       string `json:"flags"`
}

func ReadConfiguration() (fnc, error) {
	var fnconf fnc
	var tgparams tgparams

	conf, err := os.ReadFile("conf.fnc")
	if err != nil {
		return fnc{}, err
	}

	targetRegex := regexp.MustCompile(`TARGET:(\w+)`)
	if match := targetRegex.FindStringSubmatch(string(conf)); len(match) > 1 {
		if match[1] == "telegram" {
			fnconf.Target = match[1]
		} else {
			return fnc{}, fmt.Errorf("[FENRIR] Unsupported Configuration!")
		}
	}

	if fnconf.Target == "telegram" {
		scanner := bufio.NewScanner(strings.NewReader(string(conf)))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			if strings.HasPrefix(line, ">>") || line == "" || strings.Contains(line, "{") || strings.Contains(line, "}") {
				continue
			}

			parts := strings.SplitN(line, "=", 2)
			if len(parts) < 2 {
				return fnc{}, fmt.Errorf(fmt.Sprintf("[FENRIR] fnc!E!(<2=) line: %s", line))
			}

			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])

			switch key {
			case "APIS":
				tgparams.ApisPath = val
			case "MESSAGE":
				tgparams.MessagePath = val
			case "PROXIES":
				tgparams.ProxiesPath = val
			case "FLAGS":
				tgparams.Flags = val

			default:
				return fnc{}, fmt.Errorf(fmt.Sprintf("[FENRIR] fnc!E!(?nil Key) line: %s", line))
			}
		}
		if tgparams.ApisPath == "" {
			return fnc{}, fmt.Errorf("[FENRIR] !E!(?NO FNK)")
		}
		if tgparams.MessagePath == "" {
			return fnc{}, fmt.Errorf("[FENRIR] !E!(?NO FNM)")
		}
		if tgparams.ProxiesPath == "" {
			return fnc{}, fmt.Errorf("[FENRIR] !E!(?NO PROX)")
		}
		if tgparams.Flags == "" {
			tgparams.Flags = "NONE"
		}

		jsonData, _ := json.Marshal(tgparams)
		fnconf.Params = string(jsonData)

		return fnconf, nil
	}

	return fnc{}, nil
}
