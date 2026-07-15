package parser

import (
	"slices"
	"testing"
	"pathfinder/internal/models"
)

func TestCategorizeValid(t *testing.T) {
	lines := []string{
		"stations:",
		"waterloo,3,1",
		"victoria,6,7",
		"connections:",
		"waterloo-victoria",
	}

	stationLine, connectionLine, err := Categorize(lines)
	if err != nil {
		t.Errorf("expected no error: got: %v", err)
	}
	wantedStationLine := []string{"waterloo,3,1", "victoria,6,7"}
	wantedConnectionLine := []string{"waterloo-victoria"}
	if !slices.Equal(wantedStationLine, stationLine) {
		t.Errorf("expected: %+v, got: %+v", wantedStationLine, stationLine)
	}
	if !slices.Equal(wantedConnectionLine, connectionLine) {
		t.Errorf("expected: %+v, got: %+v", wantedConnectionLine, connectionLine)
	}
}
func TestCategorizeMissingStation(t *testing.T) {
	lines := []string{"connections:", "waterloo-victoria"}
	_, _, err := Categorize(lines)
	if err == nil {
		t.Errorf("expected error for missing stations section, got nil")
	}
}
func TestCategorizeMissingConnections(t *testing.T) {
	lines := []string{"stations:", "waterloo,3,1"}
	_, _, err := Categorize(lines)
	if err == nil {
		t.Errorf("expected error for missing connections section, got nil")
	}
}

func TestStationValidations(t *testing.T){
	randomTests := []struct{
		name string
		lines [] string
		want bool
	}{
		{"valid station",[]string{"waterloo,3,1"}, false},
		{"duplicate name", []string{"waterloo,3,1", "waterloo,6,7"}, true},
		{"duplicate coords", []string{"waterloo,3,1", "victoria,3,1"}, true},
		{"invalid name (capital)", []string{"Waterloo,3,1"}, true},
		{"negative x", []string{"waterloo,-1,1"}, true},
		{"negative y", []string{"waterloo,1,-1"}, true},
		{"non-numeric coord", []string{"waterloo,x,1"}, true},
		{"wrong number of fields", []string{"waterloo,3"}, true},
	}
	

	for _, testCases:= range randomTests{
		t.Run(testCases.name, func(t *testing.T){
			_, err := Stationvalidator(testCases.lines)
			if (err != nil) != testCases.want{
				t.Errorf("wanted err: %v, got: %v", testCases.want, err)
			}	
		})
		
	}
}

func TestConnectionvalidations(t *testing.T){
	stations := map[string]models.Station{
		"waterloo": {Name: "waterloo", X:3, Y:1},
		"victoria": {Name: "victoria", X:6, Y:7},
	}

	randomTests := []struct{
		name string
		lines []string
		want bool
	}{
		{"valid connection", []string{"waterloo-victoria"}, false},
		{"unknown station", []string{"waterloo-unknown"}, true},
		{"unknown station2", []string{"unknown-waterloo"}, true},
		{"duplicate connection", []string{"waterloo-victoria", "waterloo-victoria"}, true},
		{"reversed duplicate", []string{"waterloo-victoria", "victoria-waterloo"}, true},
		{"malformed line", []string{"waterloo victoria"}, true},
		{"only one station", []string{"waterloo"}, true},
	}
	for _,testCases := range randomTests{
		t.Run(testCases.name, func(t *testing.T){
			_, err := ConnectionValidator(testCases.lines, stations)
			if (err != nil) != testCases.want{
				t.Errorf("wanted err: %v, got: %v", testCases.want, err)
			}
		})
	}
}

