package msg

import (
	"os"
	"regexp"
	"strings"

	"github.com/RDLrpl/fenrir/fenrir-components/cli"
)

type Message struct {
	Target string `json:"target"`
	Data   string `json:"data"`
}

func ReadMsg(message_path string) Message {
	var msg Message

	content, err := os.ReadFile(message_path)
	if err != nil {
		panic(cli.Error_Style.Render(strings.TrimSpace(err.Error())))
	}

	reID := regexp.MustCompile(`CH:([^\s\n]+)`)
	idMatch := reID.FindStringSubmatch(string(content))
	if len(idMatch) > 1 {
		msg.Target = idMatch[1]
	}

	reMsg := regexp.MustCompile(`(?s)#START\s+(.*?)\s+#END`)
	msgMatch := reMsg.FindStringSubmatch(string(content))
	if len(msgMatch) > 1 {
		msg.Target = strings.TrimSpace(msgMatch[1])
	}

	return msg
}
