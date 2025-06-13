package shell

import (
	"fmt"
)

// getSupportedShells returns a slice containing the names of all supported shells.
// The order of the shells is not guaranteed.
func getSupportedShells() []string {
	keys := make([]string, 0, len(SupportedShells))
	for k := range SupportedShells {
		keys = append(keys, k)
	}
	return keys
}

// loadShellScript returns the integration script for the specified shell,
// or an error if unsupported.
func loadShellScript(shell string) (string, error) {
	// Get shell definition if it's a supported shell
	def, ok := SupportedShells[shell]
	if !ok {
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}

	return def.Script, nil
}
