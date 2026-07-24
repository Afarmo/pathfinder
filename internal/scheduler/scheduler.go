package scheduler

import (
	"strconv"
	"strings"
)

type Train struct {
	ID       int
	Position int
	PathID   int
}

func Scheduler(paths [][]string, trainCount int) []string {

	// Pre assign path lengths
	var pathLens = make([]int, len(paths))
	for i := range paths {
		pathLens[i] = len(paths[i]) - 1
	}

	var Trains = make([]*Train, trainCount)
	// Assign ID's to trains
	for i := range Trains {
		Trains[i] = &Train{ID: i + 1}
	}
	AssignTrains(Trains, paths)
	var turns []string
	var arrivedCount int
	// Main loop that runs as long as all trains haven't arrived
	for {

		// var turn strings.Builder
		occupied := map[string]bool{}
		var turn strings.Builder
		var moved bool
		for i := range Trains {
			train := Trains[i]
			// Train has already reached the destination
			if train.Position == pathLens[train.PathID] {
				arrivedCount++
				continue
			}
			if train.Position < pathLens[train.PathID] {
				moved = true
			}

			// Next station is full
			if _, full := occupied[paths[train.PathID][train.Position+1]]; full == true {
				continue
			}
			// Free current station
			if train.Position > 0 && train.Position < pathLens[train.PathID] {
				delete(occupied, paths[train.PathID][train.Position])
			}
			// Move train
			train.Position++

			turn.WriteByte('T')
			turn.WriteString(strconv.Itoa(train.ID))
			turn.WriteByte('-')
			turn.WriteString(paths[train.PathID][train.Position])
			turn.WriteByte(' ')

			// Occupy the new station
			if train.Position > 0 && train.Position < pathLens[train.PathID] {
				occupied[paths[train.PathID][train.Position]] = true
			}
		}

		if turn.Len() > 0 {
			turns = append(turns, strings.TrimSuffix(turn.String(), " "))

		}
		if !moved {
			break
		}
	}
	return turns
}

func AssignTrains(trains []*Train, paths [][]string) {
	counts := make([]int, len(paths))

	for i := range trains {
		bestPath := 0
		bestArrival := 1<<31 - 1 // max int
		for j, path := range paths {
			arrival := len(path) - 1 + counts[j]
			if arrival < bestArrival {
				bestArrival = arrival
				bestPath = j
			}
		}

		trains[i].PathID = bestPath
		counts[bestPath]++
	}
}
