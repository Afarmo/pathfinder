package scheduler

import (
	"strconv"
	"strings"
)

type Train struct {
	ID       int
	PathID   int
	Position int
}

func Scheduler(paths [][]string, trainCount int) []string {

	// Pre-allocate path lengths
	var pathLens = make([]int, len(paths))
	for i := range paths {
		pathLens[i] = len(paths[i])
	}

	// Assign trains to paths
	var Trains = make([]*Train, trainCount)
	pathTrainCount := make([]int, len(paths)) // trains in specific path

	for trainIndex := range Trains {

		bestPath := 0
		bestArrival := 1<<31 - 1

		for p := range paths {
			arrival := pathLens[p] - 1 + pathTrainCount[p]
			if arrival < bestArrival {
				bestArrival = arrival
				bestPath = p
			}
		}

		Trains[trainIndex] = &Train{
			ID:       trainIndex + 1,
			PathID:   bestPath,
			Position: pathTrainCount[bestPath],
		}

		pathTrainCount[bestPath]++
	}

	// Compute maximum turn by checking the longest arrival
	var maxTurn int
	for p := range paths {
		if pathTrainCount[p] == 0 {
			continue
		}
		arrival := pathLens[p] - 1 + pathTrainCount[p] - 1
		if arrival > maxTurn {
			maxTurn = arrival
		}
	}

	// Compute all moves per turn
	type Move struct {
		TrainID int
		Station string
	}

	moves := make([][]Move, maxTurn+1)

	for _, train := range Trains {

		for position := 1; position < pathLens[train.PathID]; position++ {
			turn := train.Position + position
			moves[turn] = append(moves[turn], Move{
				TrainID: train.ID,
				Station: paths[train.PathID][position],
			})
		}
	}

	// Construct all turns with moves
	var turns = make([]string, 0, maxTurn+1)

	for i := 0; i <= maxTurn; i++ {

		var turn strings.Builder

		for _, move := range moves[i] {
			turn.WriteByte('T')
			turn.WriteString(strconv.Itoa(move.TrainID))
			turn.WriteByte('-')
			turn.WriteString(move.Station)
			turn.WriteByte(' ')
		}

		if turn.Len() > 0 {
			turns = append(turns, strings.TrimSuffix(turn.String(), " "))
		}
	}

	return turns
}
