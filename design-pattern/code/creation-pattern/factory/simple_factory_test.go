package factory

import "testing"

func TestNewNotifier(t *testing.T) {
	emailNotifier := NewNotifier(email)
	emailNotifier.notify("send hello world by email")
	smsNotifier := NewNotifier(sms)
	smsNotifier.notify("send hello world by sms")
}
