package pathfinder

import (
	"fmt"
	"pathfinder/internal/models"
	"strings"
)

type Edge struct {
	To      string
	Cap     int
	Flow    int
	Reverse *Edge
}

type AllfinderGraph map[string][]*Edge

func FindAllPaths(graph models.Graph, start, end string) ([][]string, error) {
	fg := AllPath(graph, start, end)
	MaxFlow(fg, start, end)
	paths := getPaths(fg, start, end)
	if len(paths) == 0 {
		return nil, fmt.Errorf("no path found between %s and %s", start, end)
	}
	return paths, nil
}

func AllPath(graph models.Graph, start, end string) AllfinderGraph {
	fg := make(AllfinderGraph)
	addCapacityToStation(fg, graph, start, end)
	addCapacityToConnection(fg, graph, start, end)
	return fg
}

func getPaths(fg AllfinderGraph, start, end string) [][]string {
	var paths [][]string
	for {
		path := findFlowPath(fg, start, end)
		if path == nil {
			break
		}
		reducePathFlow(fg, path)
		paths = append(paths, deleteSuffixes(path))
	}
	return paths
}

func addCapacityToStation(fg AllfinderGraph, graph models.Graph, start, end string) {
	for name := range graph.Stations {
		if name == start || name == end {
			continue
		}
		AddEdge(fg, name+"-in", name+"-out", 1)
	}
}

func addCapacityToConnection(fg AllfinderGraph, graph models.Graph, start, end string) {
	for from, neighbour := range graph.Connections {
		for to := range neighbour {
			fromStation := nodeName(from, "-out", start, end)
			toStation := nodeName(to, "-in", start, end)
			AddEdge(fg, fromStation, toStation, 1)
		}
	}
}

func AddEdge(fg AllfinderGraph, from, to string, cap int) {
	forward := &Edge{To: to, Cap: cap}
	reverse := &Edge{To: from, Cap: 0}
	forward.Reverse = reverse
	reverse.Reverse = forward
	fg[from] = append(fg[from], forward)
	fg[to] = append(fg[to], reverse)
}

func nodeName(station, suffix, start, end string) string {
	if station == start || station == end {
		return station
	}
	return station + suffix
}

func findFlowPath(fg AllfinderGraph, start, end string) []string {
	visited := map[string]bool{start: true}
	var path []string

	var dfs func(node string) bool
	dfs = func(node string) bool {
		path = append(path, node)
		if node == end {
			return true
		}
		for _, edge := range fg[node] {
			if edge.Flow > 0 && !visited[edge.To] {
				visited[edge.To] = true
				if dfs(edge.To) {
					return true
				}
			}
		}
		path = path[:len(path)-1]
		return false
	}

	if dfs(start) {
		return path
	}
	return nil
}

func AugmentingPath(fg AllfinderGraph, start, end string) map[string]*Edge {
	visited := map[string]bool{start: true}
	queue := []string{start}
	from := map[string]*Edge{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, edge := range fg[current] {
			if edge.Cap-edge.Flow <= 0 {
				continue
			}
			if visited[edge.To] {
				continue
			}
			visited[edge.To] = true
			from[edge.To] = edge

			if edge.To == end {
				return from
			}
			queue = append(queue, edge.To)
		}
	}
	_, found := from[end]
	if !found {
		return nil
	}
	return from
}

func MaxFlow(fg AllfinderGraph, start, end string) {
	for {
		came := AugmentingPath(fg, start, end)
		if came == nil {
			break
		}
		PushFlow(came, start, end)
	}
}

func PushFlow(from map[string]*Edge, start, end string) {
	node := end
	for node != start {
		edge := from[node]
		edge.Flow++
		edge.Reverse.Flow--
		node = edge.Reverse.To
	}
}

func reducePathFlow(fg AllfinderGraph, path []string) {
	for i := 0; i < len(path)-1; i++ {
		for _, edge := range fg[path[i]] {
			if edge.To == path[i+1] && edge.Flow > 0 {
				edge.Flow--
				break
			}
		}
	}
}

func deleteSuffixes(path []string) []string {
	var cleanPath []string
	for _, str := range path {
		str = strings.TrimSuffix(str, "-in")
		str = strings.TrimSuffix(str, "-out")
		if len(cleanPath) == 0 || cleanPath[len(cleanPath)-1] != str {
			cleanPath = append(cleanPath, str)
		}
	}
	return cleanPath
}
