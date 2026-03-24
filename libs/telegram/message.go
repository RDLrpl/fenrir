package telegram

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	tgmessage "github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/peers"
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

	client, err := GenerateTGClient(acc)
	if err != nil {
		return err
	}

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

		peer, err = manager.Resolve(ctx, acc.Msg.Channel_id[1:])
		if err != nil {
			return fmt.Errorf("[FENRIR] TG!E!(resolve?%q): %v", acc.Msg.Channel_id[1:], err)
		}

		sender := tgmessage.NewSender(client.API())
		if _, err := sender.To(peer.InputPeer()).Text(ctx, acc.Msg.Msg); err != nil {
			return fmt.Errorf("[FENRIR] TG!E!(SEND-ERROR): %v", err)
		}

		logger.Info(fmt.Sprintf("[FENRIR]: Sended From >> %v", acc.Id))

		return nil
	})
}
