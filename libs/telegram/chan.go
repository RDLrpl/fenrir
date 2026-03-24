package telegram

import (
	"context"
	"fmt"

	"github.com/gotd/td/tg"
)

func JoinTGChan(acc Account) error {
	client, err := GenerateTGClient(acc)
	if err != nil {
		return err
	}

	return client.Run(context.Background(), func(ctx context.Context) error {
		api := client.API()

		resolved, err := api.ContactsResolveUsername(ctx, &tg.ContactsResolveUsernameRequest{
			Username: acc.Msg.Channel_id[1:],
		})
		if err != nil {
			return err
		}

		if len(resolved.Chats) == 0 {
			return fmt.Errorf("[FENRIR] Channel not found")
		}

		chat := resolved.Chats[0]
		channel, ok := chat.(*tg.Channel)
		if !ok {
			return fmt.Errorf("[FENRIR] Not Chan or you was banned")
		}
		_, err = api.ChannelsJoinChannel(ctx, &tg.InputChannel{
			ChannelID:  channel.ID,
			AccessHash: channel.AccessHash,
		})

		return err
	})
}
