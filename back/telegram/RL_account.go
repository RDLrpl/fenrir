package telegram

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/RDLrpl/fenrir/utility"
	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/tg"
	"golang.org/x/net/proxy"
)

func CLIAuth(storage string, accountid string, account utility.TelegramAccount) error {
	client, err := RegOrLog(storage, accountid, account)
	if err != nil {
		return err
	}

	return client.Run(context.Background(), func(ctx context.Context) error {
		flow := auth.NewFlow(
			auth.Constant(account.Number, account.CloudPass, auth.CodeAuthenticatorFunc(
				func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
					fmt.Printf("[FENRIR] Req!Code (%s) >> ", account.Number)
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

func RegOrLog(storage string, accountid string, account utility.TelegramAccount) (*telegram.Client, error) {
	APIID, err := strconv.Atoi(account.API_id)
	if err != nil {
		return nil, fmt.Errorf("TG!E! Invalid API_id: %v", err)
	}

	if err := os.MkdirAll(storage, 0755); err != nil {
		return nil, fmt.Errorf("TG!E! Failed session directory: %v", err)
	}

	opts := telegram.Options{
		SessionStorage: &session.FileStorage{
			Path: fmt.Sprintf("%s/%s.json", storage, accountid),
		},
	}

	host, log, pass, good := utility.Check_proxy(account.Proxy, account.ProxyLogin, account.ProxyPass)

	if good {
		var sauth *proxy.Auth
		if log != "" {
			sauth = &proxy.Auth{
				User:     log,
				Password: pass,
			}
		}

		sock5, err := proxy.SOCKS5("tcp", host, sauth, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("TG!E! PROXY ERROR: %v", err)
		}

		if dc, ok := sock5.(proxy.ContextDialer); ok {
			opts.Resolver = dcs.Plain(dcs.PlainOptions{
				Dial: dc.DialContext,
			})
		}
	}

	return telegram.NewClient(APIID, account.API_hash, opts), nil
}
