package input

import "github.com/jfixby/kraken/orderbook"

type DataListener interface {
	DoProcess(orderbook.Event)
}

type DataReader interface {
	Subscribe(DataListener)
	Run()
	Stop()
}
