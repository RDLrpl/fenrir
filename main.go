package main

import (
	"fmt"
	"os"

	"github.com/RDLrpl/Fenrir/libs/handlers"
	"github.com/RDLrpl/Fenrir/libs/utils"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#812e23"))

	fmt.Println(style.Render(utils.FenArt))
	if len(os.Args) == 2 {
		if os.Args[1] == "--clean" {
			handlers.Clean()
		}
		return
	}
	if len(os.Args) == 3 && os.Args[1] == "TG" {
		if os.Args[2] == "--ATCK!m" {
			handlers.TG_Send()
		}
		if os.Args[2] == "--ATCK!j" {
			handlers.TG_Join()

			return
		}

		if os.Args[2] == "--auth" {
			handlers.Telegram_Auth()

			return
		}

	}

	helptext := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00b7ff"))

	fmt.Println(
		helptext.Render(
			`
			usage:

			fenrir TG --auth (Auth Accs {Creating .sessions for accounts})
			fenrir TG --ATCK!m (Start message sending)
			fenrir TG --ATCK!j (Join chat (message.fnm) // only chats, simple join)

			fenrir --clean (clean all cache files (e.x. .sessions))
			`,
		))
}
