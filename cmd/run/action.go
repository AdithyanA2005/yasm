package run

import (
	"fmt"
	"path/filepath"
	"strings"
	"yasm/script"
	"yasm/usercfg"
)

// runScript loads a script's metadata, checks for unmet dependencies, and prints
// a message indicating the script is ready to run. Returns an error if metadata
// extraction fails or dependencies are missing.
func runScript(scriptName string) error {
	// Get the metadata for the script file
	scriptsDir := usercfg.GetScriptsDir()
	scriptPath := filepath.Join(scriptsDir, scriptName)
	meta, err := script.ExtractMetadata(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to run script: %w", err)
	}

	// Check if dependencies are installed
	canExecute, missing := isDepsInstalled(meta.Dependencies)
	if !canExecute {
		return fmt.Errorf("unmet dependencies: %s", strings.Join(missing, ", "))
	}

	if scriptName != "" {
		fmt.Printf("%s", scriptName)
	}

	return nil
}
