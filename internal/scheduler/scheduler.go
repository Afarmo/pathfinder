package scheduler

import (
	"fmt"
	"strings"
)

type Train struct {
	ID       int
	Position int
}

func Scheduler(path []string, trainCount int) []string {

	var Trains = make([]Train, trainCount)
	AssignID(Trains)
	var turns []string

	for !AllArrived(Trains, path) {

		var turn strings.Builder
		occupied := map[string]bool{}

		for i := range Trains {

			// Train has reached the destination
			if Trains[i].Position == len(path)-1 {
				continue
			}
			// If the next station is full then skip otherwise empty current station and move to next station
			if _, full := occupied[path[Trains[i].Position+1]]; full == true {
				continue
			} else if Trains[i].Position > 0 && Trains[i].Position < len(path)-1 {
				delete(occupied, path[Trains[i].Position])
			}

			Trains[i].Position++
			turn.WriteString(fmt.Sprintf("T%d-%s ", Trains[i].ID, path[Trains[i].Position]))
			if Trains[i].Position > 0 && Trains[i].Position < len(path)-1 {
				occupied[path[Trains[i].Position]] = true
			}
		}
		turns = append(turns, turn.String())
	}
	return turns
}

func AssignID(trains []Train) {
	for i := range trains {
		trains[i].ID = i + 1
	}
}

func AllArrived(t []Train, path []string) bool {
	for i := range t {
		if t[i].Position != len(path)-1 {
			return false
		}
	}

	return true
}
