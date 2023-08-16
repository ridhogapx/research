package main

type Cafe interface {
	Order() Coffee
	PayBill(int32) bool
}

type Customer struct {
	Name string
}

type Coffee struct {
	Name      string
	Quanntity int8
	Owner     Customer
}

type Bill struct {
	Cost   int32
	Coffee Coffee
}
