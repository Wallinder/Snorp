package event

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"snorp/internal/state"
	"strconv"
	"time"

	"github.com/coder/websocket"
)

type DiscordPayload struct {
	Op int             `json:"op"`
	S  int64           `json:"s"`
	T  string          `json:"t"`
	D  json.RawMessage `json:"d"`
}

func eventHandler(ctx context.Context, cancel context.CancelFunc, session *state.SessionState) {
	if session.Conn != nil {
		slog.Info("connection already open")
		return
	}

	url := session.Metadata.Url

	if session.Resume {
		url = session.ReadyData.ResumeGatewayURL
	}

	url += "/?v=" + session.Config.Bot.ApiVersion + "&encoding=json"

	slog.Info("connecting to websocket", "url", url)

	var err error
	session.Conn, _, err = websocket.Dial(ctx, url, nil)
	if err != nil {
		slog.Error("error opening connection", "error", err)
		return
	}

	defer func() {
		session.Conn.Close(1006, "Normal Closure")
		session.Mu.Lock()
		session.Conn = nil
		session.Mu.Unlock()
		cancel()
	}()

	for {
		_, message, err := session.Conn.Read(ctx)
		if err != nil {
			errorCode := int(websocket.CloseStatus(err))

			if SocketErrors[int(errorCode)] || errorCode == -1 {
				slog.Error("socket failure", "error", err, "code", errorCode)
				session.Mu.Lock()
				session.Resume = true
				session.Mu.Unlock()
				return
			}
			slog.Error("unrecoverable", "error", err)
			os.Exit(1)
		}

		var discordPayload DiscordPayload
		err = json.Unmarshal(message, &discordPayload)
		if err != nil {
			slog.Error("error unmarshaling json", "error", err)
			return
		}

		session.Metrics.TotalMessages.WithLabelValues(strconv.Itoa(discordPayload.Op)).Inc()

		switch discordPayload.Op {

		case HELLO:
			var heartbeat int
			err := json.Unmarshal(discordPayload.D, &heartbeat)
			if err != nil {
				slog.Error("error unmarshaling json", "error", err)
				return
			}

			go func(interval int) {
				slog.Info("starting heartbeat", "interval", interval/1000)

				ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return

					case <-ticker.C:
						sendHeartbeat(ctx, session.Conn, session.Seq)
					}
				}
			}(heartbeat)

			if session.Resume {
				resumeConnection(ctx, session.Conn, session)
			} else {
				identify(ctx, session.Conn, session.Config.Bot.Identity)
			}

		case HEARTBEAT:
			slog.Info("received heartbeat", "opcode", HEARTBEAT)
			slog.Info("sending heartbeat immediately", "opcode", HEARTBEAT)
			sendHeartbeat(ctx, session.Conn, session.Seq)

		case HEARTBEAT_ACK:
			slog.Info("received heartbeat", "opcode", HEARTBEAT_ACK)

		case DISPATCH:
			session.Mu.Lock()
			session.Seq = discordPayload.S
			session.Mu.Unlock()
			dispatchHandler(ctx, session, discordPayload.T, discordPayload.D)

		case RECONNECT:
			session.Resume = true
			slog.Info("trying to reconnect..", "opcode", RECONNECT)
			return

		case INVALID_SESSION:
			var invalid bool
			err := json.Unmarshal(discordPayload.D, &invalid)
			if err != nil {
				slog.Error("failed to unmarshal json", "error", err)
				return
			}
			slog.Warn("invalid session, trying to reconnect.", "opcode", INVALID_SESSION)

			session.Mu.Lock()
			if invalid {
				session.Resume = true
			} else {
				session.Resume = false
			}
			session.Mu.Unlock()
			return
		}
	}
}
