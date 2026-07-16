package cli
import(
	"testing"
)

func TestCliParser(t *testing.T){
	randomTest:= []struct{
		name string
		line []string
		want bool
	}{
		{"valid", []string{"pwd","filePath.map", "start", "end", "5"}, false},
		{"negative number of train", []string{"pwd","filePath.map", "start", "end", "not int"}, true},
		{"invalid extension", []string{"pwd","filePath", "start", "end", "5"}, true},
		{"insufficient arguments", []string{"pwd","filePath.map", "end", "5"}, true},
		{"same start and end station", []string{"pwd","filePath.map", "start", "start", "5"}, true},
		{"negative number of train", []string{"pwd","filePath.map", "start", "end", "-5"}, true},
	}
	for _, testCases:= range randomTest{
		t.Run(testCases.name, func(t *testing.T){
			_, err:= ParseArgs(testCases.line)
			if (err != nil) != testCases.want{
				t.Errorf("wanted: %v, got: %v",testCases.want, err)
			}
		})
	}
}