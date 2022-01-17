package orderbook

import "github.com/MauriceGit/skiplist"

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
	orders map[Price]*skiplist.SkipList
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
	orders := market.orders[ev.Price]

	order := &Order{}
	order.OrderID = b.NewOrderID()
	order.Price = ev.Price
	order.Symbol = ev.Symbol
	order.Quantity = ev.Quantity
	order.Side = ev.Side

	if len(orders) == 0 {
		orders = append(orders, order)
		market.orders[ev.Price] = orders
		bev := &BookEvent{}
		bev.EventType = ACKNOWLEDGE
		bev.UserIDAcknowledge = ev.UserID
		bev.OrderIDAcknowledge = ev.OrderID
		//bev.Input = ev
		b.BookListener.OnBookEvent(bev)
		return
	}

}

func (b *Book) getMarket(symbol Symbol) *Market {
	if b.markets == nil {
		b.markets = map[Symbol]*Market{}
	}
	market := b.markets[symbol]
	if market == nil {
		market = &Market{Symbol: symbol}
		market.orders = map[Price]*skiplist.SkipList
		b.markets[symbol] = market
	}
	return market
}

func (b *Book) NewOrderID() OrderID {
	b.orderid = b.orderid + 1
	return OrderID(b.orderid)
}

func NewBook(l BookListener) *Book {
	return &Book{BookListener: l}
}
