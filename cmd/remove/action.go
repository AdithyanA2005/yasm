package remove

import (
	"fmt"
	"os"
	"path/filepath"
	"yasm/config"
	"yasm/fzf"
	"yasm/utils"
)

// DeleteScript deletes the specified script from the scripts directory.
// If scriptName is empty, it presents a fuzzy finder menu to select a script to delete.
// The function prompts the user for confirmation before deletion.
// Returns an error if the script cannot be listed, selected, or deleted.
func removeScript(scriptName string) error {
	scriptsDir := config.GetScriptsDir()

	// If scriptname is not given use fuzzy finder to select a script
	if scriptName == "" {
		var err error

		// Fuzzy select a script if no name provided
		scriptName, err = fzf.FuzzySelectScript()
		if err != nil {
			return err
		}

		// If no script was selected, return nil (no error)
		if scriptName == "" {
			return nil
		}
	}

	// Take user confirmation before deleting the script
	scriptPath := filepath.Join(scriptsDir, scriptName)
	isConfirmed := utils.Confirm("Are you sure you want to delete `"+scriptName+"`?", "n")

	// Abort if not confirmed
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
