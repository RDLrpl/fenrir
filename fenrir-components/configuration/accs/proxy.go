package accs

import (
	"os"
	"regexp"
	"strings"

	"github.com/RDLrpl/fenrir/fenrir-components/cli"
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

func LoadProxies(proxies_path string) Proxies {
	content, err := os.ReadFile(proxies_path)
	if err != nil {
		panic(cli.Error_Style.Render(strings.TrimSpace(err.Error())))
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
		panic(cli.Error_Style.Render("Use NONE - 0 and :0 prox in .fa"))
	}

	return result
}
