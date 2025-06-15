package edit

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/adithyana2005/yasm/usercfg"
)

// EditScript opens the specified script in the user's configured editor.
// If scriptName is empty, it presents a fuzzy finder menu to select a script from the scripts directory.
// Returns an error if the script cannot be found, listed, or opened in the editor.
func editScript(scriptName string) error {
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

	// Prepare and run the editor command to open the script.
	editor := usercfg.GetEditor()
	cmd := exec.Command(editor, scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to open editor: %w", err)
	}

	return nil
}
