package fzf

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"yasm/config"
)

// FuzzySelect presents a list of entries to the user using the fzf fuzzy finder.
// It returns the selected entry as a string, or an empty string and nil error if the user cancels.
// If another error occurs while running fzf, it returns a non-nil error.
func FuzzySelect(scripts []string) (string, error) {
	// Prepare the fzf command with the list of scripts
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(scripts, "\n"))
	cmd.Stderr = os.Stderr

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		// If user cancels, fzf returns non-zero
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 130 {
				// Ctrl+C or Esc — expected user cancel
				return "", nil
			}
		}
		return "", err
	}

	// Trim the output to get the selected script
	selection := strings.TrimSpace(string(output))
	return selection, nil
}

// FuzzySelectScript presents a list of scripts to the user from which he can select one using the fzf fuzzy finder.
func FuzzySelectScript() (string, error) {
	// Read scripts directory and store entries sorted by name
	scriptsDir := config.GetScriptsDir()
	entries, err := os.ReadDir(scriptsDir)
	if err != nil {
		return "", fmt.Errorf("failed to list scripts: %w", err)
	}

	// Collect the names of all regular files from the entries slice.
	// Only files (not directories or other types) are included in the scriptNames slice.
	var scriptNames []string
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			scriptNames = append(scriptNames, entry.Name())
		}
	}

	// If no scripts found, print message and exit
	if len(scriptNames) == 0 {
		fmt.Println("No scripts found in", scriptsDir)
		return "", nil
	}

	// Run fzf to select a script
	scriptName, err := FuzzySelect(scriptNames)
	if err != nil {
		return "", fmt.Errorf("fzf exited: %w", err)
	}

	return scriptName, nil
}
