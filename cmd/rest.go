package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	restCMD = &cli.Command{
		Name:    "rest",
		Aliases: []string{"r"},
		Usage:   "The application starts the serving via rest API",
		Action:  rest,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Port for serving. if not used, the application uses its own config",
			},
		},
	}
)

func rest(c *cli.Context) error {
	fmt.Println("serving port:", c.Int("port"))
	return nil
}
