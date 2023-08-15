package main

type Order interface {
	CallWaiter(msg string)
	MakeCoffee()
}

type Customer struct {
	Name string
}

type Coffee struct {
	Name     string
	Quantity int8
}
