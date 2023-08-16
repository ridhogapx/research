package main

import (
	"fmt"

	"github.com/google/wire"
)

type Order interface {
	CallWaiter(msg string)
	MakeCoffee() *Coffee
	Start()
}

type OrderCoffee struct {
	Customer string
	Coffee   string
}

type Customer struct {
	Name string
}

type Coffee struct {
	Name     string
	Quantity int8
}

func NewCustomer(name string) Customer {
	return Customer{
		Name: name,
	}
}

func NewCoffee(customer Customer) Coffee {
	return Coffee{
		Name:     "Caffuchino",
		Quantity: 1,
	}
}

func NewOrderCoffee(customer Customer, coffee Coffee) Order {
	return &OrderCoffee{
		Customer: customer.Name,
		Coffee:   coffee.Name,
	}
}

func (order OrderCoffee) CallWaiter(msg string) {
	fmt.Println("Excuse me, may i order some coffee?")
}

func (order OrderCoffee) MakeCoffee() *Coffee {
	return &Coffee{
		Name:     "Caffuchino",
		Quantity: 1,
	}
}

func (order OrderCoffee) Start() {
	fmt.Printf("Customer %v ordering %v...", order.Coffee, order.Coffee)
}

func InitializeCoffee(string) Customer {
	wire.Build(NewCoffee, NewCustomer, NewOrderCoffee)
	return Customer{}
}
