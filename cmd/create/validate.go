package create

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"
)

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
