package input

import (
	"bufio"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"github.com/jfixby/pin/lang"
	"os"
	"path/filepath"
	"testing"
)

func TestInput(t *testing.T) {
	home := fileops.Abs("")
	testData := filepath.Join(home, "data", "test1")
	testInput := filepath.Join(testData, "in", "input_file.csv")
	//testOutput := filepath.Join(testData, "out", "output_file.csv")

	//list := fileops.ListFiles(home, func(file string) bool { return true }, fileops.ALL_CHILDREN)

	pin.D("reading", testInput)
	file, err := os.Open(testInput)
	lang.CheckErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		txt:=scanner.Text()
		pin.D("",txt)
	}

	lang.CheckErr(err)

}
