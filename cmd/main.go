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

var Reset = "\033[0m"
var Magenta = "\033[35m"
var Red = "\033[31m"

func main() {
	cfg, err := cli.ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError Parsing Arguments: %v%s\n",Red, err, Reset)
		return
	}

	graph, err := parser.MapParser(cfg.FilePath, cfg.Start, cfg.End)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError Parsing Map: %v%s\n",Red, err, Reset)
		return
	}

	paths, err := pathfinder.FindAllPaths(graph, cfg.Start, cfg.End)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError: %v%s\n",Red, err, Reset)
		return
	}

	turns := scheduler.Scheduler(paths, cfg.Trains)

	w := bufio.NewWriter(os.Stdout)
	for _, turn := range turns {
		w.WriteString(Magenta + turn + Reset)
		w.WriteByte('\n')
	}
	w.Flush()
}
