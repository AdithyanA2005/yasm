package cmd

import (
	"fmt"
	"os"
	"strings"
	"yasm/cmd/actions"
	"yasm/cmd/create"
	"yasm/cmd/edit"
	"yasm/cmd/remove"
	"yasm/cmd/run"
	"yasm/script"
	"yasm/usercfg"

	"github.com/urfave/cli/v2"
)

func Execute(args []string) error {
	app := &cli.App{
		Name:    "yasm",
		Usage:   "Yet Another Script Manager",
		Version: "0.1.0",
		Before: func(c *cli.Context) error {
			usercfg.LoadConfig()
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
		DefaultCommand: "run",
		Commands: []*cli.Command{
			create.Command(),
			run.Command(),
			edit.Command(),
			remove.Command(),
			actions.Command(),
		},
		Action: func(c *cli.Context) error {
			// If to run shell setup, do it and exit the program
			if shell := c.String("shell"); shell != "" {
				generateShellWrapper(strings.ToLower(shell))
				os.Exit(0)
			}

			// Handle list-languages flag
			if c.Bool("list-languages") {
				fmt.Println("Supported languages:")
				script.PrintLanguagesTable()
				return nil
			}

			return nil
		},
	}

	return app.Run(args)
}
