package event

import (
	"context"
	"encoding/json"
	"log"
	"menial/internal/socket/action"
	"menial/internal/state"
	"time"

	"github.com/coder/websocket"
)

type DiscordPayload struct {
	Op int             `json:"op"`
	S  int64           `json:"s"`
	T  string          `json:"t"`
	D  json.RawMessage `json:"d"`
}

func EventHandler(ctx context.Context, session *state.SessionState) {
	for {
		_, message, err := session.Conn.Read(ctx)
		if err != nil {
			errorCode := int(websocket.CloseStatus(err))

			if SocketErrors[int(errorCode)] || errorCode == -1 {
				log.Printf("Errorcode %d: %v\n", errorCode, err)
				session.Resume = true
				return
			}
			log.Fatalf("Unrecoverable error %d: %v\n", errorCode, err)
		}

		var discordPayload DiscordPayload

		err = json.Unmarshal(message, &discordPayload)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %v\n", err)
			return
		}
		session.Seq = discordPayload.S

		switch discordPayload.Op {

		case HELLO:
			var heartbeat HeartbeatInterval
			err := json.Unmarshal(discordPayload.D, &heartbeat)
			if err != nil {
				log.Printf("Error unmarshaling JSON: %v\n", err)
				return
			}

			go func(interval int) {
				log.Printf("Starting heartbeat with an interval of %d seconds!\n", interval/1000)

				ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return

					case <-ticker.C:
						SendHeartbeat(ctx, session.Conn, session.Seq)
					}
				}
			}(heartbeat.Interval)

			if session.Resume {
				ResumeConnection(ctx, session.Conn, session)
			} else {
				SendIdentify(ctx, session.Conn, session)
			}

		case HEARTBEAT:
			log.Printf("Received opcode %d, sending hearbeat immediately..\n", discordPayload.Op)
			SendHeartbeat(ctx, session.Conn, session.Seq)

		case HEARTBEAT_ACK:
			log.Println("Received heartbeat..")

		case DISPATCH:
			action.DispatchHandler(session, discordPayload.T, discordPayload.D)

		case RECONNECT:
			session.Resume = true
			log.Printf("Received %d, trying to reconnect..\n", RECONNECT)
			return

		case INVALID_SESSION:
			var invalid bool
			err := json.Unmarshal(discordPayload.D, &invalid)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
				return
			}
			if invalid {
				session.Resume = true
			} else {
				session.Resume = false
			}
			log.Printf("Received %d, trying to reconnect..\n", INVALID_SESSION)
			return
		}
	}
}
