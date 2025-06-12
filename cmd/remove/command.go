package remove

import (
	"yasm/fzf"

	"github.com/urfave/cli/v2"
)

// deleteCommand returns a new CLI command for deleting scripts.
// This command allows the user to delete an existing script by name,
// or open a menu to select and delete a script if no name is provided.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Delete an existing script",
		UsageText: "" +
			"yasm delete                Open interactive menu to select the script\n" +
			"yasm delete <script-name>  Delete the script with specified name",

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

			return removeScript(scriptName)
		},
	}
}
