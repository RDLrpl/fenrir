package fnlang

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type Proxies struct {
	Proxies []Proxy
}

type Proxy struct {
	Proto string
	Ip    string
	Port  string
	Tid   string
}

func LoadProxies(fncParams string) (Proxies, error) {
	var tg tgparams
	if err := json.Unmarshal([]byte(fncParams), &tg); err != nil {
		return Proxies{}, fmt.Errorf("[FENRIR] fnm!E!(Invalid TGParams)")
	}

	content, err := os.ReadFile(tg.ProxiesPath)
	if err != nil {
		return Proxies{}, fmt.Errorf("[FENRIR] fnm!E!(Failed to read proxies file: %v)", err)
	}

	var result Proxies
	re := regexp.MustCompile(`(?i)(SOCKS5)[\s\t]+([\d\.]+):(\d+)[\s\t-]*(\d+)?`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	if len(matches) == 0 {
		return result, fmt.Errorf("no valid proxies found. Content sample: %s", string(content))
	}

	for _, match := range matches {
		id := "0"
		if len(match) > 4 && match[4] != "" {
			id = match[4]
		}

		result.Proxies = append(result.Proxies, Proxy{
			Proto: match[1],
			Ip:    match[2],
			Port:  match[3],
			Tid:   id,
		})
	}

	return result, nil
}
