package shell

import (
	_ "embed"
)

//go:embed integrations/template.bash
var bashScript string

//go:embed integrations/template.zsh
var zshScript string

//go:embed integrations/template.fish
var fishScript string
