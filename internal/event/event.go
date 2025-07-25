package event

import (
	"encoding/json"
	"log"
	"menial/config"
	"menial/internal/state"
	socket "menial/internal/websocket"

	"github.com/coder/websocket"
)

type DiscordPayload struct {
	Op int             `json:"op"`
	S  int64           `json:"s"`
	T  string          `json:"t"`
	D  json.RawMessage `json:"d"`
}

func MessageHandler(conn *websocket.Conn, messageChannel chan []byte, config config.StaticConfig, sessionState *state.SessionState) {
	var discordPayload DiscordPayload
	for message := range messageChannel {
		err := json.Unmarshal(message, &discordPayload)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		sessionState.Seq = discordPayload.S

		switch discordPayload.Op {
		case HELLO:
			var heartbeat HeartbeatInterval
			err := json.Unmarshal(discordPayload.D, &heartbeat)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
			}

			go func(interval int) {
				log.Printf("Starting heartbeat with an interval of %d seconds!\n", interval)
				for {
					SendHeartbeat(conn, interval, discordPayload.S)
				}
			}(heartbeat.Interval)

			SendIdentify(conn, config.Identity, config.Bot.Token)

		case HEARTBEAT:
			log.Printf("Received opcode %d, sending hearbeat immediately..\n", discordPayload.Op)
			SendHeartbeat(conn, 0, discordPayload.S)

		case HEARTBEAT_ACK:
			log.Println("Received heartbeat..")

		case DISPATCH:
			switch discordPayload.T {

			case "READY":
				var readyData state.ReadyData
				err := json.Unmarshal(discordPayload.D, &readyData)
				if err != nil {
					log.Println("Error unmarshaling JSON:", err)
				}
				sessionState.ReadyData = readyData

			default:
				log.Println(string(discordPayload.D))
			}

		case RECONNECT:
			ResumeConnection(conn, config.Bot.Token, sessionState.ReadyData.SessionID, discordPayload.S)

		case INVALID_SESSION:
			var invalid bool
			err := json.Unmarshal(discordPayload.D, &invalid)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
			}
			if invalid {
				conn = socket.Connect(sessionState.Metadata.Url)
				ResumeConnection(conn, config.Bot.Token, sessionState.ReadyData.SessionID, discordPayload.S)
			} else {
				conn.Close(1000, "Normal Closure")
			}
		}
	}
}
