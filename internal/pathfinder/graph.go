package pathfinder

import (
	"pathfinder/internal/models"
)

// one directional connection in the capacity-aware flow graph.
type Edge struct {
	To      string
	Cap     int
	Flow    int
	Reverse *Edge
}

// adjacency list for the flow graph
type AllfinderGraph map[string][]*Edge

// AllPath builds the complete capacity-aware flow graph from the parsed network map
func AllPath(graph models.Graph, start, end string) AllfinderGraph {
	fg := make(AllfinderGraph)
	addCapacityToStation(fg, graph, start, end)
	addCapacityToConnection(fg, graph, start, end)
	return fg
}

// creates one connection between two stations
func AddEdge(fg AllfinderGraph, from, to string, cap int) {
	forward := &Edge{To: to, Cap: cap}
	reverse := &Edge{To: from, Cap: 0}
	forward.Reverse = reverse
	reverse.Reverse = forward
	fg[from] = append(fg[from], forward)
	fg[to] = append(fg[to], reverse)
}

// gives every intermediate station a capacity-1 bottleneck edge from its "-in" to its "-out"
func addCapacityToStation(fg AllfinderGraph, graph models.Graph, start, end string) {
	for name := range graph.Stations {
		if name == start || name == end {
			continue
		}
		AddEdge(fg, name+"-in", name+"-out", 1)
	}
}

// wiring the "-out" side of one station to the "-in" side of the other
func addCapacityToConnection(fg AllfinderGraph, graph models.Graph, start, end string) {
	for from, neighbour := range graph.Connections {
		for to := range neighbour {
			fromStation := nodeName(from, "-out", start, end)
			toStation := nodeName(to, "-in", start, end)
			AddEdge(fg, fromStation, toStation, 1)
		}
	}
}

// decides which node name to use for a station when wiring up
func nodeName(station, suffix, start, end string) string {
	if station == start || station == end {
		return station
	}
	return station + suffix
}
