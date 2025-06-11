package cmd

import (
	"fmt"
	"os"
	"strings"
	"yasm/config"
	"yasm/fzf"
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
		Flags: globalFlags(),
		Commands: []*cli.Command{
			createCommand(),
			editCommand(),
			deleteCommand(),
		},
		Action: func(c *cli.Context) error {
			if shell := c.String("shell"); shell != "" {
				scripts.PrintShellWrapper(strings.ToLower(shell))
			}

			scriptsDir := config.GetScriptsDir()
			entries, err := os.ReadDir(scriptsDir)
			if err != nil {
				return fmt.Errorf("failed to read scripts directory: %w", err)
			}

			var scriptNames []string
			for _, entry := range entries {
				if entry.Type().IsRegular() {
					scriptNames = append(scriptNames, entry.Name())
				}
			}

			if len(scriptNames) == 0 {
				fmt.Println("No scripts found in", scriptsDir)
				return nil
			}

			selected, err := fzf.FuzzySelectScript(scriptNames)
			if err != nil {
				return fmt.Errorf("fzf selection failed: %w", err)
			}

			if selected != "" {
				fmt.Printf("yasm %s\n", selected)
			}

			return nil
		},
	}

	return app.Run(args)
}
