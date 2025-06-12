package actions

import (
	"fmt"
	"yasm/usercfg"

	"github.com/urfave/cli/v2"
)

func listLanguages() *cli.Command {
	return &cli.Command{
		Name:    "list-languages",
		Usage:   "Show all configured languages for scripting",
		Aliases: []string{"ll"},
		Action: func(c *cli.Context) error {
			fmt.Println("Supported languages:")
			usercfg.PrintLanguagesTable()
			return nil
		},
	}
}
