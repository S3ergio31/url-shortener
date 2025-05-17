package domain

type Event interface{}

type ShortFound struct {
	ShortCode  string
	Repository ShortRepository
}

type EventHandler func(Event)

func Publish(event Event) {
	switch event.(type) {
	case ShortFound:
		handle(event, []EventHandler{VisitCounter})
	}
}

func handle(event Event, handlers []EventHandler) {
	for _, handler := range handlers {
		handler(event)
	}
}
