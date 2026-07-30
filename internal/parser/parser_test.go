package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStartStationDoesNotExist(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`stations:
beginning,0,0
near,1,0
far,1,3
terminus,0,3

connections:
beginning-near
beginning-terminus
near-far
terminus-far`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "nonexistent", "terminus")
	if err == nil {
		t.Fatal("expected error for nonexistent start station, instead got:", err)
	}
}

func TestEndStationDoesNotExist(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`stations:
beginning,0,0
near,1,0
far,1,3
terminus,0,3

connections:
beginning-near
beginning-terminus
near-far
terminus-far`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "beginning", "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent end station, instead got:", err)
	}
}

func TestDuplicateConnection(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`stations:
beginning,0,0
near,1,0
far,1,3
terminus,0,3

connections:
beginning-near
near-beginning
beginning-terminus
near-far
terminus-far`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "beginning", "terminus")
	if err == nil {
		t.Fatal("expected error for duplicate connections, instead got:", err)
	}
}

func TestInvalidCoordiates(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`stations:
beginning,-1,-1
near,1,0
far,1,3
terminus,0,3

connections:
beginning-near
beginning-terminus
near-far
terminus-far`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "beginning", "terminus")
	if err == nil {
		t.Fatal("expected error for invalid coordinates, instead got:", err)
	}
}

func TestDuplicateStationName(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`stations:
beginning,0,0
beginning,0,1
near,1,0
far,1,3
terminus,0,3

connections:
beginning-near
beginning-terminus
near-far
terminus-far`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "beginning", "terminus")
	if err == nil {
		t.Fatal("expected error for duplicate station name, instead got:", err)
	}
}

func TestInvalidStationName(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`stations:
Beginning,0,0
near,1,0
far,1,3
terminus,0,3

connections:
beginning-near
beginning-terminus
near-far
terminus-far`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "beginning", "terminus")
	if err == nil {
		t.Fatal("expected error for duplicate station name, instead got:", err)
	}
}

func TestStationsSection(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`beginning,0,0
near,1,0
far,1,3
terminus,0,3

connections:
beginning-near
beginning-terminus
near-far
terminus-far`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "beginning", "terminus")
	if err == nil {
		t.Fatal("expected error for missing \"stations:\" section, instead got:", err)
	}
}

func TestConnectionsSection(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`stations:
beginning,0,0
near,1,0
far,1,3
terminus,0,3

beginning-near
beginning-terminus
near-far
terminus-far`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "beginning", "terminus")
	if err == nil {
		t.Fatal("expected error for missing \"connections:\" section, instead got:", err)
	}
}

func TestInvalidConnection(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`stations:
beginning,0,0
near,1,0
far,1,3
terminus,0,3

connections:
beginning-near
beginning-terminus
near-far
terminus-far
terminus-ghost`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "beginning", "terminus")
	if err == nil {
		t.Fatal("expected error for connection with a nonexistent station, instead got:", err)
	}
}

func Test10KStationLimit(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "10k.map")

	var builder strings.Builder
	builder.WriteString("stations:\n")
	for i := 0; i < 10001; i++ {
		fmt.Fprintf(&builder, "%d,%d,%d\n", i, i+1, i+2)
	}
	builder.WriteString("connections:\n")

	if err := os.WriteFile(path, []byte(builder.String()), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := MapParser(path, "0", "1")
	if err == nil {
		t.Fatal("expected error for exceeding 10,000 station limit, instead got:", err)
	}
}
