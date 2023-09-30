package main

import "fmt"

type Person struct {
  Name string
  Age string
}

func main() {
  newPerson := Person{
    Name: "John",
    Age: "17",
  }

  pointPerson := &newPerson
  pointPerson.Age = "28"

  fmt.Println("Name:", newPerson.Name)
  fmt.Println("Age:", newPerson.Age)
}
