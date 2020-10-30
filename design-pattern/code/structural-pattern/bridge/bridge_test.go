package bridge

import "testing"

func TestCommonSms(t *testing.T) {
	m := NewCommonMessage(ViaSms())
	m.SendMessage("hello world", "Marshal")
}

func TestCommonEmail(t *testing.T) {
	m := NewCommonMessage(ViaEmail())
	m.SendMessage("hello world", "Marshal")
}

func TestUrgencySms(t *testing.T) {
	m := NewCommonMessage(ViaSms())
	m.SendMessage("hello world", "Marshal")
}

func TestUrgencyEmail(t *testing.T) {
	m := NewCommonMessage(ViaEmail())
	m.SendMessage("hello world", "Marshal")
}
