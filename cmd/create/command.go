package create

import (
	"github.com/urfave/cli/v2"
)

// createCommand returns a new CLI command for creating scripts.
// This command allows the user to specify a script name and an optional language.
// If no language is specified, "bash" is used by default.
// The command opens the new script in the editor and supports listing available languages.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create a new script",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Usage: "Language for the script (bash, python, etc).",
				Value: "bash",
			},
		},
		UsageText: "" +
			"yasm create <script-name>    							  Open new script in editor.\n" +
			"yasm create --lang <language> <script-name>  Open new script with specified language in editor.\n" +
			"\n" +
			"To view available languages, run:\n" +
			"    yasm --list-languages\n" +
			"\n" +
			"To know more about languages, run:\n" +
			"    yasm config languages",

		Action: func(c *cli.Context) error {
			scriptName := c.Args().First()
			if scriptName == "" {
				return cli.Exit("Please provide a script name.", 1)
			}

			lang := c.String("lang")
			return createScript(scriptName, lang)
		},
	}
}
