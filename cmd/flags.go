package cmd

import "github.com/urfave/cli/v2"

func globalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "shell",
			Usage:   "Output shell function for injection (bash, zsh, fish)",
			Aliases: []string{"s"},
		},
	}
}
