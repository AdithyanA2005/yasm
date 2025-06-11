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
		UsageText: "" +
			"yasm create <script-name>    									Open new script in editor.\n" +
			"yasm create --lang <language> <script-name>    Open new script with specified language in editor.\n" +
			"\n" +
			"To view available languages, run:\n" +
			"    yasm --list-languages\n" +
			"\n" +
			"To know more about languages, run:\n" +
			"    yasm config languages",

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
