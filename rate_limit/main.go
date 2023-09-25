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

func endpointHandler(w http.ResponseWriter, r *http.Request) {
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

func rateLimitter(next http.Handler) http.Handler {
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
        next.ServeHTTP(w, r)
      } 
    })
    
}

func main() {
    fmt.Println("Server is running") 
  
    mux := http.DefaultServeMux

    mux.HandleFunc("/", endpointHandler)

    var handler http.Handler = mux
    handler = rateLimitter(handler)

    server := new(http.Server)
    server.Addr = ":3000"
    server.Handler = handler
    
    server.ListenAndServe()

    
}
