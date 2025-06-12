package actions

import (
	"fmt"
	"os"
	"sort"
	"yasm/usercfg"
	"yasm/utils"

	"github.com/urfave/cli/v2"
)

func listLanguages() *cli.Command {
	return &cli.Command{
		Name:    "list-languages",
		Usage:   "Show all configured languages for scripting",
		Aliases: []string{"ll"},
		Action: func(c *cli.Context) error {
			fmt.Println("Supported languages:")

			langs := usercfg.GetLanguages()
			headers := []string{"Language", "Shebang", "Comment"}
			rows := buildLanguagesTableRows(langs)
			utils.RenderTable(os.Stdout, headers, rows)

			return nil
		},
	}
}

// buildLanguagesTableRows constructs the rows for the languages table.
// It takes a map of language names to LanguageDef structs and returns a slice of string slices,
// where each inner slice represents a row containing the language name, its shebang, and comment style.
func buildLanguagesTableRows(langs map[string]usercfg.LanguageDef) [][]string {
	keys := make([]string, 0, len(langs))
	for lang := range langs {
		keys = append(keys, lang)
	}
	sort.Strings(keys)

	rows := make([][]string, 0, len(keys))
	for _, lang := range keys {
		def := langs[lang]
		rows = append(rows, []string{lang, def.Shebang, def.Comment})
	}
	return rows
}
