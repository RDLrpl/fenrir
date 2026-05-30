package back

import (
	"github.com/RDLrpl/fenrir/utility"
)

// DISCORD

func AddDsAccount(name string, account utility.DiscordAccount) {
	config := ParseConfig()
	if config.Discord.Accounts == nil {
		config.Discord.Accounts = make(map[string]utility.DiscordAccount)
	}
	config.Discord.Accounts[name] = account
	saveConfig(config)
}

func RemoveDsAccount(name string) {
	config := ParseConfig()
	if config.Discord.Accounts == nil {
		return
	}
	if _, exists := config.Discord.Accounts[name]; exists {
		delete(config.Discord.Accounts, name)
		saveConfig(config)
	}
}

func GetDsAccount(name string) (utility.DiscordAccount, bool) {
	config := ParseConfig()
	if config.Discord.Accounts == nil {
		return utility.DiscordAccount{}, false
	}
	acc, exists := config.Discord.Accounts[name]
	return acc, exists
}

func GetDsAccounts() map[string]utility.DiscordAccount {
	config := ParseConfig()

	return config.Discord.Accounts
}

// TELEGRAM

func AddTgAccount(name string, account utility.TelegramAccount) {
	config := ParseConfig()
	if config.Telegram.Accounts == nil {
		config.Telegram.Accounts = make(map[string]utility.TelegramAccount)
	}
	config.Telegram.Accounts[name] = account
	saveConfig(config)
}

func RemoveTgAccount(name string) {
	config := ParseConfig()
	if config.Telegram.Accounts == nil {
		return
	}
	if _, exists := config.Telegram.Accounts[name]; exists {
		delete(config.Telegram.Accounts, name)
		saveConfig(config)
	}
}

func GetTgAccount(name string) (utility.TelegramAccount, bool) {
	config := ParseConfig()
	if config.Telegram.Accounts == nil {
		return utility.TelegramAccount{}, false
	}
	acc, exists := config.Telegram.Accounts[name]
	return acc, exists
}

func GetTgAccounts() map[string]utility.TelegramAccount {
	config := ParseConfig()

	return config.Telegram.Accounts
}

func AddTgTarget(key string, target string) {
	config := ParseConfig()
	if config.Telegram.Targets == nil {
		config.Telegram.Targets = make(map[string]string)
	}
	config.Telegram.Targets[key] = target
	saveConfig(config)
}

func RemoveTgTarget(key string) {
	config := ParseConfig()
	if config.Telegram.Targets == nil {
		return
	}
	if _, exists := config.Telegram.Targets[key]; exists {
		delete(config.Telegram.Targets, key)
		saveConfig(config)
	}
}

func AddDsTarget(key string, targets []string) {
	config := ParseConfig()
	if config.Discord.Targets == nil {
		config.Discord.Targets = make(map[string][]string)
	}
	config.Discord.Targets[key] = targets
	saveConfig(config)
}

func RemoveDsTarget(key string) {
	config := ParseConfig()
	if config.Discord.Targets == nil {
		return
	}
	if _, exists := config.Discord.Targets[key]; exists {
		delete(config.Discord.Targets, key)
		saveConfig(config)
	}
}

func GetDSTargets() map[string][]string {
	config := ParseConfig()

	return config.Discord.Targets
}

func AddMessage(key string, message string) {
	config := ParseConfig()
	if config.Messages == nil {
		config.Messages = make(map[string]string)
	}
	config.Messages[key] = message
	saveConfig(config)
}

func RemoveMessage(key string) {
	config := ParseConfig()
	if config.Messages == nil {
		return
	}
	if _, exists := config.Messages[key]; exists {
		delete(config.Messages, key)
		saveConfig(config)
	}
}
