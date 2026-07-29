package acceptance

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestNumbersMap(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "numbers.map")
	networkMap := []byte(`stations:
one,1,1
two,2,2
three,3,3
four,4,4
five,5,5
six,6,6

connections:
two-three
five-one
three-one
two-five
one-four
six-two
one-six
`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "two", "four", "4")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 6 {
		t.Fatalf("Expected 6 movement turns, got %d", len(lines)-1)
	}

}
