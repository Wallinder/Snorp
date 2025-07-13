package event

import (
	"encoding/json"
	"log"
	"menial/internal/config"
	"menial/internal/state"
	"time"

	"golang.org/x/net/websocket"
)

type DiscordPayload struct {
	Op int    `json:"op"`
	S  int64  `json:"s"`
	T  string `json:"t"`
	D  []byte `json:"d"`
}

func MessageHandler(conn *websocket.Conn, messageChannel chan []byte, config config.StaticConfig, session state.Session) {
	var discordPayload *DiscordPayload
	for message := range messageChannel {
		err := json.Unmarshal(message, &discordPayload)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

		switch discordPayload.Op {

		case HELLO:
			var helloData HelloData
			err := json.Unmarshal(discordPayload.D, &helloData)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
			}

			go func(interval float64) {
				time.Sleep(time.Duration(interval) * time.Second)
				log.Printf("Starting heartbeat with an interval of %f seconds!\n", interval)
				for {
					SendHeartbeat(conn, interval, discordPayload.S)
				}
			}(helloData.HeartbeatInterval / 1000)

			SendIdentify(conn, config.Identity, config.Bot.Token)

		case HEARTBEAT:
			log.Printf("Received opcode %d, sending hearbeat immediately..\n", discordPayload.Op)
			SendHeartbeat(conn, 0, discordPayload.S)

		case HEARTBEAT_ACK:
			log.Println("Received heartbeat..")

		case DISPATCH:
			//if s, ok := message["s"].(float64); ok {
			//	conf.Heartbeat.D = int(s)
			//	conf.Resume.D.Seq = int(s)
			//}
			//if tcode, ok := message["t"].(string); ok {
			//	switch tcode {

			//	case action.READY:
			//		if data, ok := message["d"].(map[string]any); ok {
			//			if gateway, ok := data["resume_gateway_url"].(string); ok {
			//				log.Printf("Updated gateway value to: %s", gateway)
			//				conf.Gateway = gateway
			//			}
			//			if sessionid, ok := data["session_id"].(string); ok {
			//				log.Printf("Updated session value to: %s", sessionid)
			//				conf.Resume.D.SessionId = sessionid
			//			}
			//		}

			//	default:
			//		log.Println(message)
			//	}
			//}

		case RECONNECT:
			ResumeConnection(conn, config.Bot.Token, session.SessionId, discordPayload.S)

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
