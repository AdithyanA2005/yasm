package shell

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// editCommand returns a new CLI command for editing scripts.
// This command allows the user to edit an existing script by name,
// or open a menu to select and edit a script if no name is provided.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "shell",
		Usage: "Print shell script used to execute yasm",
		// TODO: Add usage text for the command
		UsageText: "" +
			"yasm shell <shell-name>  Print shell script for a particular shell\n" +
			"\n" +
			"To view a list of supprted shells, run:\n" +
			"    yasm actions list-shells",
		Action: func(c *cli.Context) error {
			// Show error and usage if no arguments are given
			if c.Args().Len() == 0 {
				fmt.Fprintf(os.Stderr, "Error: <shell-name> must be provided.\n\n")
				cli.ShowSubcommandHelp(c) // show help for the 'shell' command
				return cli.Exit("", 1)
			}

			// Get the shell name from the command arguments
			shellName := c.Args().First()

			return integrateShell(shellName)
		},
	}
}
