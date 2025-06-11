package cmd

import (
	"fmt"
	"os"
	"strings"
	"yasm/common"
	"yasm/config"
	"yasm/scripts"

	"github.com/urfave/cli/v2"
)

func Execute(args []string) error {
	app := &cli.App{
		Name:    "yasm",
		Usage:   "Yet Another Script Manager",
		Version: "0.1.0",
		Before: func(c *cli.Context) error {
			config.LoadConfig()
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "shell",
				Usage: "Output shell function for injection (bash, zsh, fish)",
			},
			&cli.BoolFlag{
				Name:    "list-languages",
				Aliases: []string{"l"},
				Usage:   "Show supported script languages",
			},
		},
		Commands: []*cli.Command{
			createCommand(),
			editCommand(),
			deleteCommand(),
		},
		Action: func(c *cli.Context) error {
			// If to run shell setup, do it and exit the program
			if shell := c.String("shell"); shell != "" {
				scripts.PrintShellWrapper(strings.ToLower(shell))
				os.Exit(0)
			}

			// Handle list-languages flag
			if c.Bool("list-languages") {
				fmt.Println("Supported languages:")
				common.PrintLanguagesTable()
				return nil
			}

			// Fuzzy select a script if no name provided
			selected, err := common.FuzzySelectScript()
			if err != nil {
				return err
			}

			// If a script was selected, print the command to run it
			// TODO: Use shell integration to make this the next command that is typed in terminal
			if selected != "" {
				fmt.Printf("yasm %s\n", selected)
			}

			return nil
		},
	}

	return app.Run(args)
}
