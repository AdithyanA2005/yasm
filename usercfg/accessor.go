package usercfg

import (
	"path/filepath"
)

// GetEditor returns the preferred text editor.
// It checks the following, in order:
// 1. The editor specified in the loaded configuration.
// 2. The EDITOR environment variable.
// 3. Defaults to "nano" if neither is set.
func GetEditor() string {
	ensureConfigLoaded()

	if loadedConfig.Editor != "" {
		return loadedConfig.Editor
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

	if scriptsDir := loadedConfig.ScriptsDir; scriptsDir != "" {
		return expandPath(scriptsDir)
	} else {
		return filepath.Join(UserHomeDir, DefaultScriptsDir)
	}
}

// ShouldAddScriptsToPath returns true if the user configuration
// specifies that the scripts directory should be added to the PATH.
func ShouldAddScriptsToPath() bool {
	ensureConfigLoaded()

	return loadedConfig.AddScriptsToPath
}
