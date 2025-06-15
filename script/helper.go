package script

import (
	"slices"
	"strings"

	"github.com/adithyana2005/yasm/usercfg"
)

// getLanguageByShebang returns the language name, its definition, and a boolean indicating
// whether a matching language was found for the given shebang line.
// It compares the shebang string against the Shebang field of each language definition.
func getLanguageByShebang(shebang string) (lang string, def usercfg.LanguageDef, found bool) {
	for name, def := range usercfg.GetLanguages() {
		if slices.Contains(def.Shebang, strings.TrimSpace(shebang)) {
			return name, def, true
		}
	}

	return "", usercfg.LanguageDef{}, false
}
