package shell

type shellDefinition struct {
	Name   string
	Script string
}

// SupportedShells is a map of supported shell names to their embeded integration scripts.
var SupportedShells = map[string]shellDefinition{
	"bash": {Name: "bash", Script: bashScript},
	"zsh":  {Name: "zsh", Script: zshScript},
	"fish": {Name: "fish", Script: fishScript},
}

// Special prefix that is used with commands to inject them in user's shell.
const InjectPrefix = "::yasm-inject::"
