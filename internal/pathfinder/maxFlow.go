package pathfinder

// Edmonds-Karp loop: keep finding augmenting paths until none left
func MaxFlow(fg AllfinderGraph, start, end string) {
	for {
		came := AugmentingPath(fg, start, end)
		if came == nil {
			break
		}
		PushFlow(came, start, end)
	}
}

// searches for one path from start to end
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

// commits a path found by AugmentingPath: increases Flow by 1
func PushFlow(from map[string]*Edge, start, end string) {
	node := end
	for node != start {
		edge := from[node]
		edge.Flow++
		edge.Reverse.Flow--
		node = edge.Reverse.To
	}
}
