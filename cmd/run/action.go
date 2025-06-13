package run

import (
	"fmt"
)

func runScript(scriptName string) error {
	// TODO: Use shell integration to make this the next command that is typed in terminal
	// TODO: Implement the logic to run the script
	if scriptName != "" {
		fmt.Printf("yasm %s\n", scriptName)
	}

	return nil
}
