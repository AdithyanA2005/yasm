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

// InitConfig initializes user directories and loads the configuration file.
// It exits if user directories can't be determined or a config file exists but can't be parsed.
// If no config file is found, it falls back to an empty config.
func InitConfig() {
	var err error

	// Get user's home directory
	UserHomeDir, err = os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not determine home directory:  %v\n", err)
		os.Exit(1)
	}

	// Get user's config directory
	UserConfigDir, err = os.UserConfigDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not determine config directory:  %v\n", err)
		os.Exit(1)
	}

	// Possible config file paths
	configPaths := []string{
		filepath.Join(UserConfigDir, "yasm", "config.toml"),
		filepath.Join(UserHomeDir, ".config", "yasm", "config.toml"),
		filepath.Join(UserHomeDir, "yasm", "config.toml"),
	}

	// Try to load config from the first valid file
	for _, path := range configPaths {
		if cfg, err := readTomlFile(path); err == nil {
			loadedConfig = cfg
			return
		} else if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Failed to load config from %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	// Fallback to empty config
	loadedConfig = &ConfigDef{}
}

// readTomlFile reads a TOML configuration file from the given path,
// unmarshals its contents into a ConfigDef struct, and returns a pointer to it.
// Returns an error if the file cannot be read or if unmarshalling fails.
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
		InitConfig()
	}

	userLangs := make(map[string]LanguageDef, len(loadedConfig.Languages))

	maps.Copy(userLangs, loadedConfig.Languages)

	return userLangs
}

// Optional helper for debugging
func PrintLoadedConfig() {
	if loadedConfig == nil {
		InitConfig()
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
		InitConfig()
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
