package actions

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "actions",
		Usage: "Actions to manage scripts and configurations",
		Subcommands: []*cli.Command{
			listLanguages(),
			showConfigs(),
		},

		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				cli.ShowSubcommandHelp(c)
				return nil
			}

			// If a non-command argument (unknown command) is provided, show an error message and usage
			fmt.Fprintf(os.Stderr, "Error: '%s' is not a valid command.\n\n", c.Args().First())
			cli.ShowSubcommandHelp(c) // show help for the 'create' command
			return cli.Exit("", 1)
		},
	}
}
