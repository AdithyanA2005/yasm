package actions

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "actions",
		Usage: "Actions to manage scripts and configurations",
		Subcommands: []*cli.Command{
			listLanguages(),
		},

		Action: func(c *cli.Context) error {
			cli.ShowSubcommandHelp(c)
			return nil
		},
	}
}
