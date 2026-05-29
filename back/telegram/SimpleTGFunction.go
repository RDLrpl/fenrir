package telegram

import (
	"context"
	"fmt"

	"github.com/RDLrpl/fenrir/utility"
	"github.com/gotd/td/telegram"
	tgmessage "github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/peers"
	"github.com/gotd/td/tg"
)

func JoinTGChan(target string, storage string, accountid string, account utility.TelegramAccount) error {
	client, err := RegOrLog(storage, accountid, account)
	if err != nil {
		return err
	}

	return client.Run(context.Background(), func(ctx context.Context) error {
		api := client.API()

		resolved, err := api.ContactsResolveUsername(ctx, &tg.ContactsResolveUsernameRequest{
			Username: target[1:],
		})
		if err != nil {
			return err
		}

		if len(resolved.Chats) == 0 {
			return fmt.Errorf("TG!E! BAD Channel! Not Found! Use guide! ")
		}

		chat := resolved.Chats[0]
		channel, ok := chat.(*tg.Channel)
		if !ok {
			return fmt.Errorf("TG!E! BAD Channel or you was banned!")
		}
		_, err = api.ChannelsJoinChannel(ctx, &tg.InputChannel{
			ChannelID:  channel.ID,
			AccessHash: channel.AccessHash,
		})

		return err
	})
}

func SendTGmessage(target string, message string, storage string, accountid string, account utility.TelegramAccount) error {
	client, err := RegOrLog(storage, accountid, account)
	if err != nil {
		return err
	}

	return client.Run(context.Background(), func(ctx context.Context) error {
		status, err := client.Auth().Status(ctx)
		if err != nil {
			return fmt.Errorf("TG!E! SESSION ERROR: %v", err)
		}
		if !status.Authorized {
			return fmt.Errorf("TG!E! ACCOUNT{%s} SESSION NOT AUTH: use fenrir --auth", accountid)
		}

		manager := peers.Options{}.Build(client.API())
		if err := manager.Init(ctx); err != nil {
			return fmt.Errorf("TG!E! PEERS: %v", err)
		}

		var peer peers.Peer

		peer, err = manager.Resolve(ctx, target[1:])
		if err != nil {
			return fmt.Errorf("TG!E! resolve?%q: %v", target[1:], err)
		}

		sender := tgmessage.NewSender(client.API())
		if _, err := sender.To(peer.InputPeer()).Text(ctx, message); err != nil {
			fmt.Printf("|- TG!E! SEND-WARN |prob. SpamBlock or Timeout|: %v", err)
		}

		fmt.Printf("|- TG!O! SUCCESS >> %v", accountid)

		return nil
	})
}

func ASendTGMessage(ctx context.Context, target string, message string, accountid string, client *telegram.Client) error {
	manager := peers.Options{}.Build(client.API())
	if err := manager.Init(ctx); err != nil {
		return fmt.Errorf("TG!E! PEERS: %v", err)
	}

	var peer peers.Peer
	var err error

	if len(target) > 1 && (target[0] == '@' || target[0] == 't') {
		peer, err = manager.Resolve(ctx, target[1:])
	} else {
		peer, err = manager.Resolve(ctx, target)
	}

	if err != nil {
		return fmt.Errorf("TG!E! resolve?%q: %v", target, err)
	}

	sender := tgmessage.NewSender(client.API())
	if _, err := sender.To(peer.InputPeer()).Text(ctx, message); err != nil {
		fmt.Printf("|- TG!E! SEND-WARN |prob. SpamBlock or Timeout|: %v\n", err)
		return err
	}

	fmt.Printf("|- TG!O! SUCCESS >> %v\n", accountid)
	return nil
}
