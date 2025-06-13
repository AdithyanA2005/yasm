package shell

// SupportedShells is a map of supported shell names to their embedded integration scripts.
var SupportedShells = map[string]string{
	"bash": bashScript,
	"zsh":  zshScript,
	"fish": fishScript,
}
