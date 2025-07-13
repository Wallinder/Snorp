package event

import (
	"encoding/json"
	"log"
	"menial/config"
	"menial/internal/event/action"
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

func MessageHandler(conn *websocket.Conn, messageChannel chan []byte, config config.StaticConfig, session *state.Session) {
	var discordPayload DiscordPayload
	for message := range messageChannel {
		err := json.Unmarshal(message, &discordPayload)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		session.UpdateSeq(discordPayload.S)
		switch discordPayload.Op {

		case HELLO:
			var heartbeat HeartbeatInterval
			err := json.Unmarshal(discordPayload.D, &heartbeat)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
			}

			go func(interval float64) {
				time.Sleep(time.Duration(interval) * time.Second)
				log.Printf("Starting heartbeat with an interval of %f seconds!\n", interval)
				for {
					SendHeartbeat(conn, interval, session)
				}
			}(heartbeat.Interval / 1000)

			SendIdentify(conn, config.Identity, config.Bot.Token)

		case HEARTBEAT:
			log.Printf("Received opcode %d, sending hearbeat immediately..\n", discordPayload.Op)
			SendHeartbeat(conn, 0, session)

		case HEARTBEAT_ACK:
			log.Println("Received heartbeat..")

		case DISPATCH:
			switch discordPayload.T {

			case "READY":
				var ready action.ReadyData
				err := json.Unmarshal(discordPayload.D, &ready)
				if err != nil {
					log.Println("Error unmarshaling JSON:", err)
				}
				session.UpdateSessionId(ready.SessionID)
				session.UpdateGateway(ready.ResumeGatewayURL)

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
				ResumeConnection(conn, config.Bot.Token, session.SessionId, discordPayload.S)
			}

		case INVALID_SESSION:
			var invalid bool
			err := json.Unmarshal(discordPayload.D, &invalid)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
			}
			if invalid {
				ResumeConnection(conn, config.Bot.Token, session.SessionId, discordPayload.S)
			}
		}
	}
}
