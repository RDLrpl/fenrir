package telegram

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/tg"
	"golang.org/x/net/proxy"
)

func Auth(acc Account) error {
	apiID, err := strconv.Atoi(acc.Api.API_id)
	if err != nil {
		return fmt.Errorf("[FENRIR] TG!E!(Invalid API_id): %v", err)
	}

	sessionStorage := &session.FileStorage{
		Path: fmt.Sprintf(".sessions/%s.json", acc.Api.ID),
	}

	var sauth *proxy.Auth
	if acc.Proxy.Login != "" {
		sauth = &proxy.Auth{
			User:     acc.Proxy.Login,
			Password: acc.Proxy.Pass,
		}
	}

	s := fmt.Sprintf("%s:%s", acc.Proxy.Ip, acc.Proxy.Port)

	sock5, err := proxy.SOCKS5(acc.Proxy.Transport, s, sauth, proxy.Direct)
	if err != nil {
		return fmt.Errorf("[FENRIR] TG!E!(PROXY ERROR): %v", err)
	}

	dc := sock5.(proxy.ContextDialer)
	client := telegram.NewClient(apiID, acc.Api.API_hash, telegram.Options{
		SessionStorage: sessionStorage,
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dc.DialContext,
		}),
	})

	return client.Run(context.Background(), func(ctx context.Context) error {
		var pass = ""

		if acc.Api.Pass != "NONE" {
			pass = acc.Api.Pass
		}

		flow := auth.NewFlow(
			auth.Constant(acc.Api.Number, pass, auth.CodeAuthenticatorFunc(
				func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
					fmt.Printf("[FENRIR] Req!Code (%s) >> ", acc.Api.Number)
					var code string
					fmt.Scan(&code)
					return code, nil
				},
			)),
			auth.SendCodeOptions{},
		)
		return client.Auth().IfNecessary(ctx, flow)
	})
}
