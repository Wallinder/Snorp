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
		var msg = make([]byte, 2048)
		n, err := conn.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		messageChannel <- msg[:n]
	}
}

func Write(conn *websocket.Conn, message []byte) {
	if len(message) < 4096 {
		_, err := conn.Write(message)
		if err != nil {
			log.Fatalf("Failed to write message: %v", err)
		}
	} else {
		log.Println("Message exceeded byte limit..")
	}
}
