package input

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
)

func GetUserInput(inputName string, required bool, obscure bool) string {
	fmt.Print("Enter " + inputName + ": ")

	var input string

	for input == "" && required {
		if obscure {
			password, _ := terminal.ReadPassword(0)
			input = string(password)
			fmt.Print("\n")
		} else {
			_, _ = fmt.Scanln(&input)
		}

		if input == "" {
			fmt.Print(inputName + " can't be empty. Please try again: ")
		}
	}

	return input
}
