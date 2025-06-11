package cmd

import (
	"yasm/scripts"

	"github.com/urfave/cli/v2"
)

// editCommand returns a new CLI command for editing scripts.
// This command allows the user to edit an existing script by name,
// or open a menu to select and edit a script if no name is provided.
func editCommand() *cli.Command {
	return &cli.Command{
		Name:  "edit",
		Usage: "Edit an existing script",
		UsageText: "" +
			"yasm edit                Open a menu to select a script and edit it in editor.\n" +
			"yasm edit <script-name>  Edit the script with given name directy in editor.",

		Action: func(c *cli.Context) error {
			scriptName := c.Args().First()
			return scripts.EditScript(scriptName)
		},
	}
}
