package main

import (
	"bufio"
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

	graph, err := parser.MapParser(cfg.FilePath, cfg.Start, cfg.End)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Parsing Map: %v\n", err)
		return
	}

	paths, err := pathfinder.FindAllPaths(graph, cfg.Start, cfg.End)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	turns := scheduler.Scheduler(paths, cfg.Trains)

	w := bufio.NewWriter(os.Stdout)
	for _, turn := range turns {
		w.WriteString(turn)
		w.WriteByte('\n')
	}
	w.Flush()
}
