package wss

import (
	"log"

	"golang.org/x/net/websocket"
)

func Connect(url string) *websocket.Conn {
	ws, err := websocket.Dial(url+"/?v=10&encoding=json", "", url)
	if err != nil {
		log.Fatal(err)
	}
	return ws
}
