package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"yasm/config"
	"yasm/fzf"
)

// DeleteScript deletes the specified script from the scripts directory.
// If scriptName is empty, it presents a fuzzy finder menu to select a script to delete.
// The function prompts the user for confirmation before deletion.
// Returns an error if the script cannot be listed, selected, or deleted.
func DeleteScript(scriptName string) error {
	scriptsDir := config.GetScriptsDir()

	// If scriptname is not given use fuzzy finder to select a script
	if scriptName == "" {
		// Read entries in the scripts directory
		entries, err := os.ReadDir(scriptsDir)
		if err != nil {
			return fmt.Errorf("failed to list scripts: %w", err)
		}

		// Collect the names of all regular files from the entries slice.
		var scriptNames []string
		for _, entry := range entries {
			if entry.Type().IsRegular() {
				scriptNames = append(scriptNames, entry.Name())
			}
		}

		// Fuzzy select a script if no name provided
		scriptName, err = fzf.FuzzySelectScript(scriptNames)
		if err != nil {
			return err
		}

		// If no script was selected, return nil (no error)
		if scriptName == "" {
			return nil
		}
	}

	// Take user confirmation before deleting the script
	// TODO: Make Confirm function capable for handling this (default values)
	scriptPath := filepath.Join(scriptsDir, scriptName)
	fmt.Printf("Are you sure you want to delete '%s'? (y/N): ", scriptName)
	var response string
	fmt.Scanln(&response)
	if response != "y" && response != "Y" {
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
