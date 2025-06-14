package usercfg

var (
	LoadedConfig *ConfigDef

	// Users env variables
	UserHomeDir   string
	UserConfigDir string
	UserEditor    string

	// Default values
	DefaultEditor     = "nano"
	DefaultScriptsDir = "scripts"
	PresetLanguages   = map[string]LanguageDef{
		"bash": {
			Shebang: []string{"#!/usr/bin/env bash", "#!/usr/bin/bash"},
			Comment: "#",
		},
		"python": {
			Shebang: []string{"#!/usr/bin/env python3"},
			Comment: "#",
		},
		"sh": {
			Shebang: []string{"#!/bin/sh"},
			Comment: "#",
		},
		"zsh": {
			Shebang: []string{"#!/usr/bin/env zsh"},
			Comment: "#",
		},
	}
)

type LanguageDef struct {
	Shebang []string `toml:"shebang"`
	Comment string   `toml:"comment"`
}

type ConfigDef struct {
	ScriptsDir       string                 `toml:"scripts-dir"`
	Editor           string                 `toml:"editor"`
	AddScriptsToPath bool                   `toml:"add-scripts-to-path"`
	Languages        map[string]LanguageDef `toml:"languages"`
}
