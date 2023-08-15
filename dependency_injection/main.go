package main

type Message string

type Event struct {
	Greeter Greeter
}

type Greeter struct {
	Message Message
}

func NewMessage() Message {
	return Message("Hi, buddy!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (g Greeter) Greet() Message {
	return g.Message
}
