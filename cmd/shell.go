package cmd

import (
	"fmt"
	"os"
)

// PrintShellWrapper prints a shell function for interactive command injection
func generateShellWrapper(shell string) {
	switch shell {
	case "bash":
		fmt.Println(`yasm() {
    local CMD
    CMD=$(command yasm) || return
    READLINE_LINE="$CMD"
    READLINE_POINT=${#READLINE_LINE}
}`)
	case "zsh":
		fmt.Println(`yasm() {
    local CMD
    CMD=$(command yasm) || return
    print -z -- "$CMD"
}`)
	case "fish":
		fmt.Println(`function yasm
    set cmd (command yasm)
    if test -n "$cmd"
        commandline --replace "$cmd"
        commandline --function repaint
    end
end`)
	default:
		fmt.Fprintf(os.Stderr, "Unsupported shell: %s\n", shell)
		os.Exit(1)
	}
	os.Exit(0)
}
