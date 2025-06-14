package actions

import (
	"fmt"
	"os"
	"yasm/usercfg"
	"yasm/utils"

	"github.com/urfave/cli/v2"
)

func showConfigs() *cli.Command {
	return &cli.Command{
		Name:    "show-configs",
		Usage:   "Display current configuration and language settings",
		Aliases: []string{"sc"},
		Action: func(c *cli.Context) error {
			fmt.Println("Configuration:")
			renderConfigTable()

			fmt.Println("\nLanguages:")
			renderLanguageTable()

			return nil
		},
	}
}

// renderConfigTable renders a table displaying the user's configuration values.
// It shows the current value and the default value for each configuration field.
func renderConfigTable() {
	headers := []string{"Field", "Value", "Default Value"}
	rows := [][]string{
		{
			"srcripts-dir",
			usercfg.GetScriptsDir(),
			usercfg.LoadedConfig.ScriptsDir,
		},
		{
			"editor",
			usercfg.GetEditor(),
			usercfg.LoadedConfig.Editor,
		},
		{
			"add-scripts-to-path",
			fmt.Sprintf("%t", usercfg.ShouldAddScriptsToPath()),
			fmt.Sprintf("%t", usercfg.LoadedConfig.AddScriptsToPath),
		},
	}

	utils.RenderTable(os.Stdout, headers, rows)
}
