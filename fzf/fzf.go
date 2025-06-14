package fzf

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"yasm/script"
	"yasm/usercfg"
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

// FuzzySelectScript presents a list of scripts to the user from which
// user can select one using the fzf fuzzy finder.
func FuzzySelectScript() (string, error) {
	// Get all the entries in the scripts directory
	scriptsDir := usercfg.GetScriptsDir()
	entries, err := os.ReadDir(scriptsDir)
	if err != nil {
		return "", fmt.Errorf("failed to list scripts: %w", err)
	}

	// If no entries found, print reason return (no error)
	if len(entries) == 0 {
		fmt.Println("No scripts found in", scriptsDir)
		return "", nil
	}

	// Collect mapping of displayName => actual filename
	displayMap := make(map[string]string)
	var displayList []string

	for _, entry := range entries {
		if entry.Type().IsRegular() {
			filename := entry.Name()
			filePath := filepath.Join(scriptsDir, filename)

			meta, err := script.ExtractMetadata(filePath)
			title := meta.Title
			if err != nil || title == "" {
				title = ""
			}

			displayName := filename
			if title != "" {
				displayName += " - " + title
			}

			displayList = append(displayList, displayName)
			displayMap[displayName] = filename
		}
	}

	// Run fzf
	selected, err := FuzzySelect(displayList)
	if err != nil {
		return "", fmt.Errorf("fzf exited: %w", err)
	}

	// Return actual filename
	return displayMap[selected], nil
}
