package observer

import "fmt"

type Subject struct {
	observers []Observer
	context   string
}

type Observer interface {
	Update(subject *Subject)
}

type Reader struct {
	name string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (subject *Subject) Attach(o Observer) {
	subject.observers = append(subject.observers, o)
}

func (subject *Subject) notify() {
	for _, o := range subject.observers {
		o.Update(subject)
	}
}

func (subject *Subject) UpdateContext(context string) {
	subject.context = context
	subject.notify()
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (reader *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s \n", reader.name, s.context)
}
