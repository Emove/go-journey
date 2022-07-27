package factory

import "fmt"

const (
	email = iota
	sms
)

type Notifier interface {
	notify(content string)
}

type EmailNotifier struct {
}

type SmsNotifier struct {
}

func (notifier *EmailNotifier) notify(content string) {
	fmt.Println("send email notify to customer with content:", content)
}

func (notifier *SmsNotifier) notify(content string) {
	fmt.Println("send sms notify to customer with content:", content)
}

func NewNotifier(genre int) Notifier {
	if genre == email {
		return &EmailNotifier{}
	} else if genre == sms {
		return &SmsNotifier{}
	}
	return nil
}
