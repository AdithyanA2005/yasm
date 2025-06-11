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
				scripts.PrintLanguagesTable()
				return nil
			}

			// Read scripts directory and store entries sorted by name
			scriptsDir := config.GetScriptsDir()
			entries, err := os.ReadDir(scriptsDir)
			if err != nil {
				return fmt.Errorf("failed to read scripts directory: %w", err)
			}

			// Collect the names of all regular files from the entries slice.
			// Only files (not directories or other types) are included in the scriptNames slice.
			var scriptNames []string
			for _, entry := range entries {
				if entry.Type().IsRegular() {
					scriptNames = append(scriptNames, entry.Name())
				}
			}

			// If no scripts found, print message and exit
			if len(scriptNames) == 0 {
				fmt.Println("No scripts found in", scriptsDir)
				return nil
			}

			// Run fzf to select a script
			selected, err := fzf.FuzzySelect(scriptNames)
			if err != nil {
				return fmt.Errorf("fzf selection failed: %w", err)
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
