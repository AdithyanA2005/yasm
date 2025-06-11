package cmd

import (
	"yasm/scripts"

	"github.com/urfave/cli/v2"
)

// deleteCommand returns a new CLI command for deleting scripts.
// This command allows the user to delete an existing script by name,
// or open a menu to select and delete a script if no name is provided.
func deleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Delete an existing script",
		UsageText: "" +
			"yasm delete                Open a menu to select and delete a script.\n" +
			"yasm delete <script-name>  Delete the script with given name directy.",

		Action: func(c *cli.Context) error {
			scriptName := c.Args().First()
			return scripts.DeleteScript(scriptName)
		},
	}
}
