package message

import (
	"log"
	"menial/internal/config"
	event "menial/internal/events"
	action "menial/internal/events/actions"
	"time"

	"github.com/gorilla/websocket"
)

func Handler(conn *websocket.Conn, messageChannel chan map[string]any, conf config.Setting) {
	for message := range messageChannel {
		if opFloat, ok := message["op"].(float64); ok {
			opCode := int(opFloat)
			switch opCode {

			case event.HELLO:
				if data, ok := message["d"].(map[string]any); ok {
					if interval, ok := data["heartbeat_interval"].(float64); ok {
						interval = interval / 1000
						go func(interval int) {
							time.Sleep(time.Duration(interval) * time.Second)
							log.Printf("Starting heartbeat with an interval of %d seconds..\n", interval)
							for {
								event.SendHeartbeat(conn, interval, conf.Heartbeat)
							}
						}(int(interval))
						event.SendIdentify(conn, conf.Identify)
					}
				}

			case event.HEARTBEAT:
				log.Printf("Received opcode %d, sending hearbeat immediately..\n", opCode)
				event.SendHeartbeat(conn, 0, conf.Heartbeat)

			case event.RESUME:
				log.Println("received resume event")

			case event.HEARTBEAT_ACK:
				log.Println("Received heartbeat ack")

			case event.DISPATCH:
				if s, ok := message["s"].(float64); ok {
					conf.Heartbeat.D = int(s)
					conf.Resume.D.Seq = int(s)
				}
				if tcode, ok := message["t"].(string); ok {
					switch tcode {

					case action.READY:
						if data, ok := message["d"].(map[string]any); ok {
							if gateway, ok := data["resume_gateway_url"].(string); ok {
								conf.Env.Gateway = gateway
							}
							if sessionid, ok := data["session_id"].(string); ok {
								conf.Resume.D.SessionId = sessionid
							}
						}

					default:
						log.Println(message)
					}
				}
			case event.INVALID_SESSION:
				if data, ok := message["d"].(bool); ok {
					if data {
						event.ResumeConnection(conn, conf.Resume)
					}
				}

			case event.RECONNECT:
				event.ResumeConnection(conn, conf.Resume)

			default:
				log.Printf("Unknown message type: %d\n", opCode)
			}
		}
	}
}
