package parser

import (
	"fmt"
	"pathfinder/internal/models"
	"strconv"
	"strings"
	"unicode"
)

func ConnectionValidator(lines []string, stations map[string]models.Station) (map[string]map[string]bool, error) {
	connections := make(map[string]map[string]bool)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid connection line: %q", line)
		}

		for _, part := range parts {
			part = strings.TrimSpace(part)
			for _, ch := range part {
				if unicode.IsLower(ch) || unicode.IsNumber(ch) || ch == '_' {
					continue
				}
				return nil, fmt.Errorf("invalid connection name: %q", part)
			}
		}

		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])

		if _, ok := stations[left]; !ok {
			return nil, fmt.Errorf("nonexistent station: %q", left)
		}
		if _, ok := stations[right]; !ok {
			return nil, fmt.Errorf("nonexistent station: %q", right)
		}

		if connections[left] == nil {
			connections[left] = make(map[string]bool)
		}
		if connections[right] == nil {
			connections[right] = make(map[string]bool)
		}
		if connections[left][right] || connections[right][left] {
			return nil, fmt.Errorf("duplicate connection between %q and %q", left, right)
		}

		connections[left][right] = true
		connections[right][left] = true
	}

	return connections, nil
}

func StationValidator(lines []string) (map[string]models.Station, error) {
	stations := make(map[string]models.Station)
	coords := make(map[[2]int]string)

	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid station line: %q", line)
		}
		name := strings.TrimSpace(parts[0])
		xCoord := strings.TrimSpace(parts[1])
		yCoord := strings.TrimSpace(parts[2])

		for _, ch := range name {
			if unicode.IsLower(ch) || unicode.IsNumber(ch) || ch == '_' {
				continue
			}
			return nil, fmt.Errorf("invalid station name: %q", name)
		}

		if _, exists := stations[name]; exists {
			return nil, fmt.Errorf("duplicate station name: %q", name)
		}

		x, err := strconv.Atoi(xCoord)
		if err != nil || x < 0 {
			return nil, fmt.Errorf("invalid x coordinate for %q: %q", name, xCoord)
		}

		y, err := strconv.Atoi(yCoord)
		if err != nil || y < 0 {
			return nil, fmt.Errorf("invalid y coordinate for %q: %q", name, yCoord)
		}

		coord := [2]int{x, y}
		if other, exists := coords[coord]; exists {
			return nil, fmt.Errorf("duplicate coordinates for %q and %q", name, other)
		}
		coords[coord] = name

		stations[name] = models.Station{
			Name: name,
			X:    x,
			Y:    y,
		}
	}

	if len(stations) > 10000 {
		return nil, fmt.Errorf("station limit exceeded: %d > 10000", len(stations))
	}

	return stations, nil
}
