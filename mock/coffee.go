package main

import "github.com/stretchr/testify/mock"

type Customer struct {
	Name string
}

type Coffee struct {
	Name     string
	Quantity int8
}

type Order interface {
	NewComer(name string) Customer
}

type MockCoffee struct {
	Mock mock.Mock
}

func (m *MockCoffee) NewComer(name string) Customer {
	args := m.Mock.Called(name)
	return args.Get(0).(Customer)
}
