package models

type Station struct {
	Name string
	X    int
	Y    int
}

type Graph struct {
	Stations    map[string]*Station
	Connections map[string]map[string]bool
}
