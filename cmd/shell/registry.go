package shell

import (
	_ "embed"
)

type shellDefinition struct {
	Name   string
	Script string
}

//go:embed integrations/yasm.bash
var bashScript string

//go:embed integrations/yasm.zsh
var zshScript string

//go:embed integrations/yasm.fish
var fishScript string

var SupportedShells = map[string]shellDefinition{
	"bash": {Name: "bash", Script: bashScript},
	"zsh":  {Name: "zsh", Script: zshScript},
	"fish": {Name: "fish", Script: fishScript},
}
