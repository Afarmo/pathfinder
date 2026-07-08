package main

import (
	"fmt"
	"os"
	"pathfinder/internal/cli"
)

func main() {
	cfg, err := cli.ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Parsing Arguments: %v", err)
		return
	}
}
