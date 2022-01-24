package kraken

import (
	"github.com/jfixby/kraken/output"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"testing"
)


func TestOutput(t *testing.T) {
	home := fileops.Abs("")
	testData := filepath.Join(home, "data", "test1")
	testOutput := filepath.Join(testData, "out", "output_file.csv")

	test:=&testoutput.TestOutput{File: testOutput}

	test.LoadAll()

	test.Print()


	pin.D("EXIT")
}


