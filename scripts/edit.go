package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"yasm/config"
	"yasm/fzf"
)

func EditScript(scriptName string) error {
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
	editor := config.GetEditor()
	cmd := exec.Command(editor, scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to open editor: %w", err)
	}

	return nil
}
