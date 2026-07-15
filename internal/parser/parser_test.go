package parser

import (
	"slices"
	"testing"
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
		t.Fatal("expected error for missing stations section, got nil")
	}
}

func TestCategorizeMissingConnections(t *testing.T) {
	lines := []string{"stations:", "waterloo,3,1"}
	_, _, err := Categorize(lines)
	if err == nil {
		t.Fatal("expected error for missing connections section, got nil")
	}
}
