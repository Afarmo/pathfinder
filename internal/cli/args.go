package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	FilePath string
	Start    string
	End      string
	Trains   int
}

// ParseArgs parses and validates command-line input and constructs a Config
func ParseArgs(args []string) (Config, error) {
	if len(args) != 5 {
		return Config{}, fmt.Errorf("insufficient arguments")
	}

	filepath := args[1]
	start := args[2]
	end := args[3]
	trainsArg := args[4]

	trains, err := strconv.Atoi(trainsArg)
	if err != nil {
		return Config{}, fmt.Errorf("number of trains must be an integer")
	}
	if trains < 1 {
		return Config{}, fmt.Errorf("number of trains must be greater than 0")
	}
	if start == end {
		return Config{}, fmt.Errorf("start and end station must be different")
	}
	if !strings.HasSuffix(filepath, ".map") && !strings.HasSuffix(filepath, ".txt") {
		return Config{}, fmt.Errorf("invalid file extension")
	}

	return Config{
		FilePath: filepath,
		Start:    start,
		End:      end,
		Trains:   trains,
	}, nil
}
