package scripts

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"unicode/utf8"
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
	// Validate script name
	if err := isValidScriptName(name); err != nil {
		return err
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

// isValidScriptName validates the provided script name according to several rules:
// - Name must not be empty.
// - Name must not be "." or "..".
// - Name must not exceed 255 characters.
// - Name must not start or end with a space or dot.
// - Name must not contain any of the following characters: <>:"/\|?*
// - Name must not contain control characters (ASCII < 32).
// - Name must not be a reserved Windows device name (case-insensitive).
// Returns an error describing the first validation failure, or nil if valid.
func isValidScriptName(name string) error {
	if name == "" {
		return errors.New("script name cannot be empty")
	}

	if name == "." || name == ".." {
		return fmt.Errorf("script name cannot be '.' or '..'")
	}

	if utf8.RuneCountInString(name) > 255 {
		return fmt.Errorf("script name cannot exceed 255 characters")
	}

	// Disallow leading/trailing spaces or dots
	if strings.HasPrefix(name, " ") || strings.HasSuffix(name, " ") ||
		strings.HasPrefix(name, ".") || strings.HasSuffix(name, ".") {
		return fmt.Errorf("script name cannot start or end with a space or dot")
	}

	// Invalid characters (common restrictions across platforms)
	if strings.ContainsAny(name, `<>:"/\|?*`) {
		return fmt.Errorf("script name contains invalid characters: <>:\"/\\|?*")
	}

	// Control characters
	for _, r := range name {
		if r < 32 {
			return fmt.Errorf("script name contains control characters")
		}
	}

	// Reserved Windows device names (case-insensitive)
	reserved := []string{
		"CON", "PRN", "AUX", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
	}
	upperName := strings.ToUpper(name)
	baseName := strings.Split(upperName, ".")[0]
	if slices.Contains(reserved, baseName) {
		return fmt.Errorf("script name '%s' is a reserved name on Windows", baseName)
	}

	return nil
}
