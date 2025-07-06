package event

import (
	"encoding/json"
	"log"
	"menial/internal/config"

	"github.com/gorilla/websocket"
)

func ResumeConnection(conn *websocket.Conn, resume config.Resume) {
	message, err := json.Marshal(resume)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Resuming connection..")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal("Resuming failed", err)
	}
}
