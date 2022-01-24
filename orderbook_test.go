package kraken

import (
	"github.com/jfixby/kraken/input"
	"github.com/jfixby/kraken/orderbook"
	testoutput "github.com/jfixby/kraken/output"
	"github.com/jfixby/kraken/util"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var setup *testing.T

func TestOrderbook(t *testing.T) {
	setup = t

	skipList := util.NewIntMap()
	skipList.Set(1, "one")

	home := fileops.Abs("")
	testData := filepath.Join(home, "data", "test1")

	//testOutput := filepath.Join(testData, "out", "output_file.csv")
	//testInput := filepath.Join(testData, "in", "input_file.csv")

	testOutput := filepath.Join(testData, "out", "output_file.csv")
	testInput := filepath.Join(testData, "in", "input_file.csv")

	test := &testoutput.TestOutput{File: testOutput}
	test.LoadAll()

	reader := input.NewFileReader(testInput)
	testListener := &TestListener{
		testData: test}
	reader.Subscribe(testListener)
	reader.Run()

	var bookEventListener orderbook.BookListener = testListener
	book := orderbook.NewBook(bookEventListener)
	testListener.book = book

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
	pin.D("Input ", ev)
	t.book.DoUpdate(ev)

}

func (t *TestListener) OnBookEvent(e *orderbook.BookEvent) {
	pin.D("Output", e)
	expectedEvent := t.testData.GetEvent(t.scenario, t.counter)

	check(setup, e, expectedEvent, t.scenario, t.counter)
	t.counter++
}

func check(
	setup *testing.T,
	actual *orderbook.BookEvent,
	expected *orderbook.BookEvent,
	scenario string,
	counter int) {

	if !expected.Equal(actual) {

		pin.D(" counter", counter)
		pin.D("expected", expected)
		pin.D("  actual", actual)
		//setup.FailNow()
		panic("")
	}
}

func (t *TestListener) Reset(scenario string) {
	pin.D("Next scenario", scenario)
	t.scenario = scenario
	t.counter = 0
	t.book.Reset()

	if strings.Contains(scenario, "13") {
		t.book.TradingModeON = true
	}
	if strings.Contains(scenario, "14") {
		t.book.TradingModeON = true
	}
	if strings.Contains(scenario, "T") {
		t.book.TradingModeON = true
	}
}
