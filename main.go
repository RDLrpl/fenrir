package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/RDLrpl/Fenrir/libs/fnlang"
	"github.com/RDLrpl/Fenrir/libs/telegram"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--auth" {
		err := os.MkdirAll(".sessions", 0755)
		if err != nil {
			panic(err)
		}
		conf, err := fnlang.ReadConfiguration()
		if err != nil {
			panic(err)
		}

		Acc, err := telegram.PairAccounts(conf.Params)
		if err != nil {
			panic(err)
		}

		for _, Account := range Acc.Accs {
			if err := telegram.Auth(Account); err != nil {
				panic(err)
			}
			fmt.Println("Login successful")
		}
		return
	}

	conf, err := fnlang.ReadConfiguration()
	if err != nil {
		panic(err)
	}

	Acc, err := telegram.PairAccounts(conf.Params)
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
					fmt.Printf("Sending message from account: %v\n", acc)
					err := telegram.SendTGmessage(acc)
					if err != nil {
						fmt.Printf("Error sending message for account %v: %v\n", acc, err)
						return
					}
				}
			}
		}(Account)
	}

	wg.Wait()

}
