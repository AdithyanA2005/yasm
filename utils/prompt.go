package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Confirm displays a prompt to the user and waits for a yes/no response.
// It returns true if the user enters "y" or "yes" (case-insensitive),
// and false if the user enters "n" or "no". The function will keep
// prompting until a valid response is received or a default choice is given.
// Any input errors will result in a false return value.
func Confirm(prompt string, defaultChoice string) bool {
	reader := bufio.NewReader(os.Stdin)

	// Normalize and validate default choice
	defaultChoice = strings.ToLower(defaultChoice)
	var options string
	switch strings.ToLower(defaultChoice) {
	case "y", "yes":
		options = "Y/n"
		defaultChoice = "yes"
	case "n", "no":
		options = "y/N"
		defaultChoice = "no"
	default:
		options = "y/n"
		defaultChoice = ""
	}

	for {
		fmt.Printf("%s (%s): ", prompt, options)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return false
		}

		input = strings.TrimSpace(strings.ToLower(input))

		// Switch will only return after user enters a valid response
		switch input {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			switch defaultChoice {
			case "yes":
				return true
			case "no":
				return false
			default:
				fmt.Println("Please enter 'y' or 'n'.")
			}
		}
	}
}
