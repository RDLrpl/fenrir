package handlers

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/RDLrpl/Fenrir/libs/fnlang"
	"github.com/RDLrpl/Fenrir/libs/telegram"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func Telegram_Auth() {
	logger := log.New(os.Stderr)

	styles := log.DefaultStyles()

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Bold(true).
		Foreground(lipgloss.Color("#ff0055")).
		Background(lipgloss.Color("#3d0014"))
	logger.SetStyles(styles)

	err := os.MkdirAll(".sessions", 0755)
	if err != nil {
		logger.Error(err)
	}
	conf, err := fnlang.ReadConfiguration()
	if err != nil {
		logger.Error(err)
	}

	Acc, err := telegram.TG_PairAccounts(conf.Params)
	if err != nil {
		logger.Error(err)
	}

	for _, Account := range Acc.Accs {
		if err := telegram.Auth(Account); err != nil {
			logger.Error(err)
		}
	}
}

func TG_Join() {
	logger := log.New(os.Stderr)
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Bold(true).
		Foreground(lipgloss.Color("#ff0055")).
		Background(lipgloss.Color("#3d0014"))
	logger.SetStyles(styles)

	conf, err := fnlang.ReadConfiguration()
	if err != nil {
		logger.Error(err)
		return
	}

	Acc, err := telegram.TG_PairAccounts(conf.Params)
	if err != nil {
		logger.Error(err)
		return
	}

	var wg sync.WaitGroup

	for _, account := range Acc.Accs {
		wg.Add(1)
		go func(a telegram.Account) {
			defer wg.Done()
			if err := telegram.JoinTGChan(a); err != nil {
				logger.Error(fmt.Sprintf("[Fenrir] User %s: %v", a.Id, err))
				return
			}
			logger.Info(fmt.Sprintf("[Fenrir] Chan %s, User %s: Success", a.Msg.Channel_id, a.Id))
		}(account)
	}

	wg.Wait()
}
func Clean() {
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

	err := os.RemoveAll(".sessions")
	if err != nil {
		logger.Error(err)
		os.Exit(111)
	}
	logger.Info("[Fenrir] 'Clean' Succesfully Complete")
}

func TG_Send() {
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

	conf, err := fnlang.ReadConfiguration()
	if err != nil {
		panic(err)
	}

	Acc, err := telegram.TG_PairAccounts(conf.Params)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	delayBetweenMessages := 200 * time.Millisecond

	for _, Account := range Acc.Accs {
		wg.Add(1)
		go func(acc telegram.Account) {
			defer wg.Done()

			ticker := time.NewTicker(delayBetweenMessages)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					logger.Info(fmt.Sprintf("[FENRIR]: Sending From >> %v", acc.Id))
					err := telegram.SendTGmessage(acc)
					if err != nil {
						logger.Error(fmt.Sprintf("[ERROR]!%v: %v", acc.Id, err))
						return
					}
				}
			}
		}(Account)
	}

	wg.Wait()
}
