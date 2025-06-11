package fzf

import (
	"os"
	"os/exec"
	"strings"
)

// FuzzySelectScript runs fzf on a list of script names and returns the selected one.
func FuzzySelectScript(scripts []string) (string, error) {
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(scripts, "\n"))
	cmd.Stderr = os.Stderr

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

	selection := strings.TrimSpace(string(output))
	return selection, nil
}
