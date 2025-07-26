package socket

import (
	"context"
	"log"
	"menial/internal/state"
	"time"

	"github.com/coder/websocket"
)

const (
	UNKNOWN_ERROR         = 4000 //true
	UNKNOWN_OPCODE        = 4001 //true
	DECODE_ERROR          = 4002 //true
	NOT_AUTHENTICATED     = 4003 //true
	AUTHENTICATION_FAILED = 4004 //false
	ALREADY_AUTHENTICATED = 4005 //true
	INVALID_SEQ           = 4007 //true
	RATE_LIMITED          = 4008 //true
	SESSION_TIMEOUT       = 4009 //true
	INVALID_SHARD         = 4010 //false
	SHARDING_REQUIRED     = 4011 //false
	INVALID_API_VERION    = 4012 //false
	INVALID_INTENTS       = 4013 //false
	DISALLOWED_INTENTS    = 4014 //false
)

func Connect(ctx context.Context, url string) *websocket.Conn {
	ws, _, err := websocket.Dial(ctx, url+"/?v=10&encoding=json", nil)
	if err != nil {
		log.Fatal(err)
	}
	return ws
}

func Listen(ctx context.Context, conn *websocket.Conn, messageChannel chan []byte, state *state.SessionState) {
	for {
		_, message, err := conn.Read(ctx)
		if err != nil {
			switch websocket.CloseStatus(err) {
			case UNKNOWN_ERROR: // Try reconnecting?
				conn.CloseNow()
				log.Println("UNKNOWN_ERROR")
				conn = Connect(ctx, state.Metadata.Url)

			case UNKNOWN_OPCODE: // Don't do that!
				conn.CloseNow()
				log.Println("UNKNOWN_OPCODE")
				conn = Connect(ctx, state.Metadata.Url)

			case DECODE_ERROR: // Don't do that!
				conn.CloseNow()
				log.Println("DECODE_ERROR")
				conn = Connect(ctx, state.Metadata.Url)

			case NOT_AUTHENTICATED:
				conn.CloseNow()
				log.Println("NOT_AUTHENTICATED")
				conn = Connect(ctx, state.Metadata.Url)

			case AUTHENTICATION_FAILED:
				conn.CloseNow()
				log.Println("AUTHENTICATION_FAILED")
				conn = Connect(ctx, state.Metadata.Url)

			case ALREADY_AUTHENTICATED: // Don't do that!
				conn.CloseNow()
				log.Println("ALREADY_AUTHENTICATED")
				conn = Connect(ctx, state.Metadata.Url)

			case INVALID_SEQ: // Reconnect and start a new session
				conn.CloseNow()
				log.Println("INVALID_SEQ")
				conn = Connect(ctx, state.Metadata.Url)

			case RATE_LIMITED:
				conn.CloseNow()
				log.Println("RATE_LIMITED")
				conn = Connect(ctx, state.Metadata.Url)

			case SESSION_TIMEOUT: // Reconnect and start a new session
				conn.CloseNow()
				log.Println("SESSION_TIMEOUT")

			case INVALID_SHARD:
				log.Fatal("INVALID_SHARD")

			case SHARDING_REQUIRED:
				log.Fatal("SHARDING_REQUIRED")

			case INVALID_API_VERION:
				log.Fatal("INVALID_API_VERION")

			case INVALID_INTENTS:
				log.Fatal("INVALID_INTENTS")

			case DISALLOWED_INTENTS:
				log.Fatal("DISALLOWED_INTENTS")

			default:
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
		}
		messageChannel <- message
	}
}
