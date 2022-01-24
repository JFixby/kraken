package kraken

import (
	"github.com/jfixby/kraken/input"
	"github.com/jfixby/kraken/orderbook"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"testing"
	"time"
)

/*
Read input data from file and print into console.
 */

func TestInput(t *testing.T) {
	home := fileops.Abs("")
	testData := filepath.Join(home, "data", "test1")
	testInput := filepath.Join(testData, "in", "input_file.csv")
	//testOutput := filepath.Join(expectedOutput, "out", "output_file.csv")

	reader := input.NewFileReader(testInput)
	testListener := &InputTestListener{}
	reader.Subscribe(testListener)
	reader.Run()

	for reader.IsRunnung() {
		time.Sleep(2 * time.Second)
	}

	pin.D("EXIT")
}

type InputTestListener struct {
}

func (t *InputTestListener) DoProcess(ev *orderbook.Event) {
	pin.D("Event received", ev)
}

func (t *InputTestListener) Reset(scenario string) {
	pin.D("Next scenario", scenario)

}

