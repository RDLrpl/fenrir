package telegram

import (
	"fmt"
	"strconv"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/dcs"
	"golang.org/x/net/proxy"
)

func GenerateTGClient(acc Account) (*telegram.Client, error) {
	apiID, err := strconv.Atoi(acc.Api.API_id)
	if err != nil {
		return nil, fmt.Errorf("[FENRIR] TG!E!(Invalid API_id): %v", err)
	}

	opts := telegram.Options{
		SessionStorage: &session.FileStorage{
			Path: fmt.Sprintf(".sessions/%s.json", acc.Api.ID),
		},
	}

	if acc.Proxy.Transport != "None" {
		var sauth *proxy.Auth
		if acc.Proxy.Login != "" {
			sauth = &proxy.Auth{
				User:     acc.Proxy.Login,
				Password: acc.Proxy.Pass,
			}
		}

		addr := fmt.Sprintf("%s:%s", acc.Proxy.Ip, acc.Proxy.Port)
		sock5, err := proxy.SOCKS5(acc.Proxy.Transport, addr, sauth, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("[FENRIR] TG!E!(PROXY ERROR): %v", err)
		}

		if dc, ok := sock5.(proxy.ContextDialer); ok {
			opts.Resolver = dcs.Plain(dcs.PlainOptions{
				Dial: dc.DialContext,
			})
		}
	}

	return telegram.NewClient(apiID, acc.Api.API_hash, opts), nil
}
