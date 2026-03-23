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
	if len(os.Args) == 3 {
		if os.Args[1] == "TG" && os.Args[2] == "--start!m" {
			handlers.TG_Send()
		}
		if os.Args[1] == "TG" && os.Args[2] == "--auth" {
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
			fenrir TG --start!m (Start message sending)

			fenrir --clean (clean all cache files (e.x. .sessions))
			`,
		))
	/*
		conf, err := fnlang.ReadConfiguration()
		if err != nil {
			panic(err)
		}

		Acc, err := telegram.TG_PairAccounts(conf.Params)
		if err != nil {
			panic(err)
		}

		var wg sync.WaitGroup
		delayBetweenMessages := 600 * time.Millisecond

		for _, Account := range Acc.Accs {
			wg.Add(1)
			go func(acc telegram.Account) {
				defer wg.Done()

				ticker := time.NewTicker(delayBetweenMessages)
				defer ticker.Stop()

				for {
					select {
					case <-ticker.C:
						fmt.Printf("[FENRIR]: Sending From >> %v\n", acc.Id)
						err := telegram.SendTGmessage(acc)
						if err != nil {
							fmt.Printf("[ERROR]!%v: %v\n", acc.Id, err)
							return
						}
					}
				}
			}(Account)
		}

		wg.Wait()
	*/
}
