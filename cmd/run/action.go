package run

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"yasm/script"
	"yasm/usercfg"
)

// runScript loads a script's metadata, checks for unmet dependencies, and prints
// a message indicating the script is ready to run. Returns an error if metadata
// extraction fails or dependencies are missing.
func runScript(scriptName string, scriptArgs []string) error {
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

	// Execute the script with arguments
	cmd := exec.Command(scriptPath, scriptArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	err = cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// Don't assume Unix only — use ExitCode()
			return fmt.Errorf("script exited with code %d", exitError.ExitCode())
		}
		return fmt.Errorf("failed to run script: %w", err)
	}

	return nil
}
