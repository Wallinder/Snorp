package event

import (
	"encoding/json"
	"log"
	"menial/internal/config"

	"github.com/gorilla/websocket"
)

func SendIdentify(conn *websocket.Conn, identify config.Identify) {
	message, err := json.Marshal(identify)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Identifying..")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal("Identify failed: ", err)
	}
}
