package input

import "github.com/jfixby/kraken/orderbook"

type DataReader interface {
	Subscribe(orderbook.OrderListener)
	Run()
	Stop()
}
