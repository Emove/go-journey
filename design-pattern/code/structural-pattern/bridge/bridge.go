package bridge

import "fmt"

// 桥接模式分离抽象部分和实现部分。使得两部分独立扩展
// 桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换
// 策略模式使得抽象部分和实现部分分离，可以独立变化

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImpl interface {
	Send(text, to string)
}

type MessageSms struct {
}

type MessageEmail struct {
}

type CommonMessage struct {
	method MessageImpl
}

type UrgencyMessage struct {
	method MessageImpl
}

func ViaSms() MessageImpl {
	return &MessageSms{}
}

func ViaEmail() MessageImpl {
	return &MessageEmail{}
}

func NewCommonMessage(method MessageImpl) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func NewUrgencyMessage(method MessageImpl) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (*MessageSms) Send(text, to string) {
	fmt.Printf("send %s to %s via sms\n", text, to)
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via email\n", text, to)
}

func (com *CommonMessage) SendMessage(text, to string) {
	com.method.Send(text, to)
}

func (urgency *UrgencyMessage) SendMessage(text, to string) {
	urgency.method.Send(fmt.Sprintf("[Urgency]: %s", text), to)
}
