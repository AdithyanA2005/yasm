package shell

import (
	"fmt"
	"os"
	"yasm/utils"

	"github.com/urfave/cli/v2"
)

// integrateShell loads and prints the shell integration script for the given shell.
// Returns an error and lists supported shells if the shell is not supported.
func integrateShell(shell string) error {
	code, err := loadShellScript(shell)
	if err != nil {
		// If shell is not supported, print error and list supported shells
		fmt.Fprintf(os.Stderr, "Error: '%v' is not yet supported by yasm.\n\n", shell)
		fmt.Fprintf(os.Stderr, "Currently supported shells:\n")
		renderShellTable(os.Stderr)
		fmt.Fprintf(os.Stderr, "\nPlease open an issue in the yasm repo if you'd like to see support for '%v':\n", shell)
		fmt.Fprintf(os.Stderr, "%s\n", utils.RepoUrl)

		return cli.Exit("", 1)
	}

	fmt.Println(code)
	return nil
}
