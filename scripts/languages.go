package scripts

import "yasm/config"

var presetLanguages = map[string]config.LanguageDef{
	"bash": {
		Shebang: "#!/usr/bin/env bash",
		Comment: "#",
	},
	"python": {
		Shebang: "#!/usr/bin/env python3",
		Comment: "#",
	},
	"sh": {
		Shebang: "#!/bin/sh",
		Comment: "#",
	},
	"zsh": {
		Shebang: "#!/usr/bin/env zsh",
		Comment: "#",
	},
}

func GetLanguages() map[string]config.LanguageDef {
	all := make(map[string]config.LanguageDef)

	// Start with built-in
	for k, v := range presetLanguages {
		all[k] = v
	}

	// Override or extend with config
	for k, v := range config.GetLanguages() {
		all[k] = v
	}

	return all
}

// Supported shebangs by language
// var shebangs = map[string]string{
// 	"bash":   "#!/usr/bin/env bash",
// 	"python": "#!/usr/bin/env python3",
// 	"sh":     "#!/bin/sh",
// 	"zsh":    "#!/usr/bin/env zsh",
// }
//
// func ListSupportedLanguages() map[string]string {
// 	return shebangs
// }
