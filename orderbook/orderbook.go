package orderbook

import "github.com/jfixby/kraken/util"

type BookListener interface {
	OnBookEvent(*BookEvent)
}

type Book struct {
	BookListener BookListener
	markets      map[Symbol]*Market
	orderid      int64
}

type Market struct {
	Symbol Symbol
	//orders map[Price][]Order
	//skipList := util.NewIntMap()
	orders *util.SkipList // price :-> order list, log N search
}

type Order struct {
	OrderID  OrderID
	Quantity Quantity
	Price    Price
	Symbol   Symbol
	Side     Side
}

func (b *Book) DoUpdate(ev *Event) {

	if ev.OrderType == NEW {
		b.NewOrder(ev)
	}

}

func (b *Book) NewOrder(ev *Event) {
	market := b.getMarket(ev.Symbol)
	v, ok := market.orders.Get(int(ev.Price))

	var orders []Order
	if !ok {
		orders = []Order{}
	} else {
		orders = v.([]Order)
	}

	order := Order{}
	order.OrderID = b.NewOrderID()
	order.Price = ev.Price
	order.Symbol = ev.Symbol
	order.Quantity = ev.Quantity
	order.Side = ev.Side

	if len(orders) == 0 {
		orders = append(orders, order)
		//market.orders[ev.Price] = orders
		market.orders.Set(int(ev.Price), orders)
		bev := &BookEvent{}
		bev.EventType = ACKNOWLEDGE
		bev.UserIDAcknowledge = ev.UserID
		bev.OrderIDAcknowledge = ev.OrderID
		//bev.Input = ev
		b.BookListener.OnBookEvent(bev)
		return
	}

	if len(orders) == 0 {
		
	}

}

func (b *Book) getMarket(symbol Symbol) *Market {
	if b.markets == nil {
		b.markets = map[Symbol]*Market{}
	}
	market := b.markets[symbol]
	if market == nil {
		market = &Market{Symbol: symbol}
		market.orders = util.NewIntMap()
		b.markets[symbol] = market
	}
	return market
}

func (b *Book) NewOrderID() OrderID {
	b.orderid = b.orderid + 1
	return OrderID(b.orderid)
}

func (b *Book) Reset() *Book {
	b.orderid = 0
	b.markets = nil
	return b
}

func NewBook(l BookListener) *Book {
	return (&Book{BookListener: l}).Reset()
}
