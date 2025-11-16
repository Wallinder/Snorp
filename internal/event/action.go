package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"snorp/internal/api"
	"snorp/internal/event/interaction"
	"snorp/internal/state"

	"github.com/coder/websocket"
)

func DispatchHandler(ctx context.Context, conn *websocket.Conn, session *state.SessionState, action string, dispatchMessage json.RawMessage) {
	go session.Metrics.TotalDispatchMessages.WithLabelValues(action).Inc()

	switch action {

	case "READY":
		log.Println("Handshake complete..")
		var readyData state.ReadyData
		err := json.Unmarshal(dispatchMessage, &readyData)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		session.ReadyData = readyData
		go interaction.RegisterCommands(ctx, session)

	case "GUILD_CREATE":
		var guild api.Guild
		err := json.Unmarshal(dispatchMessage, &guild)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

	case "GUILD_DELETE":
		var guild api.Guild
		err := json.Unmarshal(dispatchMessage, &guild)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

	case "GUILD_UPDATE":
		var guild api.Guild
		err := json.Unmarshal(dispatchMessage, &guild)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

	case "CHANNEL_CREATE":
		var channel api.GuildChannels
		err := json.Unmarshal(dispatchMessage, &channel)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

	case "CHANNEL_DELETE":
		var channel api.GuildChannels
		err := json.Unmarshal(dispatchMessage, &channel)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

	case "CHANNEL_UPDATE":
		var channel api.GuildChannels
		err := json.Unmarshal(dispatchMessage, &channel)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

	case "MESSAGE_CREATE":
		var message api.Message
		err := json.Unmarshal(dispatchMessage, &message)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

	case "INTERACTION_CREATE":
		var commandResponse api.CommandResponse
		err := json.Unmarshal(dispatchMessage, &commandResponse)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		if commandResponse.Version == 1 {
			switch commandResponse.Data.Type {

			case api.CHAT_INPUT:
				go interaction.SlashInteractions(ctx, session, commandResponse)

			case api.USER_COMMAND:
				go interaction.UserInteractions(ctx, session, commandResponse)

			case api.MESSAGE:
				go interaction.MessageInteractions(ctx, session, commandResponse)

			case api.PRIMARY_ENTRY_POINT:
				go interaction.EntryInteractions(ctx, session, commandResponse)
			}
		}

	case "RESUMED":
		log.Println("Connection successfully resumed..")

	case "PRESENCE_UPDATE":
		var presence api.Presence
		err := json.Unmarshal(dispatchMessage, &presence)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		fmt.Println(string(dispatchMessage))
		fmt.Println(presence)

	default:
		fmt.Println(string(dispatchMessage))
	}
}
