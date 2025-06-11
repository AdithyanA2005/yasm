package config

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

type LanguageDef struct {
	Shebang string `toml:"shebang"`
	Comment string `toml:"comment"`
}

type Config struct {
	ScriptsDir       string                 `toml:"scripts-dir"`
	Editor           string                 `toml:"editor"`
	AddScriptsToPath bool                   `toml:"add-scripts-to-path"`
	Languages        map[string]LanguageDef `toml:"languages"`
}

var loadedConfig *Config

// LoadConfig looks for config.toml in standard locations and loads it
func LoadConfig() {
	configPaths := []string{
		filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "yasm", "config.toml"),
		filepath.Join(os.Getenv("HOME"), ".config", "yasm", "config.toml"),
		filepath.Join(os.Getenv("HOME"), "yasm", "config.toml"),
	}

	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			cfg, err := parseConfigFile(path)
			if err == nil {
				loadedConfig = cfg
				return
			}
			fmt.Fprintf(os.Stderr, "Failed to parse config file %s: %v\n", path, err)
		}
	}

	// No config found, use empty defaults
	loadedConfig = &Config{}
}

func parseConfigFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetEditor() string {
	if loadedConfig == nil {
		LoadConfig()
	}

	if loadedConfig.Editor != "" {
		return loadedConfig.Editor
	}

	editor := os.Getenv("EDITOR")
	if editor != "" {
		return editor
	}

	return "nano"
}

func GetScriptsDir() string {
	if loadedConfig == nil {
		LoadConfig()
	}

	sd := loadedConfig.ScriptsDir
	if sd != "" {
		return expandPath(sd)
	}

	return filepath.Join(os.Getenv("HOME"), "scripts")
}

func IsAddScriptsToPathEnabled() bool {
	if loadedConfig == nil {
		LoadConfig()
	}
	return loadedConfig.AddScriptsToPath
}

// expandPath replaces ~ with full home path
func expandPath(path string) string {
	if strings.HasPrefix(path, "~") {
		home := os.Getenv("HOME")
		if home != "" {
			path = filepath.Join(home, path[1:])
		}
	}
	return os.ExpandEnv(path)
}

func GetLanguages() map[string]LanguageDef {
	if loadedConfig == nil {
		LoadConfig()
	}

	userLangs := make(map[string]LanguageDef, len(loadedConfig.Languages))

	maps.Copy(userLangs, loadedConfig.Languages)

	return userLangs
}

// Optional helper for debugging
func PrintLoadedConfig() {
	if loadedConfig == nil {
		LoadConfig()
	}
	fmt.Printf("Loaded Config: %+v\n", *loadedConfig)
}
