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

		if strings.HasPrefix(txt, "#name: ") {
			tag := txt[len("#name: "):]
			if r.listener != nil {
				r.listener.Reset(tag)
				continue
			}
		}

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

		UserID, err := strconv.Atoi(arr[1])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.UserID = orderbook.UserID(UserID)

		result.Symbol = orderbook.Symbol(arr[2])

		Price, err := strconv.Atoi(arr[3])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.Price = orderbook.Price(Price)

		Qty, err := strconv.Atoi(arr[4])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.Quantity = orderbook.Quantity(Qty)

		if arr[5] == "S" {
			result.Side = orderbook.SELL
		} else if arr[5] == "B" {
			result.Side = orderbook.BUY
		} else {
			pin.E("Unknown order side", txt)
			return nil
		}

		if arr[6][len(arr[6])-1:len(arr[6])] == " " {
			arr[6] = arr[6][0 : len(arr[6])-1]
		}
		OrderID, err := strconv.Atoi(arr[6])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.OrderID = orderbook.OrderID(OrderID)

		return result
	}

	if arr[0] == "C" {
		result.OrderType = orderbook.CANCEL

		UserID, err := strconv.Atoi(arr[1])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.UserID = orderbook.UserID(UserID)

		if arr[2][len(arr[2])-1:len(arr[2])] == " " {
			arr[2] = arr[2][0 : len(arr[2])-1]
		}
		OrderID, err := strconv.Atoi(arr[2])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.OrderID = orderbook.OrderID(OrderID)

		return result
	}

	return nil
}
