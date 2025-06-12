package edit

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"yasm/usercfg"
)

// EditScript opens the specified script in the user's configured editor.
// If scriptName is empty, it presents a fuzzy finder menu to select a script from the scripts directory.
// Returns an error if the script cannot be found, listed, or opened in the editor.
func editScript(scriptName string) error {
	scriptsDir := usercfg.GetScriptsDir()

	// Prepare and run the editor command to open the script.
	// The editor is launched with the script path, and its standard input,
	// output, and error are connected to the current process.
	scriptPath := filepath.Join(scriptsDir, scriptName)
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
