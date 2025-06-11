package cmd

import (
	"yasm/scripts"

	"github.com/urfave/cli/v2"
)

func deleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Delete an existing script",
		Action: func(c *cli.Context) error {
			scriptName := c.Args().First()
			return scripts.DeleteScript(scriptName)
		},
	}
}
