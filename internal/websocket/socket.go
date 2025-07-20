package socket

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

func Listen(conn *websocket.Conn, messageChannel chan []byte) {
	for {
		var msg = make([]byte, 4096)
		n, err := conn.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		messageChannel <- msg[:n]
	}
}
