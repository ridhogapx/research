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

func NewCustomer(name string) Cafe {
	return Customer{
		Name: name,
	}
}

func (customer Customer) Order() Coffee {
	return Coffee{
		Name:      "Italiano",
		Quanntity: 1,
		Owner:     customer,
	}
}

func (Customer) PayBill(cost int32) bool {
	return true
}
