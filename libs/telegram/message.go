package telegram

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	tgmessage "github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/peers"
	"github.com/gotd/td/tg"
	"golang.org/x/net/proxy"
)

func SendTGmessage(acc Account) error {
	apiID, err := strconv.Atoi(acc.Api.API_id)
	if err != nil {
		return fmt.Errorf("invalid api_id: %w", err)
	}

	sessionStorage := &session.FileStorage{
		Path: fmt.Sprintf(".sessions/%s.json", acc.Api.ID),
	}
	s := fmt.Sprintf("%s:%s", acc.Proxy.Ip, acc.Proxy.Port)
	sock5, err := proxy.SOCKS5("tcp", s, nil, proxy.Direct)

	if err != nil {
		return fmt.Errorf("proxy dialer: %w", err)
	}

	dc := sock5.(proxy.ContextDialer)
	client := telegram.NewClient(apiID, acc.Api.API_hash, telegram.Options{
		SessionStorage: sessionStorage,
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dc.DialContext,
		}),
	})

	return client.Run(context.Background(), func(ctx context.Context) error {
		status, err := client.Auth().Status(ctx)
		if err != nil {
			return fmt.Errorf("auth status: %w", err)
		}
		if !status.Authorized {
			return fmt.Errorf("account %s is not authorized", acc.Api.ID)
		}

		manager := peers.Options{}.Build(client.API())
		fmt.Println("Man")
		if err := manager.Init(ctx); err != nil {
			return fmt.Errorf("peers init: %w", err)
		}

		var peer peers.Peer

		if id, err := strconv.ParseInt(acc.Msg.Channel_id, 10, 64); err == nil {
			ch, err := manager.GetChannel(ctx, &tg.InputChannel{ChannelID: id})
			if err == nil {
				peer = ch
			} else {
				chat, err := manager.GetChat(ctx, id)
				if err != nil {
					return fmt.Errorf("chat %d not found (not in dialogs?): %w", id, err)
				}
				peer = chat
			}
		} else {
			username := acc.Msg.Channel_id
			fmt.Println("Username")
			if len(username) > 0 && username[0] == '@' {
				username = username[1:]
			}
			peer, err = manager.Resolve(ctx, username)
			if err != nil {
				return fmt.Errorf("resolve %q: %w", username, err)
			}
		}

		fmt.Printf("Resolved peer: %+v\n", peer.InputPeer())

		sender := tgmessage.NewSender(client.API())
		if _, err := sender.To(peer.InputPeer()).Text(ctx, acc.Msg.Msg); err != nil {
			return fmt.Errorf("send failed: %w", err)
		}

		return nil
	})
}

func Auth(acc Account) error {
	apiID, err := strconv.Atoi(acc.Api.API_id)
	if err != nil {
		return fmt.Errorf("invalid api_id: %w", err)
	}

	sessionStorage := &session.FileStorage{
		Path: fmt.Sprintf(".sessions/%s.json", acc.Api.ID),
	}

	s := fmt.Sprintf("%s:%s", acc.Proxy.Ip, acc.Proxy.Port)
	sock5, err := proxy.SOCKS5("tcp", s, nil, proxy.Direct)

	if err != nil {
		return fmt.Errorf("proxy dialer: %w", err)
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
					fmt.Printf("Enter code for %s: ", acc.Api.Number)
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
