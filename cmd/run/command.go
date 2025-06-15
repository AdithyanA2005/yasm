package run

import (
	"fmt"

	"github.com/adithyana2005/yasm/fzf"
	"github.com/urfave/cli/v2"
)

// Command for running scripts.
// If no script name is provided as an argument, it opens a fuzzy finder menu
// to select a script. If a script name is provided, it runs the specified script.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "Run an existing script",
		UsageText: "" +
			"yasm run                Open interactive menu to select the run\n" +
			"yasm run <script-name>  Run the script with specified name",

		Action: func(c *cli.Context) error {
			args := c.Args().Slice()

			if len(args) == 0 {
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
				scriptName := selected
				fmt.Println("yasm run", scriptName)
				return nil
			} else {
				scriptName := args[0]
				scriptArgs := args[1:] // Remaining args passed to script
				return runScript(scriptName, scriptArgs)
			}
		},
	}
}
