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

func ExtractMetadata(filePath string) (Metadata, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return Metadata{}, err
	}
	defer file.Close()

	var md Metadata
	var prefix string

	lineCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		// Handling the first line of the script
		if lineCount == 1 {
			// Ensure that the script contains a shebang
			if !strings.HasPrefix(line, "#!") {
				return Metadata{}, cli.Exit("Error: Script does not have a shebang", 1)
			}

			// Ensure that the shebang is a of a supported language
			shebang := strings.TrimSpace(line)
			_, def, found := getLanguageByShebang(shebang)
			if !found {
				// TODO: Add some info about extending languages
				fmt.Fprintf(os.Stderr, "Error: No language config found for shebang '%s' in file '%s'\n", shebang, filePath)
				return Metadata{}, cli.Exit("", 1)
			}

			// Set the prefix for metadata lines (<comment> ScriptMetadataPrefix)
			prefix = fmt.Sprintf("%s %s", def.Comment, ScriptMetadataPrefix)
			continue
		}

		// Skip lines that dont start with the prefix
		if !strings.HasPrefix(line, prefix) {
			continue
		}

		// Define handlers for each metadata field
		handlers := map[string]func(string){
			prefix + ".title":        func(val string) { md.Title = val },
			prefix + ".description":  func(val string) { md.Description = val },
			prefix + ".tags":         func(val string) { md.Tags = parseList(val) },
			prefix + ".dependencies": func(val string) { md.Dependencies = parseList(val) },
		}

		// Extract needed metadata from the line
		for key, handler := range handlers {
			if strings.HasPrefix(line, key) {
				handler(strings.TrimSpace(strings.TrimPrefix(line, key)))
				break // Found match, skip rest
			}
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
