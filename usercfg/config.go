package usercfg

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/olekukonko/tablewriter"
	toml "github.com/pelletier/go-toml/v2"
)

// InitConfig initializes user directories and loads the configuration file.
// It exits if user directories can't be determined or a config file exists but can't be parsed.
// If no config file is found, it falls back to an empty config.
func InitConfig() {
	var err error

	// Get user's home directory
	UserHomeDir, err = os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not determine home directory:  %v\n", err)
		os.Exit(1)
	}

	// Get user's config directory
	UserConfigDir, err = os.UserConfigDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not determine config directory:  %v\n", err)
		os.Exit(1)
	}

	// Possible config file paths
	configPaths := []string{
		filepath.Join(UserConfigDir, "yasm", "config.toml"),
		filepath.Join(UserHomeDir, ".config", "yasm", "config.toml"),
		filepath.Join(UserHomeDir, "yasm", "config.toml"),
	}

	// Try to load config from the first valid file
	for _, path := range configPaths {
		if cfg, err := readTomlFile(path); err == nil {
			loadedConfig = cfg
			return
		} else if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Failed to load config from %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	// Fallback to empty config
	loadedConfig = &ConfigDef{}
}

// readTomlFile reads a TOML configuration file from the given path,
// unmarshals its contents into a ConfigDef struct, and returns a pointer to it.
// Returns an error if the file cannot be read or if unmarshalling fails.
func readTomlFile(path string) (*ConfigDef, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg ConfigDef
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ensureConfigLoaded() {
	if loadedConfig == nil {
		InitConfig()
	}
}

// PrintLanguagesTable prints a table of all supported languages to standard output.
// The table includes the language name, its shebang line, and the comment prefix.
// Languages are listed in sorted order for consistent output.
func PrintLanguagesTable() {
	langs := GetLanguages()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Language", "Shebang", "Comment"})

	// Sort keys for consistent output
	keys := make([]string, 0, len(langs))
	for lang := range langs {
		keys = append(keys, lang)
	}
	sort.Strings(keys)

	for _, lang := range keys {
		def := langs[lang]
		table.Append([]string{lang, def.Shebang, def.Comment})
	}

	table.Render()
}
