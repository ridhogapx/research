package main

import "fmt"

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

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}
