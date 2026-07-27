package parser

import (
	"bufio"
	"fmt"
	"os"
	"pathfinder/internal/models"
	"slices"
	"strings"
)

func MapParser(path string, start, end string) (models.Graph, error) {
	file, err := os.Open(path)
	if err != nil {
		return models.Graph{}, fmt.Errorf("opening network map: %w", err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		beforeComment, _, _ := strings.Cut(line, "#")
		trimmedLine := strings.TrimSpace(beforeComment)

		if trimmedLine == "" {
			continue
		}

		lines = append(lines, trimmedLine)
	}

	if err := scanner.Err(); err != nil {
		return models.Graph{}, fmt.Errorf("reading network map file: %w", err)
	}

	stations, connections, err := Categorize(lines)
	if err != nil {
		return models.Graph{}, fmt.Errorf("categorizing network map: %w", err)
	}

	validStations, err := StationValidator(stations)
	if err != nil {
		return models.Graph{}, fmt.Errorf("validating stations: %w", err)
	}
	if _, ok := validStations[start]; !ok {
		return models.Graph{}, fmt.Errorf("start station %q does not exist", start)
	}
	if _, ok := validStations[end]; !ok {
		return models.Graph{}, fmt.Errorf("end station %q does not exist", end)
	}

	validConnections, err := ConnectionValidator(connections, validStations)
	if err != nil {
		return models.Graph{}, fmt.Errorf("validating connections: %w", err)
	}

	g := models.Graph{
		Stations:    validStations,
		Connections: validConnections,
	}

	return g, nil
}

func Categorize(lines []string) ([]string, []string, error) {
	var stations []string
	var connections []string

	stationsIdx := slices.Index(lines, "stations:")
	connectionsIdx := slices.Index(lines, "connections:")

	if stationsIdx == -1 {
		return nil, nil, fmt.Errorf(`missing "stations:" section`)
	}
	if connectionsIdx == -1 {
		return nil, nil, fmt.Errorf(`missing "connections:" section`)
	}
	stations = append(stations, lines[stationsIdx+1:connectionsIdx]...)
	connections = append(connections, lines[connectionsIdx+1:]...)

	return stations, connections, nil
}
