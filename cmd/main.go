package main

import (
	"fmt"
	"os"

	"pathfinder/internal/cli"
	"pathfinder/internal/parser"
	"pathfinder/internal/pathfinder"
	"pathfinder/internal/scheduler"
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
	path, err := pathfinder.FindShortestPath(graph, cfg.Start, cfg.End)

	turns := scheduler.Scheduler(path, cfg.Trains)
	for i, turn := range turns {
		fmt.Printf("Turn %d: %s\n", i+1, turn)
	}
}
