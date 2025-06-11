package cmd

import (
	"yasm/scripts"

	"github.com/urfave/cli/v2"
)

func editCommand() *cli.Command {
	return &cli.Command{
		Name:  "edit",
		Usage: "Edit an existing script",
		Action: func(c *cli.Context) error {
			scriptName := c.Args().First()
			return scripts.EditScript(scriptName)
		},
	}
}
