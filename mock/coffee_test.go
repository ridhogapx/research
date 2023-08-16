package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

var mockCoffee = MockCoffee{Mock: mock.Mock{}}

func TestCoffee(t *testing.T) {
	mockCoffee.Mock.On("NewComer")

}
