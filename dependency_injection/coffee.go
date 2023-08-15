package main

import "fmt"

type Order interface {
	CallWaiter(msg string)
	MakeCoffee() *Coffee
}

type Customer struct {
	Name string
}

type Coffee struct {
	Name     string
	Quantity int8
}

func NewCustomer(name string) Order {
	return &Customer{
		Name: name,
	}
}

func (customer *Customer) CallWaiter(msg string) {
	fmt.Println("Excuse me, may i order some coffee?")
}

func (customer *Customer) MakeCoffee() *Coffee {
	return &Coffee{
		Name:     "Caffuchino",
		Quantity: 1,
	}
}
