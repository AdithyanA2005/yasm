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

		// Extract needed metadata from the line
		switch {
		case strings.HasPrefix(line, prefix+".title "):
			md.Title = strings.TrimSpace(strings.TrimPrefix(line, prefix+".title "))
		case strings.HasPrefix(line, prefix+".description "):
			md.Description = strings.TrimSpace(strings.TrimPrefix(line, prefix+".description "))
		case strings.HasPrefix(line, prefix+".tags "):
			md.Tags = parseList(strings.TrimSpace(strings.TrimPrefix(line, prefix+".tags ")))
		case strings.HasPrefix(line, prefix+".dependencies "):
			md.Dependencies = parseList(strings.TrimSpace(strings.TrimPrefix(line, prefix+".dependencies ")))
		}
	}

	// If there occured it will sent it, else it will return nil for error
	return md, scanner.Err()
}

func parseList(raw string) []string {
	trimmed := strings.Trim(raw, "[]")
	if trimmed == "" {
		return nil
	}
	parts := strings.Split(trimmed, ",")
	for i, p := range parts {
		parts[i] = strings.Trim(strings.Trim(p, `"`), " ")
	}
	return parts
}
