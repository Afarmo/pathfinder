package pathfinder

import (
	"fmt"
	"pathfinder/internal/models"
	"slices"
)

// Converts a boolean connections map to a slice based adjacency map
func graphConverter(gph models.Graph) map[string][]string {

	adjacency := make(map[string][]string, len(gph.Connections))

	for key, innerMap := range gph.Connections {
		connections := make([]string, 0, len(innerMap))
		for innerKey := range innerMap {
			connections = append(connections, innerKey)
		}
		adjacency[key] = connections
	}

	return adjacency
}

func FindShortestPath(graph models.Graph, start, end string) ([]string, error) {

	connections := graphConverter(graph)
	visited := map[string]bool{}
	visited[start] = true

	queue := []string{}

	parents := map[string]string{}

	// Looping stations connected to start station
	for _, station := range connections[start] {
		if station == end {
			parents[end] = start
			break
		}
		queue = append(queue, station)
		visited[station] = true
		parents[station] = start

		// Looping stations connected to previous station
		for _, substation := range connections[queue[0]] {

			// Going backwards
			if _, visitedStation := visited[substation]; visitedStation == true {
				continue
			}

			parents[substation] = station

			if substation == end {
				break
			}

			queue = append(queue, substation)
			// visited[substation] = true
		}

		if _, found := parents[end]; found == true {
			break
		}
		if len(queue) == 0 {
			return nil, fmt.Errorf("no path found between %s and %s", start, end)
		}
		queue = queue[1:]

	}

	shortestRoute := reconstructPath(parents, end)
	return shortestRoute, nil
}

func reconstructPath(shortestRoute map[string]string, end string) []string {
	var res []string
	var next string = end

	res = append(res, next)

	for i := 0; i < len(shortestRoute); i++ {

		if parent, found := shortestRoute[next]; found == true {
			res = append(res, parent)
			next = parent
		}
	}
	slices.Reverse(res)

	return res
}
