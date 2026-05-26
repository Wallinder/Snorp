package discord

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/coder/websocket"
)

type DiscordPayload struct {
	Op int             `json:"op"`
	S  int64           `json:"s"`
	T  string          `json:"t"`
	D  json.RawMessage `json:"d"`
}

type DispatchMessage struct {
	Type string
	Data []byte
}

var (
	ErrConnAlreadyOpen = errors.New("connection already open")
)

func eventHandler(ctx context.Context, discord *Discord) error {
	if discord.Websocket.Conn != nil {
		return ErrConnAlreadyOpen
	}

	url := discord.Metadata.Url

	if discord.Connection.Resume {
		url = discord.ReadyData.ResumeGatewayURL
	}

	url += "/?v=" + discord.ApiVersion + "&encoding=json"

	var err error
	discord.Websocket.Conn, _, err = websocket.Dial(ctx, url, nil)
	if err != nil {
		return err
	}
	defer func() {
		discord.Websocket.Conn.Close(1006, "Normal Closure")
		discord.SetConnection(nil)
	}()

	for {
		_, message, err := discord.Websocket.Conn.Read(ctx)

		if err != nil {
			errorCode := int(websocket.CloseStatus(err))

			if SocketErrors[errorCode] {
				discord.Connection.SetResume(true)
				return fmt.Errorf("socket failure %v: code %d", err, errorCode)
			}
			panic(err)
		}

		var discordPayload DiscordPayload
		err = json.Unmarshal(message, &discordPayload)
		if err != nil {
			return err
		}
		opCode := discordPayload.Op
		opCodeType := EventCodes[opCode]

		TotalMessages.WithLabelValues(opCodeType).Inc()

		switch opCode {

		case HELLO:
			var interval Interval
			err := json.Unmarshal(discordPayload.D, &interval)
			if err != nil {
				return err
			}

			go func(interval int) error {
				ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return ErrContextCancelled

					case <-ticker.C:
						SendHeartbeat(ctx, discord.Websocket.Conn, discord.Connection.Seq)
						if err != nil {
							return fmt.Errorf("failed to send heartbeat: %v", err)
						}
					}
				}
			}(interval.Heartbeat)

			if discord.Connection.Resume {
				SendResume(ctx,
					discord.Websocket.Conn,
					discord.Identity.Token,
					discord.ReadyData.SessionID,
					discord.Connection.Seq,
				)
				continue
			}
			discord.Identity.Send(ctx, discord.Websocket.Conn)

		case HEARTBEAT:
			SendHeartbeat(ctx, discord.Websocket.Conn, discord.Connection.Seq)

		case DISPATCH:
			dispatchMessage := DispatchMessage{
				Type: discordPayload.T,
				Data: discordPayload.D,
			}

			discord.Connection.SetSequence(discordPayload.S)

			TotalDispatchMessages.WithLabelValues(dispatchMessage.Type).Inc()

			switch dispatchMessage.Type {

			case "READY":
				var readyData ReadyData
				if err := json.Unmarshal(dispatchMessage.Data, &readyData); err != nil {
					return err
				}
				discord.SetReadyData(readyData)

			default:
				discord.DispatchChan <- dispatchMessage
			}

		case RECONNECT:
			discord.Connection.SetResume(true)

		case INVALID_SESSION:
			var invalid bool
			if err := json.Unmarshal(discordPayload.D, &invalid); err != nil {
				return err
			}
			discord.Connection.SetResume(invalid)
		}
	}
}
