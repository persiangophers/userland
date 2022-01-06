package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

var (
	commands = []*cli.Command{
		restCMD,
	}

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`. default value is `config.yaml`",
			Value:   "config.yaml",
		},
	}

	banner = `
	╦ ╦╔═╗╔═╗╦═╗╦  ╔═╗╔╗╔╔╦╗
	║ ║╚═╗║╣ ╠╦╝║  ╠═╣║║║ ║║
	╚═╝╚═╝╚═╝╩╚═╩═╝╩ ╩╝╚╝═╩╝
	`
)

func Run() error {
	fmt.Println(banner)

	app := &cli.App{
		Commands: commands,
		Flags:    flags,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app.Run(os.Args)
}
