package parser

import (
	"fmt"
	"pathfinder/internal/models"
	"strconv"
	"strings"
	"unicode"
)

func ConnectionValidator(ConnectionLine []string, stations map[string]models.Station) (map[string]map[string]bool, error) {
	connections := make(map[string]map[string]bool)

	for _, line := range ConnectionLine {
		eachConnection := strings.Split(line, "-")
		if len(eachConnection) != 2 {
			return nil, fmt.Errorf("Invalid connection line %s", eachConnection)
		}
	
		connection1 := strings.TrimSpace(eachConnection[0])
		connection2 := strings.TrimSpace(eachConnection[1])
		
		_, exists := stations[connection1]
		_, ok := stations[connection2]

		if !exists {
			return nil, fmt.Errorf("station unknown: %s", connection1)
		}
		if !ok{
			return nil, fmt.Errorf("station unknown: %s", connection2)
		}
		if connections[connection1][connection2] {
			return nil, fmt.Errorf("connections duplicated: %s - %s", connection1, connection2)
		}

		if connections[connection1] == nil {
			connections[connection1] = make(map[string]bool)
		}
		if connections[connection2] == nil {
			connections[connection2] = make(map[string]bool)
		}

		connections[connection1][connection2] = true
		connections[connection2][connection1] = true
	}
	return connections, nil

}

func Stationvalidator(stationLines []string) (map[string]models.Station, error) {
	stations := make(map[string]models.Station)
	coords := make(map[[2]int]string)

	for _, line := range stationLines {
		eachparts := strings.Split(line, ",")
		if len(eachparts) != 3 {
			return nil, fmt.Errorf("invalid station line: %s", line)
		}
		name := strings.TrimSpace(eachparts[0])
		xCoord := strings.TrimSpace(eachparts[1])
		yCoord := strings.TrimSpace(eachparts[2])

		for _, v := range name {
			if unicode.IsLower(v) || unicode.IsNumber(v) || v == '_' {
				continue
			} else {
				return nil, fmt.Errorf("invalid staion name: %s", name)
			}
		}
		_, exists := stations[name]
		if exists {
			return nil, fmt.Errorf("duplicate station name: %s", name)
		}

		x, err := strconv.Atoi(xCoord)
		if err != nil || x <= 0 {
			return nil, fmt.Errorf("error changing x coordinate for %s", xCoord)
		}

		y, err := strconv.Atoi(yCoord)
		if err != nil || y < 0 {
			return nil, fmt.Errorf("error changing y coordinate for %s", yCoord)
		}

		coord := [2]int{x, y}
		another, exists := coords[coord]
		if exists {
			return nil, fmt.Errorf("coordinates duplicated %s & %s", name, another)
		}
		coords[coord] = name

		stations[name] = models.Station{Name: name, X: x, Y: y}
	}
	return stations, nil
}
