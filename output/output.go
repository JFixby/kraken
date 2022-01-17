package testoutput

import (
	"bufio"
	"github.com/jfixby/pin"
	"os"
	"strconv"
	"strings"
)

import "encoding/json"

type EventType string

const Acknowledge EventType = "Acknowledge"
const Reject EventType = "Reject"
const Best EventType = "Best"
const Trade EventType = "Trade"

type Side string

const BUY Side = "BUY"
const SELL Side = "SELL"

type UserIDBuy int64
type UserIDSell int64
type UserIDReject int64
type UserIDAcknowledge int64

type OrderIDBuy int64
type OrderIDSell int64
type OrderIDAcknowledge int64
type OrderIDReject int64

type Price int64
type Quantity int64

func (e *TestEvent) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type TestEvent struct {
	EventType EventType

	UserIDAcknowledge UserIDAcknowledge
	UserIDSell        UserIDSell
	UserIDBuy         UserIDBuy
	UserIDReject      UserIDReject

	Price    Price
	Quantity Quantity
	Side     Side

	OrderIDBuy         OrderIDBuy
	OrderIDSell        OrderIDSell
	OrderIDAcknowledge OrderIDAcknowledge
	OrderIDReject      OrderIDReject
	ShallowAsk         bool
}

type TestOutput struct {
	File string
	data map[string][]*TestEvent
}

func (o *TestOutput) LoadAll() error {
	o.data = map[string][]*TestEvent{}

	pin.D("reading", o.File)
	file, err := os.Open(o.File)
	defer file.Close()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	tag := ""
	for scanner.Scan() {
		txt := scanner.Text()

		if strings.HasPrefix(txt, "#name: ") {
			tag = txt[len("#name: "):]
			{
				//pin.D("tag", tag)
				o.data[tag] = []*TestEvent{}
				continue
			}
		}

		event := TryToParse(txt)
		if event != nil {
			//pin.D("", event)
			o.data[tag] = append(o.data[tag], event)
		}

	}
	return nil
}

func (o *TestOutput) GetEvent(scenario string, counter int) *TestEvent {
	list := o.data[scenario]
	if list == nil {
		pin.E("scenario not found", scenario)
		pin.E("                  ", counter)
		pin.E("                  ", o.data)
		panic("")
	}
	return list[counter]
}

func TryToParse(txt string) *TestEvent {
	if txt == "" {
		return nil
	}

	if txt[0:1] == "#" {
		return nil
	}

	arr := strings.Split(txt, ", ")

	result := &TestEvent{}

	if arr[0] == "A" {
		result.EventType = Acknowledge

		userID, err := strconv.Atoi(arr[1])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.UserIDAcknowledge = UserIDAcknowledge(userID)

		orderID, err := strconv.Atoi(arr[2])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.OrderIDAcknowledge = OrderIDAcknowledge(orderID)
		return result
	}

	if arr[0] == "B" {
		result.EventType = Best

		if arr[1] == "S" {
			result.Side = SELL
		} else if arr[1] == "B" {
			result.Side = BUY
		} else {
			pin.E("Unknown order side", txt)
			return nil
		}

		if arr[2] == "-" {
			arr[2] = "-1"
		}
		price, err := strconv.Atoi(arr[2])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.Price = Price(price)

		if arr[3] == "-" {
			result.ShallowAsk = true
		} else {
			quantity, err := strconv.Atoi(arr[3])
			if err != nil {
				pin.E("invalid input", txt)
				return nil
			}
			result.Quantity = Quantity(quantity)
		}

		return result
	}

	if arr[0] == "R" {
		result.EventType = Reject

		userID, err := strconv.Atoi(arr[1])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.UserIDReject = UserIDReject(userID)

		orderID, err := strconv.Atoi(arr[2])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.OrderIDReject = OrderIDReject(orderID)
		return result
	}

	if arr[0] == "T" {
		result.EventType = Trade

		userIDBuy, err := strconv.Atoi(arr[1])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.UserIDBuy = UserIDBuy(userIDBuy)

		orderIDBuy, err := strconv.Atoi(arr[2])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.OrderIDBuy = OrderIDBuy(orderIDBuy)

		userIDSell, err := strconv.Atoi(arr[3])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.UserIDSell = UserIDSell(userIDSell)

		orderIDSell, err := strconv.Atoi(arr[4])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.OrderIDSell = OrderIDSell(orderIDSell)

		price, err := strconv.Atoi(arr[5])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.Price = Price(price)

		quantity, err := strconv.Atoi(arr[6])
		if err != nil {
			pin.E("invalid input", txt)
			return nil
		}
		result.Quantity = Quantity(quantity)
		return result
	}

	return nil
}
