package main

import (
	"github.com/stretchr/testify/mock"
)

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

type Result struct {
	Coffee MockCoffee
}

func (m *MockCoffee) NewComer(name string) string {
	args := m.Mock.Called(name)
	return args.String()
}

func Logging(res *Result) {

}
