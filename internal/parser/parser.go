package parser

import (
	"bufio"
	"fmt"
	"os"
	"pathfinder/internal/models"
	"slices"
	"strings"
)

func MapParser(path string) (models.Graph, error) {
	file, err := os.Open(path)
	if err != nil {
		return models.Graph{}, fmt.Errorf("map file dosn't exist")
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
		return models.Graph{}, fmt.Errorf("error reading map file: %w", err)
	}

	stations, connections, err := Categorize(lines)
	if err != nil {
		return models.Graph{}, fmt.Errorf("error categorizing: %w", err)
	}
	validatedconnections, err := ConnectionValidator(connections)
	if err != nil {
		return models.Graph{}, fmt.Errorf("error validating connection: %w", err)
	}
	validatedstations, err := Stationvalidator(stations)
	if err != nil {
		return models.Graph{}, fmt.Errorf("error validating station: %w", err)
	}
	g := models.Graph{
		Stations:    validatedstations,
		Connections: validatedconnections,
	}
	return g, nil
}

func Categorize(lines []string) ([]string, []string, error) {
	var stations, connections []string
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
