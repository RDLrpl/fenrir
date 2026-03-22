package telegram

import (
	"encoding/json"

	"github.com/RDLrpl/Fenrir/libs/fnlang"
)

type Accounts struct {
	Accs []Account
}

type Account struct {
	Id    string
	Api   fnlang.TGApi
	Msg   fnlang.Message
	Proxy fnlang.Proxy
}

func PairAccounts(Params string) (Accounts, error) {
	var Accs Accounts

	msg, err := fnlang.MSG(Params)
	if err != nil {
		panic(err)
	}

	tgApis, err := fnlang.TGfnk(Params)
	if err != nil {
		return Accounts{}, err
	}
	tgProxies, err := fnlang.LoadProxies(Params)
	if err != nil {
		return Accounts{}, err
	}

	for _, api := range tgApis.Apis {
		var Acc Account
		var tgapi fnlang.TGApi

		err := json.Unmarshal([]byte(api), &tgapi)
		if err != nil {
			return Accounts{}, err
		}

		taegetTid := tgapi.ID

		for _, proxy := range tgProxies.Proxies {

			if proxy.Tid == taegetTid {
				Acc.Proxy = proxy
			}
		}

		Acc.Id = tgapi.ID
		Acc.Api = tgapi
		Acc.Msg = msg
		Accs.Accs = append(Accs.Accs, Acc)
	}

	return Accs, nil
}
