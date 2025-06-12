package usercfg

var (
	loadedConfig *ConfigDef

	// Users env variables
	UserHomeDir   string
	UserConfigDir string
	UserEditor    string
)

const (
	// Default values
	DefaultEditor     = "nano"
	DefaultScriptsDir = "scripts"
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
