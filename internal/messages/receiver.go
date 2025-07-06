package message

import (
	"log"

	"github.com/gorilla/websocket"
)

func Receiver(conn *websocket.Conn, messageChannel chan map[string]any) {
	for {
		var raw map[string]any
		err := conn.ReadJSON(&raw)
		if err != nil {
			messageType, message, _ := conn.ReadMessage()
			log.Println(messageType, string(message))
			log.Fatal(err)
		}
		messageChannel <- raw
	}
}
