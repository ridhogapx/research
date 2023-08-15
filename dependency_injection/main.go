package main

type Message string

type Event struct{}

type Greeter struct {
	Message Message
}

func NewMessage() Message {
	return Message("Hi, buddy!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}
