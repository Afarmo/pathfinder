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

// ParseArgs function reads command-line input and returns a slice of arguments and an error
func ParseArgs(args []string) ([]string, error) {
	if len(args) < 5 {
		return nil, fmt.Errorf("More arguments needed")
	}
	filepath := args[1]
	start := args[2]
	end := args[3]
	trains := args[4]

	return append([]string{}, filepath, start, end, trains), nil
}

func ValidateArgs(args []string) (*Config, error) {
	trains, err := strconv.Atoi(args[3])
	if err != nil {
		return nil, fmt.Errorf("[Number of trains] needs to be a number")
	}
	if trains > 10000 {
		return nil, fmt.Errorf("Number of trains exceeds 10 000")
	}
	if trains < 1 {
		return nil, fmt.Errorf("Number of trains needs to be atleast 1")
	}
	if args[1] == args[2] {
		return nil, fmt.Errorf("Matching start station and end station")
	}
	if !strings.HasSuffix(args[0], ".map") {
		return nil, fmt.Errorf("Invalid file extension")
	}
	cfg := &Config{
		FilePath: args[0],
		Start:    args[1],
		End:      args[2],
		Trains:   trains,
	}
	return cfg, nil
}
