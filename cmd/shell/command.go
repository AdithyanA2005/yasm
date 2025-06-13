package shell

import (
	"fmt"
	"io"
	"os"
	"yasm/utils"

	"github.com/urfave/cli/v2"
)

// editCommand returns a new CLI command for editing scripts.
// This command allows the user to edit an existing script by name,
// or open a menu to select and edit a script if no name is provided.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "shell",
		Usage: "Print shell script used to execute yasm",
		UsageText: "" +
			"yasm shell <shell-name>  Print shell script for a particular shell\n" +
			"yasm shell --list        Print a list of supported shells",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "list",
				Usage:    "See a list of supported shells",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			showList := c.Bool("list")
			if showList {
				fmt.Println("Supported shells:")
				renderShellTable(os.Stdout)
				return nil
			}

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

// renderShellTable prints a table of all supported shells to provided writer.
// It displays the shell names and the total count in the header.
func renderShellTable(w io.Writer) {
	noOfSupportedLanguages := len(SupportedShells)
	headers := []string{fmt.Sprintf("Shells (%d)", noOfSupportedLanguages)}
	rows := make([][]string, 0, noOfSupportedLanguages)

	// Populate rows with each supported shell name.
	for shell := range SupportedShells {
		rows = append(rows, []string{shell})
	}

	// Render the table to standard output.
	utils.RenderTable(w, headers, rows)
}
