package pathfinder

import "strings"

// reads off the actual clean, disjoint station-to-station paths
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

// DFS from start to end, following only edges that currently carry flow (Flow > 0)
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
		path = path[:len(path)-1] // dead end, backtrack
		return false
	}

	if dfs(start) {
		return path
	}
	return nil
}

// reduces Flow by 1 along every edge in the given path
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

// removes the internal "-in"/"-out" bookkeeping suffixes
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
