package remove

import (
	"fmt"
	"os"
	"path/filepath"
	"yasm/usercfg"
	"yasm/utils"
)

// DeleteScript deletes the specified script from the scripts directory.
// If scriptName is empty, it presents a fuzzy finder menu to select a script to delete.
// The function prompts the user for confirmation before deletion.
// Returns an error if the script cannot be listed, selected, or deleted.
func removeScript(scriptName string) error {
	scriptsDir := usercfg.GetScriptsDir()
	scriptPath := filepath.Join(scriptsDir, scriptName)

	// Check if the file exists and is not a directory
	info, err := os.Stat(scriptPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("script not found: %s", scriptPath)
		}
		return fmt.Errorf("error checking script file: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("expected a file but found a directory: %s", scriptPath)
	}

	// Proceed only if user confirms before deleting the script
	isConfirmed := utils.Confirm("Are you sure you want to delete `"+scriptName+"`?", "n")
	if !isConfirmed {
		fmt.Println("Aborted.")
		return nil
	}

	// If user confirmed, delete the script
	if err := os.Remove(scriptPath); err != nil {
		return fmt.Errorf("failed to delete script: %w", err)
	}

	fmt.Println("Deleted", scriptName)
	return nil
}
