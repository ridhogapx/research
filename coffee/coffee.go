//go:build wireinject
// +build wireinject

package main

import (
	"fmt"

	"github.com/google/wire"
)

type Coffee struct {
	Name     string
	Quantity int8
	Owner    string
}

type Order struct {
	Coffee Coffee
	Waiter string
}

func NewCoffee() Coffee {
	return Coffee{
		Name:     "Caffuchino",
		Quantity: 1,
		Owner:    "John",
	}
}

func NewOrder(coffee Coffee) Order {
	return Order{
		Coffee: coffee,
		Waiter: "Doe",
	}
}

func (order Order) Ordering() {
	fmt.Printf("%v ordering %v to %v \n", order.Coffee.Owner, order.Coffee.Name, order.Waiter)
}

func InitializeCoffee() Order {
	wire.Build(NewOrder, NewCoffee)
	return Order{}
}
