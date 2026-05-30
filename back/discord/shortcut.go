package discord

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/RDLrpl/fenrir/back"
	"github.com/RDLrpl/fenrir/utility"
)

func Asendds(ctx context.Context) {
	conf := back.ParseConfig()

	var wg sync.WaitGroup
	delayBetweenMessages := 200 * time.Millisecond

	for _, acc := range conf.Discord.Accounts {
		if ctx.Err() != nil {
			break
		}

		targets, exists := conf.Discord.Targets[acc.Targ]
		if !exists || len(targets) == 0 {
			continue
		}

		wg.Add(1)
		go func(channel_id string, message string, acc utility.DiscordAccount) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				dg, err := SimpleLogIn(acc.Token, acc.Proxy, acc.ProxyLogin, acc.ProxyPass)
				if err != nil {
					return
				}

				mes, err := dg.ChannelMessageSend(channel_id, message)
				if err != nil {
					return
				}

				fmt.Printf("|- Message OK: %s (%s)\n", mes.ID, acc.Token[:8])

				select {
				case <-ctx.Done():
					return
				case <-time.After(220 * time.Millisecond):
				}
			}
		}(targets[0], conf.Messages[acc.Marg], acc)

		select {
		case <-ctx.Done():
			break
		case <-time.After(delayBetweenMessages):
		}
	}
	wg.Wait()

	fmt.Println("|- STOPED")
}
