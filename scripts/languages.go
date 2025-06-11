package scripts

import (
	"os"
	"sort"
	"yasm/config"

	"github.com/olekukonko/tablewriter"
)

var presetLanguages = map[string]config.LanguageDef{
	"bash": {
		Shebang: "#!/usr/bin/env bash",
		Comment: "#",
	},
	"python": {
		Shebang: "#!/usr/bin/env python3",
		Comment: "#",
	},
	"sh": {
		Shebang: "#!/bin/sh",
		Comment: "#",
	},
	"zsh": {
		Shebang: "#!/usr/bin/env zsh",
		Comment: "#",
	},
}

func GetLanguages() map[string]config.LanguageDef {
	all := make(map[string]config.LanguageDef)

	// Start with built-in
	for k, v := range presetLanguages {
		all[k] = v
	}

	// Override or extend with config
	for k, v := range config.GetLanguages() {
		all[k] = v
	}

	return all
}

func PrintLanguagesTable() {
	langs := GetLanguages()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Language", "Shebang", "Comment"})

	// Sort keys for consistent output
	keys := make([]string, 0, len(langs))
	for lang := range langs {
		keys = append(keys, lang)
	}
	sort.Strings(keys)

	for _, lang := range keys {
		def := langs[lang]
		table.Append([]string{lang, def.Shebang, def.Comment})
	}

	table.Render()
}

// Supported shebangs by language
// var shebangs = map[string]string{
// 	"bash":   "#!/usr/bin/env bash",
// 	"python": "#!/usr/bin/env python3",
// 	"sh":     "#!/bin/sh",
// 	"zsh":    "#!/usr/bin/env zsh",
// }
//
// func ListSupportedLanguages() map[string]string {
// 	return shebangs
// }
