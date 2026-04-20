package util

import (
	"fmt"
	"strings"

	"github.com/RDLrpl/fenrir/fenrir-components/cli"
)

func PrintUsage() {
	fmt.Println(
		cli.Red_Style.Render(strings.TrimSpace(cli.FenArt)),
		"\n\n",
		cli.Help_Style.Render(strings.TrimSpace(cli.Usage)),
	)
}
