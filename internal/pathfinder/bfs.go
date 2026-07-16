package pathfinder

import (
	"fmt"
	"slices"

	"pathfinder/internal/models"
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

	queue := []string{start}

	parents := map[string]string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, station := range connections[current] {

			if !visited[station] {

				parents[station] = current
				visited[station] = true
				if station == end {
					break
				}
				queue = append(queue, station)
			}

		}

		if _, found := parents[end]; found {
			break
		}
	}

	if _, found := parents[end]; !found {
		return nil, fmt.Errorf("no path found between %s and %s", start, end)
	}

	shortestRoute := reconstructPath(parents, end)
	return shortestRoute, nil
}

func reconstructPath(parents map[string]string, end string) []string {
	var path []string
	var child string = end

	path = append(path, child)

	for {
		parent, found := parents[child]
		if !found {
			break
		}

		path = append(path, parent)
		child = parent
	}
	slices.Reverse(path)

	return path
}
