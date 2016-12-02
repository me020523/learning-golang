package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/**
 * handler func for websocket: ws://$host:$port/echo
 */
func echo(response http.ResponseWriter, request *http.Request) {
	upgrade := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrade.Upgrade(response, request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	t, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = conn.WriteMessage(t, msg)
	if err != nil {
		log.Println(err.Error())
	}
}
func main() {
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":8080", nil)
}
