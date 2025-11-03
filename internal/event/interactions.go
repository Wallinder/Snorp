package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"snorp/internal/api"
	"snorp/internal/jobs"
	"snorp/internal/sql"
	"snorp/internal/state"
	"snorp/pkg/svv"
)

func InteractionHandler(ctx context.Context, session *state.SessionState, commandResponse api.CommandResponse) {
	if commandResponse.Version == 1 {
		switch commandResponse.Data.Type {

		case api.CHAT_INPUT:
			switch commandResponse.Data.Name {

			case jobs.SVV:
				for _, options := range commandResponse.Data.Options {
					switch options.Name {

					case jobs.SVV_REGNUMMER:

						callbackMessage := api.MessageCallback{
							Type: api.CHANNEL_MESSAGE_WITH_SOURCE,
							Data: api.MessageCallbackData{},
						}

						data, err := svv.GetVehicle(session.Config.SVV.ApiKey, options.Value)
						if err != nil {
							callbackMessage.Data.Content = fmt.Sprintf("Failed to lookup vehicle, %s", err)
							api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
						}
						for _, vehicle := range data.KjoretoydataListe {
							callbackMessage.Data.Content = fmt.Sprintf("%+v\n", vehicle)
							api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
						}
					}
				}
			}

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

		case api.PRIMARY_ENTRY_POINT:
			return
		}
	}
}
