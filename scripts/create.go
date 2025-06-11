package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"yasm/config"
)

func CreateScript(name string, lang string) error {
	if name == "" {
		return fmt.Errorf("script name cannot be empty")
	}

	if lang == "" {
		lang = "bash"
	}

	// Get language definitions
	languages := GetLanguages()
	langDef, ok := languages[lang]
	if !ok {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	scriptsDir := config.GetScriptsDir()
	if err := os.MkdirAll(scriptsDir, 0755); err != nil {
		return fmt.Errorf("failed to create scripts directory: %w", err)
	}

	// Use raw name to create script file
	scriptPath := filepath.Join(scriptsDir, name)

	// Check if script already exists
	if _, err := os.Stat(scriptPath); err == nil {
		return fmt.Errorf("script '%s' already exists", name)
	}

	// Create and write template content
	file, err := os.OpenFile(scriptPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("failed to create script: %w", err)
	}
	defer file.Close()

	template := fmt.Sprintf(
		"%s\n%s %s\n%s %s\n",
		langDef.Shebang,
		langDef.Comment, "Script created by yasm",
		langDef.Comment, "Add your code below",
	)

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
