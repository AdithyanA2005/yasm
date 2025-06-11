package main

import (
	"fmt"
	"os"
	"yasm/cmd"
)

func main() {
	err := cmd.Execute(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
