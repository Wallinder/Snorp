package interaction

import (
	"context"
	"fmt"
	"log"
	"snorp/internal/api"
	"snorp/internal/jobs"
	"snorp/internal/state"
	"snorp/pkg/svv"
)

func SlashInteractions(mainCtx context.Context, session *state.SessionState, commandResponse api.CommandResponse) {
	callbackMessage := api.MessageCallback{
		Type: api.CHANNEL_MESSAGE_WITH_SOURCE,
		Data: api.MessageCallbackData{},
	}

	switch commandResponse.Data.Name {

	case SVV:
		for _, options := range commandResponse.Data.Options {
			switch options.Name {

			case SVV_REGNUMMER:
				data, err := svv.GetVehicle(session.Config.SVV.ApiKey, options.Value)
				if err != nil {
					callbackMessage.Data.Content = "Failed to lookup vehicle, need to be in format: AB1234567"
					api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
					return
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

	case CREATE:
		for _, options := range commandResponse.Data.Options {
			switch options.Name {

			case CREATE_CHANNEL:

				switch options.Value {

				case CHANNEL_WELCOME:
					if session.Jobs.Welcome[commandResponse.GuildID] != "" {
						callbackMessage.Data.Content = "This guild already have a welcome channel"
						api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
						return
					}

					callbackMessage.Data.Content = "Created a welcome channel"
					api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)

					newChannel := &api.GuildChannels{
						Name:  "new-phone-who-dis",
						Type:  api.GUILD_TEXT,
						Topic: "snorp:welcome",
					}

					channelID, err := api.FindOrCreateChannel(session, newChannel, commandResponse.GuildID)
					if err != nil {
						log.Println(err)
						return
					}
					session.Jobs.Welcome[commandResponse.GuildID] = channelID

				case CHANNEL_STEAM_NEWS:
					if session.Jobs.SteamNews[commandResponse.GuildID] {
						callbackMessage.Data.Content = "This guild already have an active steam-news job"
						api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
						return
					}

					callbackMessage.Data.Content = "Created a steam-news channel"
					api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)

					go jobs.SteamNewsFeed(mainCtx, session, commandResponse.GuildID)

				case CHANNEL_STEAM_SALES:
					if session.Jobs.SteamSales[commandResponse.GuildID] {
						callbackMessage.Data.Content = "This guild already have an active steam-sales job"
						api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
						return
					}

					callbackMessage.Data.Content = "Created a steam-sales channel"
					api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)

					go jobs.SteamSalesFeed(mainCtx, session, commandResponse.GuildID)
				}
			}
		}
	}
}
