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

	shebang, ok := shebangs[lang]
	if !ok {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	scriptsDir := config.GetScriptsDir()
	if err := os.MkdirAll(scriptsDir, 0755); err != nil {
		return fmt.Errorf("failed to create scripts directory: %w", err)
	}

	// Add extension based on language (optional)
	// For now, use raw name
	scriptPath := filepath.Join(scriptsDir, name)

	if _, err := os.Stat(scriptPath); err == nil {
		return fmt.Errorf("script '%s' already exists", name)
	}

	// Pre-fill file with shebang
	file, err := os.OpenFile(scriptPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("failed to create script: %w", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s\n\n", shebang)
	if err != nil {
		return fmt.Errorf("failed to write shebang: %w", err)
	}

	// Open in editor
	editor := config.GetEditor()
	cmd := exec.Command(editor, scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
