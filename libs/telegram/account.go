package telegram

import (
	"encoding/json"
	"fmt"

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

func TG_PairAccounts(Params string) (Accounts, error) {
	var Accs Accounts

	msg, err := fnlang.MSG(Params)
	if err != nil {
		return Accounts{}, fmt.Errorf("[FENRIR] TG!E!(MESSAGE .fnm): %v", err)
	}

	tgApis, err := fnlang.TGfnk(Params)
	if err != nil {
		return Accounts{}, fmt.Errorf("[FENRIR] TG!E!(KEYS .fnk): %v", err)
	}
	tgProxies, err := fnlang.LoadProxies(Params)
	if err != nil {
		return Accounts{}, fmt.Errorf("[FENRIR] TG!E!(PROXY .prox): %v", err)
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
