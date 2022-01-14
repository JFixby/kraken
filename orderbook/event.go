package orderbook

import "encoding/json"

type OrderType string
const NEW OrderType = "NEW"
const CANCEL OrderType = "CANCEL"
const FLUSH OrderType = "FLUSH"

type OrderSide string
const BUY OrderSide  = "BUY"
const SELL OrderSide  = "SELL"

type Symbol string
type UserID int64
type OrderID int64
type Price float64
type Quantity float64

type Event struct {

	Symbol Symbol

	OrderType OrderType

	UserIDBuy UserID
	UserIDSell UserID
	UserIDAcknowledge UserID
	UserIDCancel OrderID

	OrderIDBuy OrderID
	OrderIDSell OrderID
	OrderIDAcknowledge OrderID
	OrderIDCancel OrderID

	Price Price
	Quantity Quantity

}

func (e *Event) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}