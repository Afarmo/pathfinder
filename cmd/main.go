package main

import (
	"fmt"
	"os"
	"pathfinder/internal/cli"
	"pathfinder/internal/parser"
)

func main() {
	cfg, err := cli.ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Parsing Arguments: %v\n", err)
		return
	}
	graph, err := parser.MapParser(cfg.FilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Parsing Map: %v\n", err)
		return
	}
}
