package actions

import (
	"fmt"
	"yasm/script"

	"github.com/urfave/cli/v2"
)

func listLanguages() *cli.Command {
	return &cli.Command{
		Name:    "list-languages",
		Usage:   "Show all configured languages for scripting",
		Aliases: []string{"ll"},
		Action: func(c *cli.Context) error {
			fmt.Println("Supported languages:")
			script.PrintLanguagesTable()
			return nil
		},
	}
}
