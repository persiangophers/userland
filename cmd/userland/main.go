package main

import (
	"fmt"
	"github.com/persiangophers/userland/cmd/userland/commands"
	"os"
)

func main() {
	commands.AddCommands()
	if err := commands.RootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
