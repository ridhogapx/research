package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
      Body: "Hi! You've reaced API!",
    }

    err := json.NewEncoder(w).Encode(&msg)

    if err != nil {
      return
    }
}

func main() {
    fmt.Println("Server is running")  

    http.HandleFunc("/ping", endpointHandler)
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
      fmt.Println("Failed to listen server in port 8080")
    }

    
}
