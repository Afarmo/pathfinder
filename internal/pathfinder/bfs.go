package pathfinder

import "fmt"

type Stations struct {
	Station    string
	Neighbours []*Stations
}

type Graph struct {
	Stations map[string][]string
}

// Testing
func NewGraph() Graph {
	return Graph{
		Stations: map[string][]string{
			"A": {"B", "C"},
			"B": {"A", "D"},
			"C": {"B", "D"},
			"D": {"B", "C"},
		},
	}
}

func FindShortestPath(g Graph, start, end string) ([]string, error) {

	visited := make(map[string]bool)
	tempRoute := []string{}
	shortestRoute := []string{}
	tempRoute = append(tempRoute, start)
	visited[start] = true
	// Looping neighbours of start
	for _, station := range g.Stations[start] {
		if station == end {
			shortestRoute = append(append(shortestRoute, tempRoute...), station)
			break
		}

		tempRoute = append(tempRoute, station)

		visited[station] = true
		// Looping neighbours of start neighbours
		for _, substation := range g.Stations[station] {

			// Going backwards
			if _, visitedStation := visited[substation]; visitedStation == true {
				continue
			}
			visited[substation] = true
			tempRoute = append(tempRoute, substation)
			if end == substation {

				shortestRoute = append(shortestRoute, tempRoute...)
			}
		}
		tempRoute = nil
		tempRoute = append(tempRoute, start)
	}
	if len(shortestRoute) == 0 {
		return nil, fmt.Errorf("end station could not be reached")
	}
	fmt.Printf("Path: %v\n\n", shortestRoute)
	return shortestRoute, nil
}
