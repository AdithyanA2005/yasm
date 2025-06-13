package shell

import (
	"fmt"
	"strings"
)

// loadShellScript returns the integration script for the specified shell,
// or an error if unsupported.
func loadShellScript(shell string) (string, error) {
	// Get shell definition if it's a supported shell
	shellScript, ok := SupportedShells[shell]
	if !ok {
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}

	return sanitize(shellScript), nil
}

// sanitize replaces the placeholders in the shell script with the actual values.
// TODO: Implement this
func sanitize(input string) string {
	input = strings.ReplaceAll(input, "{{}}", "")
	return input
}
