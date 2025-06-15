package main

import (
	"fmt"
	"os"

	"github.com/adithyana2005/yasm/cmd"
)

func main() {
	err := cmd.Execute(os.Args)
	// If any error occurs print it
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
