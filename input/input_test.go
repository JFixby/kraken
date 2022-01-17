package input

import (
	"github.com/jfixby/kraken/orderbook"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"testing"
	"time"
)

func TestInput(t *testing.T) {
	home := fileops.Abs("")
	testData := filepath.Join(home, "data", "test1")
	testInput := filepath.Join(testData, "in", "input_file.csv")
	//testOutput := filepath.Join(testData, "out", "output_file.csv")

	reader := NewFileReader(testInput)
	testListener := &TestListener{}
	reader.Subscribe(testListener)
	reader.Run()

	for reader.IsRunnung() {
		time.Sleep(2 * time.Second)
	}

	pin.D("EXIT")
}

type TestListener struct {
}

func (t TestListener) DoProcess(ev *orderbook.OrderEvent) {
	pin.D("OrderEvent received", ev)
}

func (t TestListener) Reset(scenario string) {
	pin.D("Next scenario", scenario)
}

