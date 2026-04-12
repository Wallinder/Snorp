package event

import (
	"context"
	"encoding/json"
	"log/slog"
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

func EventHandler(ctx context.Context, cancel context.CancelFunc, session *state.SessionState) {
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
		session.SetConnection(nil)
		cancel()
	}()

	for {
		_, message, err := session.Conn.Read(ctx)

		if err != nil {
			errorCode := int(websocket.CloseStatus(err))

			if SocketErrors[errorCode] {
				slog.Error("socket failure", "error", err, "code", errorCode)
				session.SetResume(true)
				return
			}
			state.LogAndExit("unrecoverable", err, 1)
		}

		var discordPayload DiscordPayload
		err = json.Unmarshal(message, &discordPayload)
		if err != nil {
			slog.Error("error unmarshaling json", "error", err)
			return
		}

		TotalMessages.WithLabelValues(strconv.Itoa(discordPayload.Op)).Inc()

		switch discordPayload.Op {

		case HELLO:
			var interval Interval
			err := json.Unmarshal(discordPayload.D, &interval)
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
			}(interval.Heartbeat)

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
			session.SetSequence(discordPayload.S)
			Dispatcher(ctx, session, discordPayload.T, discordPayload.D)

		case RECONNECT:
			session.SetResume(true)
			slog.Info("reconnecting", "opcode", RECONNECT)
			return

		case INVALID_SESSION:
			var invalid bool
			err := json.Unmarshal(discordPayload.D, &invalid)
			if err != nil {
				slog.Error("failed to unmarshal json", "error", err)
				return
			}
			slog.Warn("invalid session, trying to reconnect.", "opcode", INVALID_SESSION)

			if invalid {
				session.SetResume(true)
			} else {
				session.SetResume(false)
			}
			return
		}
	}
}
