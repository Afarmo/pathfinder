package main

import (
	"fmt"
	"os"

	"pathfinder/internal/cli"
	"pathfinder/internal/parser"
	"pathfinder/internal/pathfinder"
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
	allpath, err := pathfinder.FindAllPaths(graph, cfg.Start, cfg.End)
	if err != nil{
		fmt.Println("error")
	}
	fmt.Println(allpath)
	// path, err := pathfinder.FindShortestPath(graph, cfg.Start, cfg.End)

}
