package socket

import (
	"context"
	"log"

	"github.com/coder/websocket"
)

func Connect(ctx context.Context, url string) *websocket.Conn {
	ws, _, err := websocket.Dial(ctx, url+"/?v=10&encoding=json", nil)
	if err != nil {
		log.Fatal(err)
	}
	return ws
}

func Listen(ctx context.Context, conn *websocket.Conn, messageChannel chan []byte) {
	for {
		_, message, err := conn.Read(ctx)
		if err != nil {
			ErrorHandler(conn, err)
		}
		messageChannel <- message
	}
}
