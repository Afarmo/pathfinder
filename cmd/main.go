package main

import (
	"fmt"
	"os"
	"time"

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
	_, err = pathfinder.FindShortestPath(graph, cfg.Start, cfg.End)
	path := [][]string{
		{"jungle", "grasslands", "suburbs", "clouds", "wetlands", "desert"},
		{"jungle", "farms", "downtown", "metropolis", "industrial", "desert"},
		{"jungle", "green_belt", "village", "mountain", "treetop", "desert"},
	}
	t := time.Now()
	turns := scheduler.Scheduler(path, cfg.Trains)
	elapsed := time.Since(t)
	for i, turn := range turns {
		fmt.Printf("Turn %d: %s\n", i+1, turn)
	}
	// for _, turn := range turns {
	// 	fmt.Println(turn)
	// }
	fmt.Print("\nTime elapsed (Scheduler): ", elapsed)
	fmt.Println("\nTime elapsed (Printer)  :", time.Since(t)-elapsed)

}
