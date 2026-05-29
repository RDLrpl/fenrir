package discord

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/RDLrpl/fenrir/utility"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/net/proxy"
)

func SimpleLogIn(token string, proxyf string, proxyf_login string, proxyf_pass string) (*discordgo.Session, error) {
	host, log, pass, good := utility.Check_proxy(proxyf, proxyf_login, proxyf_pass)

	httpClient := http.DefaultClient

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

		transport := &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return sock5.Dial(network, addr)
			},
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}

		httpClient = &http.Client{
			Transport: transport,
		}
	}

	dg, err := discordgo.New(token)
	if err != nil {
		return nil, fmt.Errorf("session: %v", err)
	}

	dg.Client = httpClient

	return dg, nil
}

func Send(channel_id string, token string, message string, proxy string, proxy_login string, proxy_pass string) error {
	dg, err := SimpleLogIn(token, proxy, proxy_login, proxy_pass)
	if err != nil {
		log.Fatalf("DS!E! %v", err)
	}

	msg, err := dg.ChannelMessageSend(channel_id, message)
	if err != nil {
		log.Fatalf("DS!E! Message Send: %v", err)
	}

	fmt.Printf("|- Message OK: %s\n", msg.ID)
	return nil
}
