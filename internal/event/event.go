package event

import (
	"context"
	"encoding/json"
	"log"
	"menial/config"
	"menial/internal/action"
	"menial/internal/state"

	"github.com/coder/websocket"
)

type DiscordPayload struct {
	Op int             `json:"op"`
	S  int64           `json:"s"`
	T  string          `json:"t"`
	D  json.RawMessage `json:"d"`
}

func MessageHandler(ctx context.Context, conn *websocket.Conn, messageChannel chan []byte, config config.StaticConfig, sessionState *state.SessionState) {
	var discordPayload DiscordPayload
	if sessionState.Resume {
		ResumeConnection(ctx, conn, config.Bot.Token, sessionState)
	}
	for message := range messageChannel {
		if string(message) == "CTX_CLOSED" {
			return
		}
		err := json.Unmarshal(message, &discordPayload)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
			return
		}
		sessionState.Seq = discordPayload.S

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
				for {
					SendHeartbeat(ctx, conn, interval, discordPayload.S)
				}
			}(heartbeat.Interval)

			SendIdentify(ctx, conn, config.Bot.Identity, config.Bot.Token)

		case HEARTBEAT:
			log.Printf("Received opcode %d, sending hearbeat immediately..\n", discordPayload.Op)
			SendHeartbeat(ctx, conn, 0, discordPayload.S)

		case HEARTBEAT_ACK:
			log.Println("Received heartbeat..")

		case DISPATCH:
			action.DispatchHandler(sessionState, discordPayload.T, discordPayload.D)

		case RECONNECT:
			sessionState.Resume = true
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
				sessionState.Resume = true
			} else {
				sessionState.Resume = false
			}
			log.Printf("Received %d, trying to reconnect..\n", INVALID_SESSION)
			return
		}
	}
}
