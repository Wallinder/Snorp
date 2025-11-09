package interaction

import (
	"context"
	"encoding/json"
	"log"
	"snorp/internal/api"
	"snorp/internal/state"

	"gorm.io/gorm/clause"
)

func MessageInteractions(ctx context.Context, session *state.SessionState, commandResponse api.CommandResponse) {
	tx := session.DB.WithContext(ctx)

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
				Content: "Message archived, my liege",
			}

			result := tx.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(&state.ArchivedMessages{
				ID:         message.ID,
				Type:       message.Type,
				AuthorID:   message.Author.ID,
				GlobalName: message.Author.GlobalName,
				Username:   message.Author.Username,
				Content:    message.Content,
				Timestamp:  message.Timestamp,
			})
			if result.Error != nil {
				callbackMessage.Data.Content = "Failed to archive message, sorry"
				api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
			} else {
				api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
			}
		}
	}
}
