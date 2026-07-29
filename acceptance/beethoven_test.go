package acceptance

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestBeethovenMap(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "beethoven.map")
	networkMap := []byte(`stations:
beethoven,1,6
verdi,7,1
albinoni,1,1
handel,3,14
mozart,14,9
part,10,0

connections:
beethoven-handel
handel-mozart
beethoven-verdi
verdi-part
verdi-albinoni
beethoven-albinoni
albinoni-mozart
mozart-part`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "beethoven", "part", "9")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 6 {
		t.Fatalf("Expected 6 movement turns, got %d", len(lines)-1)
	}

}
