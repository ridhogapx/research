package main

type Order interface {
	CallWaiter(msg string)
	MakeCoffee() Coffee
}

type Customer struct {
	Name string
}

type Coffee struct {
	Name     string
	Quantity int8
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Name: name,
	}
}
