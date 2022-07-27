package listener

import "testing"

func TestListener(t *testing.T) {
	source := NewEvenSource()
	source.AddListener(NewOrderHandler())
	source.AddListener(NewOrderHandler())
	source.AddListener(NewOrderHandler())

	event := NewEvent("事件触发")
	source.notify(event)
}
