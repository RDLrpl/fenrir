package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/RDLrpl/fenrir/back"
	"github.com/RDLrpl/fenrir/back/discord"
	"github.com/RDLrpl/fenrir/back/telegram"
	"github.com/RDLrpl/fenrir/ui"
	"github.com/RDLrpl/fenrir/utility"
)

func main() {
	fmt.Println(utility.FenArt)
	back.CreateOrCheckConf()

	arguments := os.Args
	args := len(arguments)
	if args >= 2 {
		if arguments[1] == "cli" && args >= 4 {
			switch arguments[2] {
			case "tg":
				telegram_manipulation(arguments[3])
			case "ds":
				discord_manipulation(arguments[3])
			case "fn":
				fenrir_manipulation(arguments[3])
			default:
				//usage
			}
		} else {
			// usage
		}
	} else {
		ui.RunApp()
	}
}

func telegram_manipulation(arg string) {
	switch {
	case strings.HasPrefix(arg, "auth-"):
		conf := back.ParseConfig()
		target := strings.TrimPrefix(arg, "auth-")

		if target == "1" {
			for id, acc := range conf.Telegram.Accounts {
				fmt.Printf("'%s' ?\n", id)
				err := telegram.CLIAuth(conf.Telegram.Sessions, id, acc)

				if err != nil {
					panic(err)
				}
				fmt.Printf("'%s' GOOD\n", id)
			}
		} else {
			acc, exists := conf.Telegram.Accounts[target]
			if !exists {
				fmt.Printf("Account '%s' does not exist\n", target)
				return
			}
			err := telegram.CLIAuth(conf.Telegram.Sessions, target, acc)
			if err != nil {
				panic(err)
			}
			fmt.Printf("GOOD\n")
		}
	case strings.HasPrefix(arg, "join-"):
		conf := back.ParseConfig()
		target := strings.TrimPrefix(arg, "join-")

		if target == "1" {
			for id, acc := range conf.Telegram.Accounts {
				fmt.Printf("'%s' ?\n", id)
				chann := conf.Telegram.Targets[acc.Targ]

				err := telegram.JoinTGChan(chann, conf.Telegram.Sessions, id, acc)
				if err != nil {
					panic(err)
				}
				fmt.Printf("'%s' GOOD\n", id)
			}
		} else {
			acc, exists := conf.Telegram.Accounts[target]
			if !exists {
				fmt.Printf("Account '%s' does not exist\n", target)
				return
			}
			chann := conf.Telegram.Targets[acc.Targ]

			err := telegram.JoinTGChan(chann, conf.Telegram.Sessions, target, acc)
			if err != nil {
				panic(err)
			}
			fmt.Printf("GOOD\n")
		}
	case strings.HasPrefix(arg, "send-"):
		conf := back.ParseConfig()
		target := strings.TrimPrefix(arg, "send-")

		if target == "1" {
			for id, acc := range conf.Telegram.Accounts {
				fmt.Printf("'%s' ?\n", id)
				chann := conf.Telegram.Targets[acc.Targ]
				message := conf.Messages[acc.Marg]

				err := telegram.SendTGmessage(chann, message, conf.Telegram.Sessions, id, acc)
				if err != nil {
					panic(err)
				}
			}
		} else {
			acc, exists := conf.Telegram.Accounts[target]
			if !exists {
				fmt.Printf("Account '%s' does not exist\n", target)
				return
			}
			chann := conf.Telegram.Targets[acc.Targ]
			message := conf.Messages[acc.Marg]

			err := telegram.SendTGmessage(chann, message, conf.Telegram.Sessions, target, acc)
			if err != nil {
				panic(err)
			}
		}
	case arg == "asend1":
		conf := back.ParseConfig()

		var wg sync.WaitGroup
		delayBetweenMessages := 200 * time.Millisecond

		for id, acc := range conf.Telegram.Accounts {
			wg.Add(1)
			go func(id string, acc utility.TelegramAccount) {
				defer wg.Done()

				client, err := telegram.RegOrLog(conf.Telegram.Sessions, id, acc)
				if err != nil {
					fmt.Printf("Ooops %s: %s\n", id, err)
					return
				}

				err = client.Run(context.Background(), func(ctx context.Context) error {
					ticker := time.NewTicker(delayBetweenMessages)
					defer ticker.Stop()

					for range ticker.C {
						chann := conf.Telegram.Targets[acc.Targ]
						message := conf.Messages[acc.Marg]

						err = telegram.ASendTGMessage(ctx, chann, message, id, client)
						if err != nil {
							fmt.Printf("RETRY %s: %s\n", id, err)
						}
					}
					return nil
				})

				if err != nil {
					fmt.Printf("Stopped [%s] FATAL: %s\n", id, err)
				}
			}(id, acc)
		}
		wg.Wait()
	default:
		// usage
	}
}

func discord_manipulation(arg string) {
	switch {
	case strings.HasPrefix(arg, "join-"):
		conf := back.ParseConfig()
		target := strings.TrimPrefix(arg, "join-")

		if target == "1" {
			for id, acc := range conf.Discord.Accounts {
				fmt.Printf("'%s' ?\n", id)

				targetapis := conf.Discord.Targets[acc.Targ]

				err := discord.Join(targetapis[1], acc.Token)
				if err != nil {
					panic(err)
				}
			}
		} else {
			acc, exists := conf.Discord.Accounts[target]
			if !exists {
				fmt.Printf("Account '%s' does not exist\n", target)
				return
			}

			targetapis := conf.Discord.Targets[acc.Targ]
			err := discord.Join(targetapis[1], acc.Token)
			if err != nil {
				panic(err)
			}
		}
	case strings.HasPrefix(arg, "send-"):
		conf := back.ParseConfig()
		target := strings.TrimPrefix(arg, "send-")

		if target == "1" {
			for id, acc := range conf.Discord.Accounts {
				fmt.Printf("'%s' ?\n", id)

				targetapis := conf.Discord.Targets[acc.Targ]
				message := conf.Messages[acc.Marg]

				err := discord.Send(targetapis[0], acc.Token, message, acc.Proxy, acc.ProxyLogin, acc.ProxyPass)
				if err != nil {
					panic(err)
				}
			}
		} else {
			acc, exists := conf.Discord.Accounts[target]
			if !exists {
				fmt.Printf("Account '%s' does not exist\n", target)
				return
			}
			targetapis := conf.Discord.Targets[acc.Targ]
			message := conf.Messages[acc.Marg]

			err := discord.Send(targetapis[0], acc.Token, message, acc.Proxy, acc.ProxyLogin, acc.ProxyPass)
			if err != nil {
				panic(err)
			}
		}
	case arg == "asend1":
		discord.Asendds(context.Background())
	default:
		//usage
	}

}

func fenrir_manipulation(arg string) {
	switch {
	case arg == "i-fenrirCAAU":
		err := back.AutoDownloadCAAU()
		if err != nil {
			panic(err)
		}
	}
}
