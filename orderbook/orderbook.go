package orderbook

type Book struct {
}

type Result struct {
	Input *Event
}

func (b *Book) DoUpdate(ev *Event) *Result {
	r := &Result{}
	r.Input = ev

	return r
}

func NewBook() *Book {
	return &Book{}
}
