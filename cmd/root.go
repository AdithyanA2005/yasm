package cmd

import (
	"fmt"
	"os"
	"yasm/cmd/actions"
	"yasm/cmd/create"
	"yasm/cmd/edit"
	"yasm/cmd/remove"
	"yasm/cmd/run"
	"yasm/cmd/shell"
	"yasm/usercfg"

	"github.com/urfave/cli/v2"
)

func Execute(args []string) error {
	app := &cli.App{
		Name:    "yasm",
		Usage:   "Yet Another Script Manager",
		Version: "0.1.0",
		UsageText: "" +
			"yasm                    Opens an interactive menu to run a script\n" +
			"yasm [command options]  Use --help to see usage of particular command",
		Before: func(c *cli.Context) error {
			usercfg.InitConfig()
			return nil
		},
		DefaultCommand: "run",
		Commands: []*cli.Command{
			create.Command(),
			run.Command(),
			edit.Command(),
			remove.Command(),
			actions.Command(),
			shell.Command(),
		},
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(os.Stderr, "Error: '%s' is not a valid command.\n\n", command)
			cli.ShowAppHelp(c)
		},
	}

	return app.Run(args)
}
