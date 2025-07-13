package wss

import (
	"log"

	"golang.org/x/net/websocket"
)

func Listen(conn *websocket.Conn, messageChannel chan []byte) {
	for {
		var msg = make([]byte, 2048)
		n, err := conn.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		messageChannel <- msg[:n]
	}
}
