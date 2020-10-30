package facade

import "testing"

func TestNewOrderFacade(t *testing.T) {
	facade := NewOrderFacade()
	facade.Order()
}
