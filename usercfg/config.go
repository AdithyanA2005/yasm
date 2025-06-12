package usercfg

import (
	"fmt"
	"os"
	"path/filepath"

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
			LoadedConfig = cfg
			return
		} else if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Failed to load config from %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	// Fallback to empty config
	LoadedConfig = &ConfigDef{}
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
