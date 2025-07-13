package event

import (
	"encoding/json"
	"log"
	"menial/internal/config"
	"time"

	"golang.org/x/net/websocket"
)

type HelloData struct {
	HeartbeatInterval float64 `json:"heartbeat_interval"`
}

type Heartbeat struct {
	Op int   `json:"op"`
	D  int64 `json:"d"`
}

func SendHeartbeat(conn *websocket.Conn, interval float64, seq int64) {
	message, err := json.Marshal(Heartbeat{
		Op: 1,
		D:  seq,
	})
	if err != nil {
		log.Fatal("Failed to marshal:", err)
	}
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal("Heartbeat failed", err)
	}
	time.Sleep(time.Duration(interval) * time.Second)
}

type Identify struct {
	Op int          `json:"op"`
	D  IdentifyData `json:"d"`
}

type IdentifyData struct {
	Token      string             `json:"token"`
	Intents    int64              `json:"intents"`
	Properties IdentifyProperties `json:"properties"`
}
type IdentifyProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

func SendIdentify(conn *websocket.Conn, conf config.Identity, token string) {
	message, err := json.Marshal(Identify{
		Op: 2,
		D: IdentifyData{
			Token:   token,
			Intents: conf.Intents,
			Properties: IdentifyProperties{
				Os:      conf.Properties.Os,
				Browser: conf.Properties.Browser,
				Device:  conf.Properties.Device,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Identifying..")
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal("Identify failed: ", err)
	}
}

type Resume struct {
	Op int
	D  ResumeData
}

type ResumeData struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

func ResumeConnection(conn *websocket.Conn, token string, sessionId string, seq int64) {
	message, err := json.Marshal(Resume{
		Op: 6,
		D: ResumeData{
			Token:     token,
			SessionId: sessionId,
			Seq:       seq,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Resuming connection..")
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal("Resuming failed", err)
	}
}
