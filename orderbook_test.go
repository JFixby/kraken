package kraken

import (
	"github.com/jfixby/kraken/input"
	"github.com/jfixby/kraken/orderbook"
	testoutput "github.com/jfixby/kraken/output"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"testing"
	"time"
)

var setup *testing.T

func TestOrderbook(t *testing.T) {
	setup = t

	book := orderbook.NewBook()

	home := fileops.Abs("")
	testData := filepath.Join(home, "data", "test1")
	testOutput := filepath.Join(testData, "out", "output_file.csv")
	testInput := filepath.Join(testData, "in", "input_file.csv")

	test := &testoutput.TestOutput{File: testOutput}
	test.LoadAll()

	reader := input.NewFileReader(testInput)
	testListener := &TestListener{
		testData: test,
		book:     book}
	reader.Subscribe(testListener)
	reader.Run()

	for reader.IsRunnung() {
		time.Sleep(2 * time.Second)
	}

	pin.D("EXIT")
}

type TestListener struct {
	testData *testoutput.TestOutput
	scenario string
	book     *orderbook.Book
	counter  int
}

func (t *TestListener) DoProcess(ev *orderbook.Event) {

	result := t.book.DoUpdate(ev)
	expectedEvent := t.testData.GetEvent(t.scenario, t.counter)

	check(setup, result, expectedEvent, t.scenario, t.counter)

	pin.D("Event received", ev)
	t.counter++
}

func check(
	setup *testing.T,
	actual *orderbook.Result,
	expected *testoutput.TestEvent,
	scenario string,
	counter int) {

	pin.D("expected", expected)
	pin.D("  actual", actual)
}

func (t *TestListener) Reset(scenario string) {
	pin.D("Next scenario", scenario)
	t.scenario = scenario
	t.counter = 0
}
