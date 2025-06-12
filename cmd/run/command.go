package run

import (
	"yasm/fzf"

	"github.com/urfave/cli/v2"
)

// Command for running scripts.
// If no script name is provided as an argument, it opens a fuzzy finder menu
// to select a script. If a script name is provided, it runs the specified script.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "Run a particular script",
		UsageText: "" +
			"yasm run                Open a menu to select the script to run.\n" +
			"yasm run <script-name>  Directly run the script with the given name.",

		Action: func(c *cli.Context) error {
			scriptName := c.Args().First()

			if scriptName == "" {
				// Fuzzy select a script if it was not provided as an argument
				selected, err := fzf.FuzzySelectScript()
				if err != nil {
					return err
				}

				// If no script was selected, return nil (no error)
				if selected == "" {
					return nil
				}

				// Update scriptName to the selected script
				scriptName = selected
			}

			return runScript(scriptName)
		},
	}
}
