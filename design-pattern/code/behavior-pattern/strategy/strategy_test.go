package strategy

import "testing"

func TestPayByCash(t *testing.T) {
	payment := NewPayment("Emove", "123456", 66, &Cash{})
	payment.Pay()
}

func TestPayByBank(t *testing.T) {
	payment := NewPayment("Emove", "123456", 88, &Bank{})
	payment.Pay()
}
