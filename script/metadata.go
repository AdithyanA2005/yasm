package script

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

type Metadata struct {
	Title        string
	Description  string
	Tags         []string
	Dependencies []string
}

// ExtractMetadata reads a script file, determines its language from the shebang,
// and extracts metadata fields (title, description, tags, dependencies) from
// specially formatted comment lines at the top of the file.
//
// Metadata lines must follow the format:
//
//	<comment_char> script.<field>: <value>
//
// Returns a Metadata struct populated with extracted values, or an error.
func ExtractMetadata(filePath string) (Metadata, error) {
	// Open the script file for reading.
	file, err := os.Open(filePath)
	if err != nil {
		return Metadata{}, err
	}
	defer file.Close()

	var md Metadata
	var prefix string
	var commentChar string

	scanner := bufio.NewScanner(file)

	// Read the first line to check for a shebang and determine script language.
	if scanner.Scan() {
		trimmedLine := strings.TrimSpace(scanner.Text())

		// Ensure that the script contains a shebang (#!).
		if !strings.HasPrefix(trimmedLine, "#!") {
			return Metadata{}, cli.Exit("Error: Script does not have a shebang", 1)
		}

		// Determine language config from the shebang.
		_, def, found := getLanguageByShebang(trimmedLine)
		if !found {
			// Print error if language is not supported.
			fmt.Fprintf(os.Stderr, "Error: No language config found for shebang '%s' in file '%s'\n", trimmedLine, filePath)
			return Metadata{}, cli.Exit("", 1)
		}

		commentChar = def.Comment
		// Build the metadata prefix, e.g. "# @yasm."
		prefix = fmt.Sprintf("%s %s", commentChar, ScriptMetadataPrefix)
	}

	// Map metadata keys to handlers that set fields on the Metadata struct.
	handlers := map[string]func(string){
		prefix + ".title":        func(val string) { md.Title = val },
		prefix + ".description":  func(val string) { md.Description = val },
		prefix + ".tags":         func(val string) { md.Tags = parseList(val) },
		prefix + ".dependencies": func(val string) { md.Dependencies = parseList(val) },
	}

	// Scan each line after the shebang to extract metadata.
	// Stop when reaching the first non-comment, non-empty line (script body).
	for scanner.Scan() {
		trimmedLine := strings.TrimSpace(scanner.Text())

		// If the line is a metadata line, process it using the registered handlers.
		if strings.HasPrefix(trimmedLine, prefix) {
			// Call the handler for the matching metadata key.
			for key, handler := range handlers {
				if strings.HasPrefix(trimmedLine, key) {
					handler(strings.TrimSpace(strings.TrimPrefix(trimmedLine, key)))
					break
				}
			}
			continue // Continue scanning for more metadata lines.
		}

		// If the line is not empty/whitespace and not a comment,
		// it marks the start of the actual script content. Exit the loop.
		if trimmedLine != "" && !strings.HasPrefix(trimmedLine, commentChar) {
			break
		}
	}

	return md, scanner.Err()
}

// parseList splits a raw string into a slice of strings using whitespace as the delimiter.
// It returns a slice containing each field found in the input string.
func parseList(raw string) []string {
	// Used Fields instead of Split to handle multiple spaces issue
	return strings.Fields(raw)
}
