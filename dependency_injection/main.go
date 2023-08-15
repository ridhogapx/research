package main

type Message string

type Event struct{}

func NewMessage() Message {
	return Message("Hi, buddy!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

type Greeter struct {
	Message Message
}
