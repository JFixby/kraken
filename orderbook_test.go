package kraken

import (
	"github.com/jfixby/kraken/input"
	"github.com/jfixby/kraken/orderbook"
	testoutput "github.com/jfixby/kraken/output"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var setup *testing.T

// Both component test and usage example
func TestOrderBook(t *testing.T) {
	setup = t

	// input data
	home := fileops.Abs("")
	testData := filepath.Join(home, "data", "test1")
	testOutput := filepath.Join(testData, "out", "output_file.csv")
	testInput := filepath.Join(testData, "in", "input_file.csv")

	// expected output
	expectedOutput := &testoutput.TestOutput{File: testOutput}
	expectedOutput.LoadAll()

	// TestEnvironment wraps and tests Book component
	testEnvironment := &TestEnvironment{
		expectedOutput: expectedOutput}
	var bookEventListener orderbook.BookListener = testEnvironment

	//create book and subscribe it to TestEnvironment
	book := orderbook.NewBook(bookEventListener)
	testEnvironment.book = book

	// expected input will be read as a file and converted into event stream
	// fed to test environment
	inputFileReader := input.NewFileReader(testInput)
	inputFileReader.Subscribe(testEnvironment)
	inputFileReader.Run()

	// wait for tests to finish
	for inputFileReader.IsRunnung() {
		time.Sleep(2 * time.Second)
	}

	pin.D("EXIT")
}

type TestEnvironment struct {
	expectedOutput *testoutput.TestOutput
	scenario       string
	book           *orderbook.Book
	counter        int
}

// Receives input events and feeds them to the Book
func (t *TestEnvironment) DoProcess(ev *orderbook.Event) {
	pin.D("Input ", ev)
	t.book.DoUpdate(ev)
}

// Listents for events spawned by the Book and checks them against expected
func (t *TestEnvironment) OnBookEvent(e *orderbook.BookEvent) {
	pin.D("Output", e)
	expectedEvent := t.expectedOutput.GetEvent(t.scenario, t.counter)

	check(setup, e, expectedEvent, t.scenario, t.counter)
	t.counter++
}

// compares expected output with actual
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

// Resets book on each scenario
func (t *TestEnvironment) Reset(scenario string) {
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
