package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"yasm/config"
)

// CreateScript creates a new script file with the specified name and language.
// If the language is not provided, it defaults to "bash".
// The function checks for supported languages, creates the scripts directory if needed,
// ensures the script does not already exist, writes a template to the new file,
// and opens the script in the user's configured editor.
// Returns an error if any step fails.
func CreateScript(name string, lang string) error {
	// Validate script name
	// TODO: Add additional checks
	if name == "" {
		return fmt.Errorf("script name cannot be empty")
	}

	// If no language is specified, default to "bash"
	if lang == "" {
		lang = "bash"
	}

	// Get language definitions
	languages := GetLanguages()
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
	scriptPath := filepath.Join(scriptsDir, name)

	// Check if script already exists
	if _, err := os.Stat(scriptPath); err == nil {
		return fmt.Errorf("script '%s' already exists", name)
	}

	// Open the script file for writing
	file, err := os.OpenFile(scriptPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("failed to create script: %w", err)
	}
	defer file.Close()

	// Generate the script template
	template := generateScript(langDef.Shebang, langDef.Comment, name)

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

// generateScript returns a template string for a new script file.
// It includes a shebang, YASM config block, and a placeholder for the main script.
// shebang: the interpreter directive (e.g., #!/bin/bash)
// comment: the comment prefix for the language (e.g., #, //)
// name: the script name to include in the config block
func generateScript(shebang, comment, name string) string {
	return fmt.Sprintf(`%[1]s

%[2]s YASM CONFIG START >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
%[2]s @yasm.title %[3]s
%[2]s @yasm.description 
%[2]s @yasm.tags []
%[2]s @yasm.dependencies []
%[2]s <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< YASM CONFIG END

%[2]s ################################################
%[2]s ### MAIN SCRIPT STARTS #########################
%[2]s ################################################

%[2]s Write your script here.
`, shebang, comment, name)
}
