package common

import (
	"bufio"
	"maps"
	"os"
	"sort"
	"strings"
	"yasm/config"

	"github.com/olekukonko/tablewriter"
)

// Preset languages with default shebangs and comments
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

// GetLanguages returns a map of all available language definitions.
// It starts with built-in preset languages and then overrides or extends them
// with any user-defined languages from the configuration.
func GetLanguages() map[string]config.LanguageDef {
	all := make(map[string]config.LanguageDef)

	// Start with built-in
	maps.Copy(all, presetLanguages)

	// Override or extend with config
	maps.Copy(all, config.GetLanguages())

	return all
}

// PrintLanguagesTable prints a table of all supported languages to standard output.
// The table includes the language name, its shebang line, and the comment prefix.
// Languages are listed in sorted order for consistent output.
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

func DetectCommentChar(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", nil // empty file
	}
	firstLine := scanner.Text()

	for _, lang := range GetLanguages() {
		if strings.TrimSpace(firstLine) == strings.TrimSpace(lang.Shebang) {
			return lang.Comment, nil
		}
	}

	// fallback
	return "#", nil
}
