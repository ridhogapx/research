package main

import "github.com/go-redis/redis/v8"

func main() {
  client := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
  })
}
