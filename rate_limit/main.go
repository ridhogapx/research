package main

import "net/http"

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
}

func main() {

}
