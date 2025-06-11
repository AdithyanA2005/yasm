package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Confirm prompts the user for a yes/no confirmation.
func Confirm(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/n]: ", prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return false
		}
		input = strings.TrimSpace(strings.ToLower(input))
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
