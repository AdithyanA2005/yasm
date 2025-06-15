package usercfg

import (
	"maps"
	"path/filepath"

	"github.com/adithyana2005/yasm/utils"
)

// GetEditor returns the preferred text editor.
// It checks the following, in order:
// 1. The editor specified in the loaded configuration.
// 2. The EDITOR environment variable.
// 3. Defaults to "nano" if neither is set.
func GetEditor() string {
	ensureConfigLoaded()

	if LoadedConfig.Editor != "" {
		return LoadedConfig.Editor
	} else if UserEditor != "" {
		return UserEditor
	} else {
		return DefaultEditor
	}
}

// GetScriptsDir returns the directory path where user scripts are stored.
// It checks the following, in order:
// 1. The ScriptsDir specified in the loaded configuration (after expanding any paths).
// 2. Defaults to "$HOME/{DefaultScriptsDirName}" if not set in the configuration.
func GetScriptsDir() string {
	ensureConfigLoaded()

	if scriptsDir := LoadedConfig.ScriptsDir; scriptsDir != "" {
		return utils.ExpandUserPath(scriptsDir)
	} else {
		return filepath.Join(UserHomeDir, DefaultScriptsDir)
	}
}

// ShouldAddScriptsToPath returns true if the user configuration
// specifies that the scripts directory should be added to the PATH.
func ShouldAddScriptsToPath() bool {
	ensureConfigLoaded()

	return LoadedConfig.AddScriptsToPath
}

// GetLanguages returns a map of all available language definitions.
// It merges built-in preset languages with any user-defined overrides or extensions
// from the loaded configuration. User-defined languages take precedence.
func GetLanguages() map[string]LanguageDef {
	ensureConfigLoaded()

	all := make(map[string]LanguageDef)
	maps.Copy(all, PresetLanguages)
	maps.Copy(all, LoadedConfig.Languages)

	return all
}
