package shell

type shellDefinition struct {
	Name   string
	Script string
}

var SupportedShells = map[string]shellDefinition{
	"bash": {Name: "bash", Script: bashScript},
	"zsh":  {Name: "zsh", Script: zshScript},
	"fish": {Name: "fish", Script: fishScript},
}
const InjectPrefix = "::yasm-inject::"
