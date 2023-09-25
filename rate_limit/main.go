package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

type Message struct {
  Status string `json:"status"`
  Body string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, request *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    msg := Message{
      Status: "Successfull",
      Body: "Hi! You've reached API!",
    }

    err := json.NewEncoder(w).Encode(&msg)

    if err != nil {
      return
    }
}

func rateLimitter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
    limiter := rate.NewLimiter(2, 4)
    
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      if !limiter.Allow() {
        message := Message{
          Status: "Failed",
          Body: "The API is hitting Rate Limit, please try again later",
        }

        w.WriteHeader(http.StatusTooManyRequests)   
        json.NewEncoder(w).Encode(&message)
        return
      } else {
        next(w, r)
      } 
    })
    
}


func main() {
    fmt.Println("Server is running")  

    http.HandleFunc("/ping", endpointHandler)
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
      fmt.Println("Failed to listen server in port 8080")
    }

    
}
