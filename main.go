package main

import (
	"log"

	"github.com/persiangophers/userland/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
