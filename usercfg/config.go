package usercfg

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	toml "github.com/pelletier/go-toml/v2"
)

var (
	userHomeDir   string
	userConfigDir string
	loadedConfig  *ConfigDef
)

func LoadConfig() {
	var err error

	userHomeDir, err = os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not determine home directory:  %v\n", err)
		os.Exit(1)
	}

	userConfigDir, err = os.UserConfigDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not determine config directory:  %v\n", err)
		os.Exit(1)
	}

	configPaths := []string{
		filepath.Join(userConfigDir, "yasm", "config.toml"),
		filepath.Join(userHomeDir, ".config", "yasm", "config.toml"),
		filepath.Join(userHomeDir, "yasm", "config.toml"),
	}

	for _, path := range configPaths {
		if cfg, err := readTomlFile(path); err == nil {
			loadedConfig = cfg
			return
		} else if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Failed to load config from %s: %v\n", path, err)
		}
	}

	// Fallback to empty config
	loadedConfig = &ConfigDef{}
}


func readTomlFile(path string) (*ConfigDef, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg ConfigDef
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetEditor() string {
	if loadedConfig == nil {
		LoadConfig()
	}

	if loadedConfig.Editor != "" {
		return loadedConfig.Editor
	}

	editor := os.Getenv("EDITOR")
	if editor != "" {
		return editor
	}

	return "nano"
}

func GetScriptsDir() string {
	if loadedConfig == nil {
		LoadConfig()
	}

	sd := loadedConfig.ScriptsDir
	if sd != "" {
		return expandPath(sd)
	}

	return filepath.Join(os.Getenv("HOME"), "scripts")
}

func IsAddScriptsToPathEnabled() bool {
	if loadedConfig == nil {
		LoadConfig()
	}
	return loadedConfig.AddScriptsToPath
}

// expandPath replaces ~ with full home path
func expandPath(path string) string {
	if strings.HasPrefix(path, "~") {
		home := os.Getenv("HOME")
		if home != "" {
			path = filepath.Join(home, path[1:])
		}
	}
	return os.ExpandEnv(path)
}

func getLanguages() map[string]LanguageDef {
	if loadedConfig == nil {
		LoadConfig()
	}

	userLangs := make(map[string]LanguageDef, len(loadedConfig.Languages))

	maps.Copy(userLangs, loadedConfig.Languages)

	return userLangs
}

// Optional helper for debugging
func PrintLoadedConfig() {
	if loadedConfig == nil {
		LoadConfig()
	}
	fmt.Printf("Loaded Config: %+v\n", *loadedConfig)
}

// Preset languages with default shebangs and comments
var PresetLanguages = map[string]LanguageDef{
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
func GetLanguages() map[string]LanguageDef {
	all := make(map[string]LanguageDef)

	// Start with built-in
	maps.Copy(all, PresetLanguages)

	// Override or extend with config
	maps.Copy(all, getLanguages())

	return all
}

func ensureConfigLoaded() {
	if loadedConfig == nil {
		LoadConfig()
	}
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
