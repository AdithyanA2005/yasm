package metadata

import (
	"bufio"
	"os"
	"strings"
	"yasm/common"
)

type Metadata struct {
	Title        string
	Description  string
	HasArguments string
	Usage        string
	Tags         []string
	Dependencies []string
}

func ExtractMetadata(filePath string, commentChar string) (Metadata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Metadata{}, err
	}
	defer file.Close()

	var md Metadata
	prefix := commentChar + " @yasm."

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, prefix) {
			continue
		}

		// Strip comment prefix
		line = strings.TrimPrefix(line, prefix)

		switch {
		case strings.HasPrefix(line, "title "):
			md.Title = strings.TrimSpace(strings.TrimPrefix(line, "title "))
		case strings.HasPrefix(line, "description "):
			md.Description = strings.TrimSpace(strings.TrimPrefix(line, "description "))
		case strings.HasPrefix(line, "has_arguments "):
			md.HasArguments = strings.TrimSpace(strings.TrimPrefix(line, "has_arguments "))
		case strings.HasPrefix(line, "usage "):
			usageLine := strings.TrimSpace(strings.TrimPrefix(line, "usage "))
			md.Usage = strings.Trim(usageLine, `"`)
		case strings.HasPrefix(line, "tags "):
			md.Tags = parseList(strings.TrimSpace(strings.TrimPrefix(line, "tags ")))
		case strings.HasPrefix(line, "dependencies "):
			md.Dependencies = parseList(strings.TrimSpace(strings.TrimPrefix(line, "dependencies ")))
		}
	}

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

func ExtractCommentChar(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", nil // empty file
	}
	firstLine := scanner.Text()

	for _, lang := range common.GetLanguages() {
		if strings.TrimSpace(firstLine) == strings.TrimSpace(lang.Shebang) {
			return lang.Comment, nil
		}
	}

	// fallback
	return "#", nil
}
