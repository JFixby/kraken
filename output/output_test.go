package testoutput

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"testing"
)


func TestInput(t *testing.T) {
	home := fileops.Abs("")
	testData := filepath.Join(home, "data", "test1")
	testOutput := filepath.Join(testData, "out", "output_file.csv")

	test:=&TestOutput{File:testOutput}

	test.LoadAll()


	pin.D("EXIT")
}


