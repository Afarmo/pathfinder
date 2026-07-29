package acceptance

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
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

	trainFormat := strings.Fields(lines[0])
	if trainFormat[0] != "T1-victoria" && trainFormat[0] != "T1-euston" {
		t.Fatalf("Expected T1-victoria or T1-euston, got %q", trainFormat)
	}
}
