package telegram

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/dcs"
	tgmessage "github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/peers"
	"github.com/gotd/td/tg"
	"golang.org/x/net/proxy"
)

func SendTGmessage(acc Account) error {
	logger := log.New(os.Stderr)

	styles := log.DefaultStyles()

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Bold(true).
		Foreground(lipgloss.Color("#00ff37")).
		Background(lipgloss.Color("#008643"))

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Bold(true).
		Foreground(lipgloss.Color("#ff0055")).
		Background(lipgloss.Color("#3d0014"))
	logger.SetStyles(styles)

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
		status, err := client.Auth().Status(ctx)
		if err != nil {
			return fmt.Errorf("[FENRIR] TG!E!(AUTH ERROR): %v", err)
		}
		if !status.Authorized {
			return fmt.Errorf("[FENRIR] TG!E!(ACCOUNT{%s} NOT AUTH): use fenrir --auth", acc.Api.ID)
		}

		manager := peers.Options{}.Build(client.API())
		if err := manager.Init(ctx); err != nil {
			return fmt.Errorf("[FENRIR] TG!E!(PEERS): %v", err)
		}

		var peer peers.Peer

		if id, err := strconv.ParseInt(acc.Msg.Channel_id, 10, 64); err == nil {
			ch, err := manager.GetChannel(ctx, &tg.InputChannel{ChannelID: id})
			if err == nil {
				peer = ch
			} else {
				chat, err := manager.GetChat(ctx, id)
				if err != nil {
					return fmt.Errorf("[FENRIR] TG!E!(%d not dialog): %v", id, err)
				}
				peer = chat
			}
		} else {
			username := acc.Msg.Channel_id
			if len(username) > 0 && username[0] == '@' {
				username = username[1:]
			}
			peer, err = manager.Resolve(ctx, username)
			if err != nil {
				return fmt.Errorf("[FENRIR] TG!E!(resolve?%q): %v", username, err)
			}
		}

		sender := tgmessage.NewSender(client.API())
		if _, err := sender.To(peer.InputPeer()).Text(ctx, acc.Msg.Msg); err != nil {
			return fmt.Errorf("[FENRIR] TG!E!(SEND-ERROR): %v", err)
		}

		logger.Info(fmt.Sprintf("[FENRIR]: Sended From >> %v", acc.Id))

		return nil
	})
}
