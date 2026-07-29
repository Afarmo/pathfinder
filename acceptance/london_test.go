package acceptance

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestLondonMap(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "london.map")
	networkMap := []byte(`stations:
waterloo,3,1
victoria,6,7
euston,11,23
st_pancras,5,15

connections:
waterloo-victoria
waterloo-euston
st_pancras-euston
victoria-st_pancras`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "waterloo", "st_pancras", "2")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 2 {
		t.Fatalf("Expected 2 movement turns, got %d", len(lines)-1)
	}

}

func TestLondonMapOneTrain(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "london.map")
	networkMap := []byte(`stations:
waterloo,3,1
victoria,6,7
euston,11,23
st_pancras,5,15

connections:
waterloo-victoria
waterloo-euston
st_pancras-euston
victoria-st_pancras`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "waterloo", "st_pancras", "1")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 2 {
		t.Fatalf("Expected 2 movement turns, got %d", len(lines)-1)
	}

	firstMovementTurn := strings.Fields(lines[0])
	secondMovementTurn := strings.Fields(lines[1])
	if len(firstMovementTurn) != 1 || len(secondMovementTurn) != 1 {
		t.Fatalf("Expected 1 move per movement turn , got %d and %d", len(firstMovementTurn), len(firstMovementTurn))
	}

}

func TestLondonMapThreeTrains(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "london.map")
	networkMap := []byte(`stations:
waterloo,3,1
victoria,6,7
euston,11,23
st_pancras,5,15

connections:
waterloo-victoria
waterloo-euston
st_pancras-euston
victoria-st_pancras`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "waterloo", "st_pancras", "3")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 3 {
		t.Fatalf("Expected 3 movement turns, got %d", len(lines)-1)
	}

}

func TestLondonMapFourTrains(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "london.map")
	networkMap := []byte(`stations:
waterloo,3,1
victoria,6,7
euston,11,23
st_pancras,5,15

connections:
waterloo-victoria
waterloo-euston
st_pancras-euston
victoria-st_pancras`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "waterloo", "st_pancras", "4")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 3 {
		t.Fatalf("Expected 3 movement turns, got %d", len(lines)-1)
	}

}

func TestLondonMap100Trains(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "london.map")
	networkMap := []byte(`stations:
waterloo,3,1
victoria,6,7
euston,11,23
st_pancras,5,15

connections:
waterloo-victoria
waterloo-euston
st_pancras-euston
victoria-st_pancras`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "waterloo", "st_pancras", "100")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 51 {
		t.Fatalf("Expected 51 movement turns, got %d", len(lines)-1)
	}

}

func TestLondonsMapMoves(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "london.map")
	networkMap := []byte(`stations:
waterloo,3,1
victoria,6,7
euston,11,23
st_pancras,5,15

connections:
waterloo-victoria
waterloo-euston
st_pancras-euston
victoria-st_pancras`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "waterloo", "st_pancras", "2")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")

	firstMovementTurn := strings.Fields(lines[0])
	secondMovementTurn := strings.Fields(lines[1])
	if len(firstMovementTurn) != 2 || len(secondMovementTurn) != 2 {
		t.Fatalf("Expected 2 move's per movement turn , got %d and %d", len(firstMovementTurn), len(firstMovementTurn))
	}
}
