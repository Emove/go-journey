package listener

import "fmt"

// 监听器模式
// 监听器模式虽然在设计中，或者说在角色划分上看起来喝观察者模式不同
// 但是在实际上，监听器模式和观察者模式在设计思想上是统一的
// 他们有着一样的设计思想，在观察者模式中，subject角色相当于监听器中的Event和EventSource
// 在subject中，担负着事件触发和观察者的通知，而监听器则更加细化些
// 这两者没有谁好谁坏，在理解的概念上，更喜欢哪种就用哪种就好

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
