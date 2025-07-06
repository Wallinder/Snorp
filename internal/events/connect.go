package event

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func Connect(url string) *websocket.Conn {
	var connector = &websocket.Dialer{
		// https://github.com/gorilla/websocket/blob/main/client.go#L53
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}

	wss, _, err := connector.Dial(url+"/?v=10&encoding=json", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return wss
}
