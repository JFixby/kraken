package input

import (
	"bufio"
	"github.com/jfixby/kraken/orderbook"
	"github.com/jfixby/pin"
	"os"
	"strconv"
	"strings"
)

type FileReader struct {
	inputFile string
	listener  DataListener
	runFlag   bool
}

func NewFileReader(inputFile string) *FileReader {
	return &FileReader{
		inputFile,
		nil,
		false,
	}
}

func (r *FileReader) Subscribe(l DataListener) {
	r.listener = l
}

func (r *FileReader) IsRunnung() bool {
	return r.runFlag
}

func (r *FileReader) Stop() {
	r.runFlag = false
}

func (r *FileReader) Run() {
	if r.runFlag {
		return
	}

	r.runFlag = true
	go r.runthread()
}

func (r *FileReader) runthread() {
	input := r.inputFile
	pin.D("reading", input)
	file, err := os.Open(input)
	defer file.Close()
	if err != nil {
		r.runFlag = false
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() && r.runFlag {
		txt := scanner.Text()
		//		pin.D("", txt)
		event := ParseEvent(txt)
		if r.listener != nil {
			if event != nil {
				r.listener.DoProcess(event)
			}
		}

	}

	r.runFlag = false
}

func ParseEvent(txt string) *orderbook.Event {

	if txt == "" {
		return nil
	}

	if txt[0:1] == "#" {
		return nil
	}

	arr := strings.Split(txt, ", ")

	result := &orderbook.Event{}

	//pin.D(txt, arr)

	if arr[0] == "F" {
		result.OrderType = orderbook.FLUSH
		return result
	}

	if arr[0] == "N" {
		result.OrderType = orderbook.NEW
		return result
	}

	if arr[0] == "C" {
		result.OrderType = orderbook.CANCEL

		UserIDCancel, err := strconv.Atoi(arr[1])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.UserIDCancel = orderbook.OrderID(UserIDCancel)

		OrderIDCancel, err := strconv.Atoi(arr[2])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.OrderIDCancel = orderbook.OrderID(OrderIDCancel)

		return result
	}

	return nil
}
