package input

import "github.com/jfixby/kraken/orderbook"

type DataListener interface {
	Reset(scenario string)
	DoProcess(*orderbook.Event)
}

type DataReader interface {
	Subscribe(DataListener)
	Run()
	Stop()
}
