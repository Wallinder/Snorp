package event

import (
	"encoding/json"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func MessageHandler(conn *websocket.Conn, messageChannel chan []byte, eventStruct *EventStructs) {
	for message := range messageChannel {
		err := json.Unmarshal(message, &eventStruct.DiscordPayload)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

		switch eventStruct.DiscordPayload.Op {

		case HELLO:
			err := json.Unmarshal(message, &eventStruct.Hello)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
			}

			go func(interval float64) {
				time.Sleep(time.Duration(interval) * time.Second)
				log.Printf("Starting heartbeat with an interval of %f seconds!\n", interval)
				for {
					SendHeartbeat(conn, interval, eventStruct.DiscordPayload.S)
				}
			}(eventStruct.Hello.D.HeartbeatInterval / 1000)

			SendIdentify(conn)

		case HEARTBEAT:
			log.Printf("Received opcode %d, sending hearbeat immediately..\n", eventStruct.DiscordPayload.Op)
			//SendHeartbeat(conn, 0, conf.Heartbeat)

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

		case INVALID_SESSION:
			//if data, ok := message["d"].(bool); ok {
			//	if data {
			//		ResumeConnection(conn, conf)
			//	}
			//}

		case RECONNECT:
			//ResumeConnection(conn, conf)
		}
	}
}
