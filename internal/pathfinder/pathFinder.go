package pathfinder

import (
	"fmt"
	"pathfinder/internal/models"
)

// ties the whole pipeline together
func FindAllPaths(graph models.Graph, start, end string) ([][]string, error) {
	fg := AllPath(graph, start, end)
	MaxFlow(fg, start, end)
	paths := getPaths(fg, start, end)

	if len(paths) == 0 {
		return nil, fmt.Errorf("no path found between %s and %s", start, end)
	}
	return paths, nil
}
