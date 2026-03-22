package main

import (
	"fmt"
	"os"

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

	}
	conf, err := fnlang.ReadConfiguration()
	if err != nil {
		panic(err)
	}

	Acc, err := telegram.PairAccounts(conf.Params)

	for _, Account := range Acc.Accs {
		n := 1
		for n > 0 {
			fmt.Println(Account)
			err := telegram.SendTGmessage(Account)
			if err != nil {
				panic(err)
			}
		}
	}
}
