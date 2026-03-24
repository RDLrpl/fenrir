package fnlang

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Proxies struct {
	Proxies []Proxy
}

type Proxy struct {
	Proto     string
	Transport string
	Ip        string
	Port      string
	Login     string
	Pass      string
	Tid       string
}

func LoadProxies(fncParams string) (Proxies, error) {
	var tg tgparams
	if err := json.Unmarshal([]byte(fncParams), &tg); err != nil {
		return Proxies{}, err
	}

	content, err := os.ReadFile(tg.ProxiesPath)
	if err != nil {
		return Proxies{}, err
	}

	var result Proxies

	re := regexp.MustCompile(`(?i)(SOCKS5)[\s\t]+(?:([a-z0-9]+):)?([a-z0-9\-\.]+):(\d+)(?:\[(?:L:([^,\]]+),\s*P:([^\]]+)|NONE)\])?[\s\t-]*(\w+)?`)

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ">>") {
			continue
		}

		if strings.HasPrefix(line, "NONE") {
			parts := strings.Split(line, " - ")
			if len(parts) < 2 {
				continue
			}
			result.Proxies = append(result.Proxies, Proxy{
				Transport: "None",
				Tid:       strings.TrimSpace(parts[1]),
			})
			continue
		}

		match := re.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}

		result.Proxies = append(result.Proxies, Proxy{
			Proto:     match[1],
			Transport: match[2],
			Ip:        match[3],
			Port:      match[4],
			Login:     match[5],
			Pass:      match[6],
			Tid:       match[7],
		})
	}

	if len(result.Proxies) == 0 {
		return result, fmt.Errorf("NO VALID PROXIES (PROX <-)")
	}

	return result, nil
}
