package socket

import (
	"log"

	"github.com/coder/websocket"
)

// true = reconnect
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

func ErrorHandler(conn *websocket.Conn, err error) {
	switch websocket.CloseStatus(err) {
	case UNKNOWN_ERROR: // Try reconnecting?
		log.Println("UNKNOWN_ERROR")

	case UNKNOWN_OPCODE: // Don't do that!
		log.Println("UNKNOWN_OPCODE")

	case DECODE_ERROR: // Don't do that!
		log.Println("DECODE_ERROR")

	case NOT_AUTHENTICATED:
		log.Println("NOT_AUTHENTICATED")

	case AUTHENTICATION_FAILED:
		log.Println("AUTHENTICATION_FAILED")
		conn.CloseNow()

	case ALREADY_AUTHENTICATED: // Don't do that!
		log.Println("ALREADY_AUTHENTICATED")

	case INVALID_SEQ: // Reconnect and start a new session
		log.Println("INVALID_SEQ")

	case RATE_LIMITED:
		log.Println("RATE_LIMITED")

	case SESSION_TIMEOUT: // Reconnect and start a new session
		log.Println("SESSION_TIMEOUT")

	case INVALID_SHARD:
		conn.CloseNow()
		log.Fatal("INVALID_SHARD")

	case SHARDING_REQUIRED:
		conn.CloseNow()
		log.Fatal("SHARDING_REQUIRED")

	case INVALID_API_VERION:
		conn.CloseNow()
		log.Fatal("INVALID_API_VERION")

	case INVALID_INTENTS:
		conn.CloseNow()
		log.Fatal("INVALID_INTENTS")

	case DISALLOWED_INTENTS:
		conn.CloseNow()
		log.Fatal("DISALLOWED_INTENTS")

	default:
		log.Fatal(err)
	}
}
