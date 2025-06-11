package cmd

import (
	"yasm/config"

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
	}

	return app.Run(args)
}
