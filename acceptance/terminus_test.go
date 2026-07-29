package acceptance

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestTerminusMap(t *testing.T) {
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
	cmd := exec.Command("go", "run", "../cmd", path, "beginning", "terminus", "20")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 11 {
		t.Fatalf("Expected 11 movement turns, got %d", len(lines)-1)
	}

}
