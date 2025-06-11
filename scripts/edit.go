package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"yasm/common"
	"yasm/config"
)

// EditScript opens the specified script in the user's configured editor.
// If scriptName is empty, it presents a fuzzy finder menu to select a script from the scripts directory.
// Returns an error if the script cannot be found, listed, or opened in the editor.
func EditScript(scriptName string) error {
	scriptsDir := config.GetScriptsDir()

	// If scriptname is not given use fuzzy finder to select a script
	if scriptName == "" {
		var err error

		// Fuzzy select a script if no name provided
		scriptName, err = common.FuzzySelectScript()
		if err != nil {
			return err
		}

		// If no script was selected, return nil (no error)
		if scriptName == "" {
			return nil
		}
	}

	// Prepare and run the editor command to open the script.
	// The editor is launched with the script path, and its standard input,
	// output, and error are connected to the current process.
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
