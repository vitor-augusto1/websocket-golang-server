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
  // Upgrading the http connection
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println(err)
  }
  log.Println("Client connected")
  err = conn.WriteMessage(1, []byte("Hi client!"))
  if err != nil {
    log.Println(err)
  }
  reader(conn)
}


func reader(conn *websocket.Conn) {
  for {
    messageType, p, err := conn.ReadMessage()
    if err != nil {
      log.Println(err)
      return
    }
    // Printing the message
    fmt.Println(string(p))
    if err := conn.WriteMessage(messageType, p); err != nil {
      log.Println(err)
      return
    }
  }
}


