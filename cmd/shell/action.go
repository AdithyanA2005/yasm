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
		supported := getSupportedShellNames()

		msg := fmt.Sprintf(
			"Error: '%v' is not yet supported by yasm.\n\n"+
				"For the time being, yasm supports the following shells:\n%s\n"+
				"Please open an issue in the yasm repo if you'd like to see support for '%v':\n%s\n",
			shell, formatShellList(supported), shell, utils.RepoUrl,
		)

		fmt.Fprint(os.Stderr, msg)

		return cli.Exit("", 1)
	}

	fmt.Println(code)
	return nil
}

// formatShellList formats a slice of shell names into a human-readable
// bulleted list, with each shell on its own line prefixed by "  - ".
func formatShellList(shells []string) string {
	var result string
	for _, shell := range shells {
		result += fmt.Sprintf("  - %s\n", shell)
	}
	return result
}
