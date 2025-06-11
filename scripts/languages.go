package scripts

// Supported shebangs by language
var shebangs = map[string]string{
	"bash":   "#!/usr/bin/env bash",
	"python": "#!/usr/bin/env python3",
	"sh":     "#!/bin/sh",
	"zsh":    "#!/usr/bin/env zsh",
}

func ListSupportedLanguages() map[string]string {
	return shebangs
}
