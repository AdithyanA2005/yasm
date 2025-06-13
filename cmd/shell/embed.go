package shell

import (
	_ "embed"
)

//go:embed templates/template.bash
var bashScript string

//go:embed templates/template.zsh
var zshScript string

//go:embed templates/template.fish
var fishScript string
