package acceptance

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestSpaceMap(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "space.map")
	networkMap := []byte(`stations:
bond_square,20,6
apple_avenue,7,7
orange_junction,6,1
space_port,1,11

connections:
bond_square-apple_avenue
apple_avenue-orange_junction
orange_junction-space_port`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "bond_square", "space_port", "4")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 6 {
		t.Fatalf("Expected 6 movement turns, got %d", len(lines)-1)
	}

}
