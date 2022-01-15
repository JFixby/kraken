package orderbook

import "encoding/json"

type OrderType string
const NEW OrderType = "NEW"
const CANCEL OrderType = "CANCEL"
const FLUSH OrderType = "FLUSH"

type Side string
const BUY Side  = "BUY"
const SELL Side  = "SELL"

type Symbol string
type UserID int64
type OrderID int64
type Price int64
type Quantity int64

type Event struct {

	OrderType OrderType
	UserID UserID
	Symbol Symbol
	Price Price
	Quantity Quantity
	Side Side
	OrderID OrderID

}

func (e *Event) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}