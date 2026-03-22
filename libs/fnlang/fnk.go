package fnlang

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Fnk struct {
	ApisType string
	Apis     []string
}

type TGApi struct {
	ID       string `json:"id"`
	API_id   string `json:"api_id"`
	API_hash string `json:"api_hash"`
	Number   string `json:"number"`
	Pass     string `json:"pass"`
}

func TGfnk(fncParams string) (Fnk, error) {
	var tg tgparams

	err := json.Unmarshal([]byte(fncParams), &tg)
	if err != nil {
		return Fnk{}, fmt.Errorf("[FENRIR] fnk!E!(Invalid TGParams)")
	}

	apis, err := os.Open(tg.ApisPath)
	if err != nil {
		return Fnk{}, fmt.Errorf("[FENRIR] fnk!E!(Invalid fnk-Path)")
	}
	defer apis.Close()

	var resultApis []string
	scanner := bufio.NewScanner(apis)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, ">>") || line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) == 6 && parts[0] == "TG" {
			apiData := TGApi{
				ID:       parts[1],
				API_id:   parts[2],
				API_hash: parts[3],
				Number:   parts[4],
				Pass:     parts[5],
			}

			jsonData, _ := json.Marshal(apiData)
			resultApis = append(resultApis, string(jsonData))
		} else {
			return Fnk{}, fmt.Errorf(fmt.Sprintf("[FENRIR] fnk!E!(Invalid Format) line: %s", line))
		}
	}

	return Fnk{
		ApisType: "TG",
		Apis:     resultApis,
	}, nil
}
