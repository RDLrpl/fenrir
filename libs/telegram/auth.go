package telegram

import (
	"context"
	"fmt"

	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

func Auth(acc Account) error {
	client, err := GenerateTGClient(acc)
	if err != nil {
		return err
	}

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
