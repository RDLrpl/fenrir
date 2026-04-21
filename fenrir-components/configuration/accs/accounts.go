package accs

import (
	"bufio"
	"os"
	"strings"

	"github.com/RDLrpl/fenrir/fenrir-components/cli"
)

type Accounts struct {
	TGApis []TGApi
}

type TGApi struct {
	ID       string
	API_id   string
	API_hash string
	Number   string
	Pass     string
	Prox     string
}

func TGfa(Acc_path string) Accounts {
	apis, err := os.Open(Acc_path)
	if err != nil {
		panic(cli.Error_Style.Render(strings.TrimSpace(err.Error())))
	}
	defer apis.Close()

	var resultApis Accounts
	scanner := bufio.NewScanner(apis)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, ">>") || line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) == 7 && parts[0] == "TG" {
			apiData := TGApi{
				ID:       parts[1],
				API_id:   parts[2],
				API_hash: parts[3],
				Number:   parts[4],
				Pass:     parts[5],
				Prox:     parts[6],
			}

			resultApis.TGApis = append(resultApis.TGApis, apiData)
		} else {
			panic(cli.Error_Style.Render("Bad Format. Check <<fa>> usage"))
		}
	}

	return resultApis
}
