package main

type Customer struct {
	Name string
}

type Coffee struct {
	Name     string
	Quantity int8
}

type Order interface {
	NewComer(name string) Customer
	OrderCoffee() Coffee
}
