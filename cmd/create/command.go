package create

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// createCommand returns a new CLI command for creating scripts.
// This command allows the user to specify a script name and an optional language.
// If no language is specified, "bash" is used by default.
// The command opens the new script in the editor and supports listing available languages.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create a new script and open in editor",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				DefaultText: "bash",
				Usage:       "Language for the script",
				Value:       "bash",
			},
		},
		UsageText: "" +
			"yasm create <script-name>    				Open new script\n" +
			"yasm create [options] <script-name>  Open new script with specified language template\n" +
			"\n" +
			"To view available languages, run:\n" +
			"    yasm config list-languages",

		Action: func(c *cli.Context) error {
			// Get the script name from the command arguments
			scriptName := c.Args().First()
			if c.Args().Len() == 0 {
				// Show error and usage
				fmt.Fprintf(os.Stderr, "Error: <script-name> must be provided.\n\n")
				cli.ShowSubcommandHelp(c) // show help for the 'create' command
				return cli.Exit("", 1)
			}

			// If no language is specified by command, default to "bash"
			lang := c.String("lang")
			if lang == "" {
				lang = "bash"
			}

			return createScript(scriptName, lang)
		},
	}
}
