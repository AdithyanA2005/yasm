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
// prompting until a valid response is received. Any input errors will
// result in a false return value.
func Confirm(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/n]: ", prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return false
		}

		// Tring spaces and convert to lowercase
		input = strings.TrimSpace(strings.ToLower(input))

		// Switch will only return after user enters a valid response
		switch input {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Please enter 'y' or 'n'.")
		}
	}
}
