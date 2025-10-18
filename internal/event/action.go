package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"snorp/internal/api"
	"snorp/internal/sql"
	"snorp/internal/state"

	"github.com/coder/websocket"
)

func DispatchHandler(ctx context.Context, conn *websocket.Conn, session *state.SessionState, action string, dispatchMessage json.RawMessage) {
	switch action {

	case "READY":
		log.Println("Handshake completed..")
		var readyData state.ReadyData
		err := json.Unmarshal(dispatchMessage, &readyData)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		session.ReadyData = readyData

	case "GUILD_CREATE":
		var guild api.Guild
		err := json.Unmarshal(dispatchMessage, &guild)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		go func() {
			conn, err := session.Pool.Acquire(ctx)
			if err != nil {
				log.Printf("Error acquiring connection %v\n", err)
				return
			}

			sql.InsertGuild(ctx, conn, &guild)
			for _, channel := range guild.Channels {
				sql.InsertChannel(ctx, conn, &channel, guild.ID)
			}
			for _, member := range guild.Members {
				sql.InsertUser(ctx, conn, &member)
				sql.InsertUserGuildMapping(ctx, conn, &member, guild.ID)
			}
		}()

	case "MESSAGE_CREATE":
		var message api.Message
		err := json.Unmarshal(dispatchMessage, &message)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		fmt.Println(message)

	case "RESUMED":
		log.Println("Connection successfully resumed..")

	default:
		fmt.Println(string(dispatchMessage))
	}
}
