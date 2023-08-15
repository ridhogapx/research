package main

import (
	"errors"
	"fmt"
	"time"
)

type Message string

type Event struct {
	Greeter Greeter
}

type Greeter struct {
	Message Message
	Grumpy  bool
}

func NewMessage() Message {
	return Message("Hi, buddy!")
}

func NewGreeter(m Message) Greeter {
	var grumpy bool

	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}

	return Greeter{Message: m, Grumpy: grumpy}
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("error event greeter is grumpy")
	}

	return Event{Greeter: g}, nil
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away")
	}

	return g.Message
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	e := InitializeEvent()
	e.Start()
}
