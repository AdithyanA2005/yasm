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

	// Resolve symlinks
	resolvedPath, err := filepath.EvalSymlinks(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to resolve script path: %w", err)
	}

	// Check if the script exists
	info, err := os.Stat(resolvedPath)
	if err != nil {
		return fmt.Errorf("script not found: %w", err)
	}

	// Check if the script is a regular file
	if !info.Mode().IsRegular() {
		return fmt.Errorf("script is not a regular file: %s", scriptPath)
	}

	// Check if the script is executable
	if info.Mode().Perm()&0111 == 0 {
		return fmt.Errorf("script is not executable: %s", scriptPath)
	}

	// Extract needed metadata from script
	meta, err := script.ExtractMetadata(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to extract metadata: %w", err)
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
