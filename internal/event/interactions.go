package event

import (
	"context"
	"encoding/json"
	"log"
	"snorp/internal/api"
	"snorp/internal/jobs"
	"snorp/internal/sql"
	"snorp/internal/state"
)

func InteractionHandler(ctx context.Context, session *state.SessionState, commandResponse api.CommandResponse) {
	if commandResponse.Version == 1 {
		switch commandResponse.Data.Type {

		case api.CHAT_INPUT:
			return

		case api.USER_COMMAND:
			return

		case api.MESSAGE:
			var data map[string]json.RawMessage
			err := json.Unmarshal(commandResponse.Data.Resolved.Messages, &data)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
				return
			}
			for _, value := range data {
				var message api.Message
				err := json.Unmarshal(value, &message)
				if err != nil {
					log.Println("Error unmarshaling JSON:", err)
					return
				}

				switch commandResponse.Data.Name {

				case jobs.ARCHIVE_MESSAGE:
					callbackMessage := api.MessageCallback{
						Type: api.CHANNEL_MESSAGE_WITH_SOURCE,
					}
					callbackMessage.Data = api.MessageCallbackData{
						Content: "Message Saved",
					}
					err = sql.InsertMessage(ctx, session.Pool, message)
					if err != nil {
						callbackMessage.Data.Content = "Failed to save message"
						api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
					} else {
						api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
					}
				}
			}

		case api.PRIMARY_ENTRY_POINT:
			return
		}
	}
}
