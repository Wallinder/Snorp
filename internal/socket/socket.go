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
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := conn.Read(ctx)
			if err != nil {
				errorCode := int(websocket.CloseStatus(err))
				if errorCode == -1 {
					log.Println(err)
					continue
				}
				if SocketErrors[int(errorCode)] {
					log.Printf("Errorcode %d: %v\n", errorCode, err)
					//errorChannel <- true
					return
				}
				log.Fatalf("Unrecoverable error %d: %v\n", errorCode, err)
			}
			messageChannel <- message
		}
	}
}
