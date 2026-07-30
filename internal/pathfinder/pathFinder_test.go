package pathfinder

import (
	"os"
	"path/filepath"
	"pathfinder/internal/parser"
	"testing"
)

func TestPathExists(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "terminus.map")
	networkMap := []byte(`stations:
beginning,0,0
near,1,0
far,1,3
terminus,0,3

connections:
beginning-near
near-far`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	start := "beginning"
	end := "terminus"
	g, err := parser.MapParser(path, start, end)
	if err != nil {
		t.Fatal("unexpected parsing error:", err)
	}
	_, err = FindAllPaths(g, start, end)
	if err == nil {
		t.Fatal("expected error for no path existing between start and end, instead got:", err)
	}

}
