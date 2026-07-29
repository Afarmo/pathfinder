package acceptance

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestSmallLargeMap(t *testing.T) {
	tmp := os.TempDir()
	path := filepath.Join(tmp, "small_large.map")
	networkMap := []byte(`stations:
small,4,0
large,4,6
00,0,0
01,0,1
02,0,2
03,0,3
04,0,4
05,0,5
10,1,0
11,1,1
12,1,2
13,1,3
14,1,4
15,1,5
20,2,0
21,2,1
22,2,2
23,2,3
24,2,4
25,2,5
30,3,0
31,3,1
32,3,2
33,3,3
34,3,4
35,3,5
36,3,6

connections:
24-25
24-23
23-12
small-32
32-33
33-34
34-35
35-36
36-22
small-10
10-11
10-20
11-12
11-14
12-large
12-03
small-13
13-14
14-15
small-00
00-01
01-02
02-03
03-04
20-21
20-25
21-15
21-22
21-30
22-large
25-30
30-31
31-large
04-05
05-large
`)

	if err := os.WriteFile(path, networkMap, 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../cmd", path, "small", "large", "9")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(output), "\n")
	if len(lines)-1 != 8 {
		t.Fatalf("Expected 8 movement turns, got %d", len(lines)-1)
	}

}
