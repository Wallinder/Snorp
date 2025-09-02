package event

import (
	"context"
	"encoding/json"
	"log"
	"menial/internal/action"
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

func MessageHandler(ctx context.Context, conn *websocket.Conn, session *state.SessionState, cancel context.CancelFunc) {
	defer cancel()

	if session.Resume {
		ResumeConnection(ctx, conn, session.Config.Bot.Token, session)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Return signal received from context")
			return

		case message, ok := <-session.Messages:
			if !ok {
				log.Println("Message channel closed")
				return
			}
			var discordPayload DiscordPayload

			err := json.Unmarshal(message, &discordPayload)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
				return
			}
			session.Seq = discordPayload.S

			switch discordPayload.Op {
			case HELLO:
				var heartbeat HeartbeatInterval
				err := json.Unmarshal(discordPayload.D, &heartbeat)
				if err != nil {
					log.Println("Error unmarshaling JSON:", err)
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
							SendHeartbeat(ctx, conn, discordPayload.S)
						}
					}
				}(heartbeat.Interval)
				if !session.Resume {
					SendIdentify(ctx, conn, session.Config.Bot.Identity, session.Config.Bot.Token)
				}

			case HEARTBEAT:
				log.Printf("Received opcode %d, sending hearbeat immediately..\n", discordPayload.Op)
				SendHeartbeat(ctx, conn, discordPayload.S)

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
}
