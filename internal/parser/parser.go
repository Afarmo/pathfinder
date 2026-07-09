package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)
type Station struct {
	Name string
	X    int
	Y    int
}

func MapParser(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("map file dosn't exist")
	}
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		eachLine := scanner.Text()
		beforeComment, _, _ := strings.Cut(eachLine, "#")
		trimmedLine := strings.TrimSpace(beforeComment)
		if trimmedLine == "" {
			continue
		}
		lines = append(lines, trimmedLine)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading map file: %w", err)
	}
	return lines, nil
}

var nameRule = regexp.MustCompile(`^[a-z0-9_]+$`)

func Stationvalidator (stationLines []string) (map[string] *Station, error){
	stations := make(map[string]*Station)
	coords := make(map[[2]int]string)

	for _, line := range stationLines{
		eachparts := strings.Split(line, ",")
		if len(eachparts) != 3 {
			return nil, fmt.Errorf("invalid station line: %s", line)
		}
		name:= strings.TrimSpace(eachparts[0])
		xCoord:= strings.TrimSpace(eachparts[1])
		yCoord:= strings.TrimSpace(eachparts[2])

		if !nameRule.MatchString(name){
			return nil,fmt.Errorf("invalid staion name: %s", name)
		}
		_, exists := stations[name]
		if exists{
			return nil, fmt.Errorf("duplicate station name: %s", name)
		}

		x, err := strconv.Atoi(xCoord)
		if err != nil || x < 0{
			return nil, fmt.Errorf("error changing x coordinate for %s", xCoord)
		}

		y, err := strconv.Atoi(yCoord)
		if err != nil || y < 0{
			return nil, fmt.Errorf("error changing y coordinate for %s", yCoord)
		}

		coord := [2]int{x,y}
		another, exists:= coords[coord]
		if exists{
			return nil, fmt.Errorf("coordinates duplicated %s & %s", name, another)
		}
		coords[coord] = name

		stations[name] = &Station{Name: name, X:x, Y:y}
	}
	return stations , nil
}

func Categorize(lines []string) (stations []string, connections []string, err error) {

	stationsIdx := slices.Index(lines, "stations:")
	connectionsIdx := slices.Index(lines, "connections:")

	if stationsIdx == -1 {
		return nil, nil, fmt.Errorf("missing 'stations:' section")
	}
	if connectionsIdx == -1 {
		return nil, nil, fmt.Errorf("missing 'connections:' section")
	}
	stations = append(stations, lines[stationsIdx+1:connectionsIdx]...)
	connections = append(connections, lines[connectionsIdx+1:]...)
	return stations, connections, nil

}
