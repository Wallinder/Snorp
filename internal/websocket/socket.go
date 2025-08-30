package socket

import (
	"context"
	"log"

	"github.com/coder/websocket"
)

func Connect(ctx context.Context, url string) (*websocket.Conn, error) {
	log.Printf("Connecting to socket: %s\n", url)
	ws, _, err := websocket.Dial(ctx, url+"/?v=10&encoding=json", nil)
	if err != nil {
		log.Printf("Error opening connection: %v\n", err)
		return ws, err
	}
	return ws, err
}

func Listen(ctx context.Context, conn *websocket.Conn, messageChannel chan []byte) {
	for {
		_, message, err := conn.Read(ctx)
		if err != nil {
			errorCode := int(websocket.CloseStatus(err))

			if SocketErrors[int(errorCode)] {
				log.Printf("Error %d: Trying to reconnect..\n", errorCode)
				messageChannel <- []byte("CTX_CLOSED")
				return
			}
			log.Fatalf("Unrecoverable error %d\n", errorCode)
		}
		select {
		case <-ctx.Done():
			return
		case messageChannel <- message:
		}
	}
}
