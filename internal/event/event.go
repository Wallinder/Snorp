package event

import (
	"encoding/json"
	"log"
	"menial/config"
	"menial/internal/state"
	"time"

	"golang.org/x/net/websocket"
)

type DiscordPayload struct {
	Op int             `json:"op"`
	S  int64           `json:"s"`
	T  string          `json:"t"`
	D  json.RawMessage `json:"d"`
}

func MessageHandler(conn *websocket.Conn, messageChannel chan []byte, config config.StaticConfig, state *state.SessionState) {
	var discordPayload DiscordPayload
	for message := range messageChannel {
		err := json.Unmarshal(message, &discordPayload)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		state.Seq = discordPayload.S
		
		switch discordPayload.Op {
		case HELLO:
			var heartbeat HeartbeatInterval
			err := json.Unmarshal(discordPayload.D, &heartbeat)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
			}

			go func(interval int) {
				log.Printf("Starting heartbeat with an interval of %f seconds!\n", interval)
				for {
					SendHeartbeat(conn, interval, state)
				}
			}(heartbeat.Interval)

			SendIdentify(conn, config.Identity, config.Bot.Token)

		case HEARTBEAT:
			log.Printf("Received opcode %d, sending hearbeat immediately..\n", discordPayload.Op)
			SendHeartbeat(conn, 0, state)

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
				state.ReadyData = readyData

			default:
				log.Println(string(discordPayload.D))
			}

		case RECONNECT:
			var null *string
			err := json.Unmarshal(discordPayload.D, &null)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
			}
			if null == nil {
				ResumeConnection(conn, config.Bot.Token, state.ReadyData.SessionID, discordPayload.S)
			}

		case INVALID_SESSION:
			var invalid bool
			err := json.Unmarshal(discordPayload.D, &invalid)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
			}
			if invalid {
				ResumeConnection(conn, config.Bot.Token, state.ReadyData.SessionID, discordPayload.S)
			}
		}
	}
}
