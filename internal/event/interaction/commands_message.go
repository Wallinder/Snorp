package interaction

import (
	"context"
	"encoding/json"
	"log"
	"snorp/internal/api"
	"snorp/internal/sql"
	"snorp/internal/state"
)

func MessageInteractions(ctx context.Context, session *state.SessionState, commandResponse api.CommandResponse) {
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

		case ARCHIVE_MESSAGE:
			callbackMessage := api.MessageCallback{
				Type: api.CHANNEL_MESSAGE_WITH_SOURCE,
			}
			callbackMessage.Data = api.MessageCallbackData{
				Content: "Message saved, my liege",
			}

			err = sql.InsertMessage(ctx, session.Pool, message)

			if err != nil {
				callbackMessage.Data.Content = "Failed to save message, sorry"
				api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
			} else {
				api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
			}
		}
	}
}
