package main

import (
	"fmt"
	"os"
)

func main() {
  args := os.Args[1:]

  fmt.Println("Hello,", args[0])

}
