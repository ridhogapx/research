package main

type Message string

type Greeter struct{}

type Event struct{}

func NewMessage() Message {
	return Message("Hi, buddy!")
}
