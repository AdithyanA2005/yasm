package usercfg

var (
	loadedConfig *ConfigDef

	// Users env variables
	UserHomeDir   string
	UserConfigDir string
	UserEditor    string

	// Default values
	DefaultEditor     = "nano"
	DefaultScriptsDir = "scripts"
	PresetLanguages   = map[string]LanguageDef{
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
)

type LanguageDef struct {
	Shebang string `toml:"shebang"`
	Comment string `toml:"comment"`
}

type ConfigDef struct {
	ScriptsDir       string                 `toml:"scripts-dir"`
	Editor           string                 `toml:"editor"`
	AddScriptsToPath bool                   `toml:"add-scripts-to-path"`
	Languages        map[string]LanguageDef `toml:"languages"`
}
