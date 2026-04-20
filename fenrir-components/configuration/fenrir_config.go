package configuration

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/RDLrpl/fenrir/fenrir-components/cli"
)

type Config struct {
	Instruction string

	ApisPath    string
	ProxiesPath string
	MessagePath string
	// todo
}

func ReadConfiguration() Config {
	var target string

	conf, err := os.ReadFile("fenrir.conf")
	if err != nil {
		panic(cli.Error_Style.Render(strings.TrimSpace(err.Error())))
	}

	targetRegex := regexp.MustCompile(`TARGET:(\w+)`)
	if match := targetRegex.FindStringSubmatch(string(conf)); len(match) > 1 {
		if match[1] == "telegram" {
			target = match[1]
		} else {
			panic(cli.Error_Style.Render("Unsupported Target Configuration"))
		}
	}

	if target == "telegram" {
		apis, msg, prox := telegram_target(string(conf))

		return Config{
			ApisPath:    apis,
			ProxiesPath: prox,
			MessagePath: msg,
			Instruction: target,
		}
	}

	panic(cli.Error_Style.Render("Unsupported Target Configuration"))
}

func telegram_target(fl string) (string, string, string) {
	var apis, msg, prox string

	scanner := bufio.NewScanner(strings.NewReader(fl))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, ">>") || line == "" || strings.Contains(line, "{") || strings.Contains(line, "}") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			panic(cli.Error_Style.Render("Invalid Configuration"))
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "APIS":
			apis = val
		case "MESSAGE":
			msg = val
		case "PROXIES":
			prox = val

		default:
			panic(cli.Error_Style.Render("Invalid Configuration"))
		}
	}

	return apis, msg, prox
}
