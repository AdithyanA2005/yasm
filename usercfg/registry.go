package usercfg

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
