package event

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/internal/state"
	"snorp/pkg/discord"
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
	if session.WsConn != nil {
		slog.Info("connection already open")
		return
	}

	url := session.Metadata.Url

	if session.Connection.Resume {
		url = session.ReadyData.ResumeGatewayURL
	}

	url += "/?v=" + session.Config.Bot.ApiVersion + "&encoding=json"

	slog.Info("connecting to websocket", "url", url)

	var err error
	session.WsConn, _, err = websocket.Dial(ctx, url, nil)
	if err != nil {
		slog.Error("error opening connection", "error", err)
		return
	}

	defer func() {
		session.WsConn.Close(1006, "Normal Closure")
		session.SetConnection(nil)
	}()

	for {
		_, message, err := session.WsConn.Read(ctx)

		if err != nil {
			errorCode := int(websocket.CloseStatus(err))

			if SocketErrors[errorCode] {
				slog.Error("socket failure", "error", err, "code", errorCode)
				session.Connection.SetResume(true)
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
		opCode := discordPayload.Op
		opCodeType := EventCodes[opCode]

		TotalMessages.WithLabelValues(opCodeType).Inc()

		slog.Info("event", "opcode", opCodeType, "type", discordPayload.T)

		switch opCode {

		case discord.HELLO:
			var interval discord.Interval
			err := json.Unmarshal(discordPayload.D, &interval)
			if err != nil {
				slog.Error("error unmarshaling json", "error", err)
				return
			}

			go func(interval int) {
				ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return

					case <-ticker.C:
						discord.SendHeartbeat(ctx, session.WsConn, session.Connection.Seq)
						if err != nil {
							slog.Error("failed to send heartbeat", "error", err)
							return
						}
					}
				}
			}(interval.Heartbeat)

			if session.Connection.Resume {
				discord.SendResume(ctx,
					session.WsConn,
					session.Config.Bot.Identity.Token,
					session.ReadyData.SessionID,
					session.Connection.Seq,
				)
				continue
			}
			session.Config.Bot.Identity.Send(ctx, session.WsConn)

		case discord.HEARTBEAT:
			discord.SendHeartbeat(ctx, session.WsConn, session.Connection.Seq)

		case discord.DISPATCH:
			session.Connection.SetSequence(discordPayload.S)
			dispatcher(ctx, session, discordPayload.T, discordPayload.D)

		case discord.RECONNECT:
			session.Connection.SetResume(true)
			return

		case discord.INVALID_SESSION:
			var invalid bool
			if err := json.Unmarshal(discordPayload.D, &invalid); err != nil {
				slog.Error("failed to unmarshal json", "error", err)
				return
			}
			session.Connection.SetResume(invalid)
			return
		}
	}
}
