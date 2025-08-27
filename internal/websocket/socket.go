package socket

import (
	"context"
	"log"
	"menial/internal/state"

	"github.com/coder/websocket"
)

func Connect(ctx context.Context, url string) (*websocket.Conn, error) {
	ws, _, err := websocket.Dial(ctx, url+"/?v=10&encoding=json", nil)
	if err != nil {
		return ws, err
	}
	return ws, err
}

func Listen(ctx context.Context, conn *websocket.Conn, messageChannel chan []byte, state *state.SessionState) {
	for {
		_, message, err := conn.Read(ctx)
		if err != nil {
			errorCode := int(websocket.CloseStatus(err))

			if SocketErrors[int(errorCode)] {
				conn.Close(1006, "Reconnecting..")
				log.Printf("Error %d: Trying to reconnect..\n", errorCode)
				messageChannel <- []byte("CTX_CLOSED")
				return
			}
			log.Fatalf("Unrecoverable error %d\n", errorCode)
			return
		}
		select {
		case <-ctx.Done():
			return
		case messageChannel <- message:
		}
	}
}
