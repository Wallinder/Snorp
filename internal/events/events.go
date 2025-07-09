package event

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func SendHeartbeat(conn *websocket.Conn, interval float64, s int64) {
	message, err := json.Marshal(fmt.Sprintf("{op: 1, d: %d}", s))
	if err != nil {
		log.Fatal("Failed to marshal:", err)
	}
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal("Heartbeat failed", err)
	}
	time.Sleep(time.Duration(interval) * time.Second)
}

func SendIdentify(conn *websocket.Conn) {
	message, err := json.Marshal(identify)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Identifying..")
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal("Identify failed: ", err)
	}
}

func ResumeConnection(conn *websocket.Conn) {
	message, err := json.Marshal(resume)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Resuming connection..")
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal("Resuming failed", err)
	}
}
