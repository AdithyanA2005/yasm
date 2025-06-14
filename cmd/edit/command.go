package edit

import (
	"yasm/fzf"

	"github.com/urfave/cli/v2"
)

// editCommand returns a new CLI command for editing scripts.
// This command allows the user to edit an existing script by name,
// or open a menu to select and edit a script if no name is provided.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "edit",
		Usage: "Edit an existing script in editor",
		UsageText: "" +
			"yasm edit                Open interactive menu to select the script\n" +
			"yasm edit <script-name>  Edit the script with specified name",

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

			return editScript(scriptName)
		},
	}
}
