package commands

import (
	"github.com/persiangophers/userland/cmd/userland/api"
	"github.com/persiangophers/userland/cmd/userland/user"
	"github.com/spf13/cobra"
	"log"
)

var cmdList = []*cobra.Command{
	{
		Use:   "start-api",
		Short: "Start the api server",
		Long: `
			Start the api server on preferred port (default is :3000).
			Use 'userland start-api :<port>'
			to start api on your preferred port`,
		Run: func(cmd *cobra.Command, args []string) {

			port := ""

			if len(args) > 0 {
				port = args[0]
			}

			log.Fatalf("Error occured : %s", api.StartApi(port))
		},
	},
	{
		Use:   "login",
		Short: "Login to Panel",
		Long:  `Proceeding login you will receive a access token to manage things inside server`,
		Run: func(cmd *cobra.Command, args []string) {
			err := user.Login()
			if err != nil {
				log.Fatalf("Error occured : %s", err)
			}
		},
	},
	{
		Use:   "add",
		Short: "Add user manually",
		Long: `Manually add user to app database using
								userland add <database>`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				switch args[0] {
				case "user":
					err := user.Add()
					if err != nil {
						log.Fatalf("Error occured : %s", err)
					}
				}
			} else {
				log.Fatal("You to specify a database")
			}
		},
	},
}

var RootCmd = &cobra.Command{
	Use:     "userland",
	Short:   "userland is a cms for managing users",
	Long:    `Authorize, authenticate and identify users`,
	Version: "1.0.0",
}

func AddCommands() {
	for _, cmd := range cmdList {
		RootCmd.AddCommand(cmd)
	}
}
