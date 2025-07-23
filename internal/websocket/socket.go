package socket

import (
	"context"
	"log"

	"github.com/coder/websocket"
)

func Connect(url string) *websocket.Conn {
	ws, _, err := websocket.Dial(context.TODO(), url+"/?v=10&encoding=json", nil)
	if err != nil {
		log.Fatal(err)
	}
	return ws
}

func Listen(conn *websocket.Conn, messageChannel chan []byte) {
	for {
		_, message, err := conn.Read(context.TODO())
		if err != nil {
			ErrorHandler(conn, err)
		}
		messageChannel <- message
	}
}
