package fzf

import (
	"os"
	"os/exec"
	"strings"
)

// FuzzySelectScript presents a list of scripts to the user using the fzf fuzzy finder.
// It returns the selected script as a string, or an empty string and nil error if the user cancels.
// If another error occurs while running fzf, it returns a non-nil error.
func FuzzySelectScript(scripts []string) (string, error) {
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
