package run

import (
	"fmt"
	"yasm/cmd/shell"
)

func runScript(scriptName string) error {
	// TODO: Use shell integration to make this the next command that is typed in terminal
	// TODO: Implement the logic to run the script
	if scriptName != "" {
		fmt.Printf("%syasm %s\n", shell.InjectPrefix, scriptName)
	}

	return nil
}
