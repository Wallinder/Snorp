package interaction

import (
	"context"
	"fmt"
	"snorp/internal/api"
	"snorp/internal/state"
	"snorp/pkg/svv"
)

func SlashInteractions(ctx context.Context, session *state.SessionState, commandResponse api.CommandResponse) {
	switch commandResponse.Data.Name {

	case SVV:
		for _, options := range commandResponse.Data.Options {
			switch options.Name {

			case SVV_REGNUMMER:

				callbackMessage := api.MessageCallback{
					Type: api.CHANNEL_MESSAGE_WITH_SOURCE,
					Data: api.MessageCallbackData{},
				}

				data, err := svv.GetVehicle(session.Config.SVV.ApiKey, options.Value)
				if err != nil {
					callbackMessage.Data.Content = "Failed to lookup vehicle, need to be in format: AB1234567"
					api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
				}
				for _, vehicle := range data.KjoretoydataListe {
					callbackMessage.Data.Content = fmt.Sprintf(
						"ID: %s\nStatus: %s\nGodkjent: %s\nRegistrert: %s\n",
						vehicle.KjoretoyID.Kjennemerke,
						vehicle.Registrering.Registreringsstatus.KodeBeskrivelse,
						vehicle.Godkjenning.TekniskGodkjenning.Godkjenningsundertype.KodeVerdi,
						vehicle.Forstegangsregistrering.RegistrertForstegangNorgeDato,
					)
					api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
				}
			}
		}
	}
}
