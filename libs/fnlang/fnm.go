package fnlang

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Message struct {
	Channel_id string
	Msg        string
}

func MSG(fncParams string) (Message, error) {
	var tg tgparams

	err := json.Unmarshal([]byte(fncParams), &tg)
	if err != nil {
		return Message{}, fmt.Errorf("[FENRIR] fnm!E!(Invalid TGParams)")
	}

	content, err := os.ReadFile(tg.MessagePath)
	if err != nil {
		return Message{}, fmt.Errorf("[FENRIR] fnm!E!(Invalid fnm-Path)")
	}

	res := Message{}

	reID := regexp.MustCompile(`CH:([^\s\n]+)`)
	idMatch := reID.FindStringSubmatch(string(content))
	if len(idMatch) > 1 {
		res.Channel_id = idMatch[1]
	}

	reMsg := regexp.MustCompile(`(?s)#START\s+(.*?)\s+#END`)
	msgMatch := reMsg.FindStringSubmatch(string(content))
	if len(msgMatch) > 1 {
		res.Msg = strings.TrimSpace(msgMatch[1])
	}

	return res, nil
}
