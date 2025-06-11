package cmd

import (
	"fmt"
	"yasm/scripts"

	"github.com/urfave/cli/v2"
)

func createCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create a new script",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Usage: "Language for the script (bash, python, etc).",
				Value: "bash",
			},
		},

		Action: func(c *cli.Context) error {
			scriptName := c.Args().First()
			if scriptName == "" {
				fmt.Println("Please provide a script name.")
				return nil
			}
			lang := c.String("lang")
			return scripts.CreateScript(scriptName, lang)
		},
	}
}
