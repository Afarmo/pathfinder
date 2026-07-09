package main

import (
	"fmt"
	"os"
	"pathfinder/internal/cli"
)

func main() {
	args, err := cli.ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Parsing Arguments: %v", err)
		return
	}

	cfg, err := cli.ValidateArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Validating Arguments: %v", err)
		return
	}
	fmt.Println("Validated Arguments:", cfg.FilePath, cfg.Start, cfg.End, cfg.Trains)
}
