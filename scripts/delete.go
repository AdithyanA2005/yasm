package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"yasm/config"
	"yasm/fzf"
)

func DeleteScript(scriptName string) error {
	scriptsDir := config.GetScriptsDir()
	// var err error

	if scriptName == "" {
		entries, err := os.ReadDir(scriptsDir)
		if err != nil {
			return fmt.Errorf("failed to list scripts: %w", err)
		}
		var scriptNames []string
		for _, entry := range entries {
			if entry.Type().IsRegular() {
				scriptNames = append(scriptNames, entry.Name())
			}
		}
		scriptName, err = fzf.FuzzySelectScript(scriptNames)
		if err != nil {
			return err
		}
		if scriptName == "" {
			return nil
		}
	}

	scriptPath := filepath.Join(scriptsDir, scriptName)
	fmt.Printf("Are you sure you want to delete '%s'? (y/N): ", scriptName)
	var response string
	fmt.Scanln(&response)
	if response != "y" && response != "Y" {
		fmt.Println("Aborted.")
		return nil
	}

	if err := os.Remove(scriptPath); err != nil {
		return fmt.Errorf("failed to delete script: %w", err)
	}

	fmt.Println("Deleted", scriptName)
	return nil
}
