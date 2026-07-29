package acceptance

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestJungleMap(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "jungle.map")
	networkMap := []byte(`stations:
jungle,5,16
green_belt,6,1
village,5,7
mountain,9,16
treetop,0,4
grasslands,15,13
suburbs,4,9
clouds,0,0
wetlands,2,12
farms,11,10
downtown,4,4
metropolis,3,20
industrial,1,18
desert,9,0

connections:
jungle-grasslands
mountain-treetop
clouds-wetlands
downtown-metropolis
green_belt-village
suburbs-clouds
industrial-desert
jungle-farms
village-mountain
wetlands-desert
grasslands-suburbs
jungle-green_belt
farms-downtown
treetop-desert
metropolis-industrial
mountain-wetlands
farms-mountain
`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "jungle", "desert", "10")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 8 {
		t.Fatalf("Expected 8 movement turns, got %d", len(lines)-1)
	}

}
