package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

func Run() error {
	app := cli.App{
		Commands: []*cli.Command{serveCMD},
	}

	return app.Run(os.Args)
}
