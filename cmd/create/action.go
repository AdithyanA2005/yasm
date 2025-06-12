package create

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"yasm/config"
	"yasm/script"
)

// createScript creates a new script file with the specified name and language.
// If the language is not provided, it defaults to "bash".
// The function checks for supported languages, creates the scripts directory if needed,
// ensures the script does not already exist, writes a template to the new file,
// and opens the script in the user's configured editor.
// Returns an error if any step fails.
func createScript(scriptName string, lang string) error {
	// Validate script name
	if err := isValidScriptName(scriptName); err != nil {
		return err
	}

	// If no language is specified, default to "bash"
	if lang == "" {
		lang = "bash"
	}

	// Get language definitions
	languages := script.GetLanguages()
	langDef, ok := languages[lang]
	if !ok {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	// Ensure the scripts directory exists
	scriptsDir := config.GetScriptsDir()
	if err := os.MkdirAll(scriptsDir, 0755); err != nil {
		return fmt.Errorf("failed to create scripts directory: %w", err)
	}

	// Get the path for the new script
	scriptPath := filepath.Join(scriptsDir, scriptName)

	// Check if script already exists
	if _, err := os.Stat(scriptPath); err == nil {
		return fmt.Errorf("script '%s' already exists", scriptName)
	}

	// Open the script file for writing
	file, err := os.OpenFile(scriptPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("failed to create script: %w", err)
	}
	defer file.Close()

	// Generate the script template
	template := generateScriptTemplate(langDef.Shebang, langDef.Comment, scriptName)

	// Write the template to the script file
	_, err = file.WriteString(template)
	if err != nil {
		return fmt.Errorf("failed to write to script: %w", err)
	}

	// Open in editor
	editor := config.GetEditor()
	cmd := exec.Command(editor, scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
