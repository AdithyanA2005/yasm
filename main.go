package main

import (
	"log"
	"os"
	"yasm/cmd"
)

func main() {
	err := cmd.Execute(os.Args)
	// If any error occurs print it
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}
