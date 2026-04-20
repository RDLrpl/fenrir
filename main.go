package main

import (
	"os"

	"github.com/RDLrpl/fenrir/fenrir-components/handlers/util"
)

func main() {
	arguments := os.Args
	args := len(arguments)

	switch {
	case args == 2:
		// fenrir --clean
		switch arguments[1] {
		case "--clean":
			//todo
		case "--run":
			//todo
		default:
			util.PrintUsage()
		}
	case args == 4:
		// fenrir cli tg --auth
		if arguments[1] == "cli" && arguments[2] == "tg" {
			switch arguments[3] {
			case "--auth":
				//todo
			case "--atckm":
				//todo
			case "--atckj":
				//todo
			default:
				util.PrintUsage()
			}
		} else {
			util.PrintUsage()
		}
	default:
		util.PrintUsage()
	}

}
