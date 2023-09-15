package main

import (
	"fmt"
	"sync"
)

func main() {
  s := []string{"foo", "bar", "john", "doe", "yoa"}
  l := len(s)

  var wg sync.WaitGroup
  wg.Add(l)

  fmt.Println("Loop is running...")

  for i := 0; i < l; i ++ {
    go func(i int) {
      defer wg.Done()
      val := s[i]
      fmt.Printf("%v Value: %v \n", i, val)
    }(i)

  }
  // doLoop(i)

  wg.Wait()

  fmt.Println("finished loop...")

}

func doLoop(i int) {
  var wg sync.WaitGroup
  defer wg.Done()
  fmt.Println(i)
}
