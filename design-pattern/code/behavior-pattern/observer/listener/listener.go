package listener

import "fmt"

type Event struct {
	context string
}

type EventSource struct {
	listeners []EventListener
}

type EventListener interface {
	Handle(event Event)
}

func NewEvenSource() *EventSource {
	return &EventSource{
		listeners: make([]EventListener, 0),
	}
}

func (es *EventSource) AddListener(listener EventListener) {
	es.listeners = append(es.listeners, listener)
}

func (es *EventSource) notify(event Event) {
	for _, listener := range es.listeners {
		listener.Handle(event)
	}
}

func NewEvent(context string) Event {
	return Event{
		context: context,
	}
}

type OrderHandler struct {
	order int
}

var order = 0

func NewOrderHandler() *OrderHandler {
	order++
	return &OrderHandler{
		order: order,
	}
}

func (handler *OrderHandler) Handle(event Event) {
	fmt.Printf("order %d handler receive an event %s\n", handler.order, event.context)
}
