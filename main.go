package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader {
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
}


func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Home Page")
}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {
  // Checking the origin
  upgrader.CheckOrigin = func(r *http.Request) bool { return true }
}
