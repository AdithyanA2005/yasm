package script

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adithyana2005/yasm/utils"
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
	}

	type Target struct {
		Key     string
		Handler func(string)
	}

	// Map metadata keys to targets that set fields on the Metadata struct.
	targets := []Target{
		{Key: ScriptMetadataPrefix + ".title", Handler: func(val string) { md.Title = val }},
		{Key: ScriptMetadataPrefix + ".description", Handler: func(val string) { md.Description = val }},
		{Key: ScriptMetadataPrefix + ".tags", Handler: func(val string) { md.Tags = parseList(val) }},
		{Key: ScriptMetadataPrefix + ".dependencies", Handler: func(val string) { md.Dependencies = parseList(val) }},
	}

	// Scan each line after the shebang to extract metadata.
	// Stop when reaching the first non-comment, non-empty line (script body).
lineLoop:
	for scanner.Scan() {
		trimmedLine := strings.TrimSpace(scanner.Text())

		// If the line is a metadata line, process it using the registered handlers.
		for _, target := range targets {
			if value, found := utils.ExtractAfterSubstring(trimmedLine, target.Key); found {
				target.Handler(strings.TrimSpace(value))
				continue lineLoop // Continue to the next line after processing
			}
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
