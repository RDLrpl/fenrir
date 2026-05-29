package back

import (
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
