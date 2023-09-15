package main

import "fmt"

func main() {
  s := []string{"foo", "bar", "john", "doe", "yoa"}
  l := len(s)

  fmt.Println("Loop is running...")

  for i := 0; i < l; i ++ {
    val := s[i]
    fmt.Printf("%v Value: %v \n", i, val)
  }

  fmt.Println("finished loop...")

}
