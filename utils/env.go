package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// ExpandUserPath expands environment variables and the ~ symbol in the given path.
// If the home directory cannot be resolved, it returns the original path.
func ExpandUserPath(path string) string {
	path = os.ExpandEnv(path)

	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path // fallback: return original if home can't be resolved
		}
		switch {
		case path == "~":
			return home
		case strings.HasPrefix(path, "~/"):
			return filepath.Join(home, path[2:])
		}
	}

	return path
}
