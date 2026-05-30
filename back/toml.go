package back

import (
	"fmt"
	"log"
	"os"

	"github.com/RDLrpl/fenrir/utility"
	"github.com/pelletier/go-toml/v2"
)

func ParseConfig() utility.Configuration {
	var config utility.Configuration

	conf_file, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal(conf_file, &config)
	if err != nil {
		panic(err)
	}

	return config
}
func createDefaultConfig() {
	defaultConfig := utility.Configuration{
		Messages: make(map[string]string),
		Telegram: utility.TelegramConfig{
			Sessions: ".sessions",
			Targets:  make(map[string]string),
			Accounts: make(map[string]utility.TelegramAccount),
		},
		Discord: utility.DiscordConfig{
			Targets:  make(map[string][]string),
			Accounts: make(map[string]utility.DiscordAccount),
		},
	}

	data, err := toml.Marshal(defaultConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("config.toml", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateOrCheckConf() {
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		createDefaultConfig()
		return
	}
	confFile, err := os.ReadFile("config.toml")
	if err != nil {
		createDefaultConfig()
		return
	}

	var config utility.Configuration
	err = toml.Unmarshal(confFile, &config)
	if err != nil {
		fmt.Println("BAD CONFIG, RECREATING!")
		createDefaultConfig()
	}
}

func saveConfig(config utility.Configuration) {
	data, err := toml.Marshal(config)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = os.WriteFile("config.toml", data, 0644)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
