package interaction

import (
	"context"
	"fmt"
	"slices"
	"snorp/internal/api"
	"snorp/internal/jobs"
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

	case STEAM:
		for _, options := range commandResponse.Data.Options {
			switch options.Name {

			case STEAM_NEWS:
				callbackMessage := api.MessageCallback{
					Type: api.CHANNEL_MESSAGE_WITH_SOURCE,
					Data: api.MessageCallbackData{},
				}
				if slices.Contains(session.Jobs.SteamNews, commandResponse.GuildID) {
					callbackMessage.Data.Content = "This guild already have an active steam-news job"
					api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
					return
				}
				callbackMessage.Data.Content = "Created a steam-news channel"
				api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)

				go jobs.SteamNewsFeed(ctx, session, commandResponse.GuildID)
				session.Jobs.SteamNews = append(session.Jobs.SteamNews, commandResponse.GuildID)

			case STEAM_SALES:
				callbackMessage := api.MessageCallback{
					Type: api.CHANNEL_MESSAGE_WITH_SOURCE,
					Data: api.MessageCallbackData{},
				}
				if slices.Contains(session.Jobs.SteamSales, commandResponse.GuildID) {
					callbackMessage.Data.Content = "This guild already have an active steam-sales job"
					api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)
					return
				}
				callbackMessage.Data.Content = "Created a steam-sales channel"
				api.InteractionMsgCallback(session, commandResponse.ID, commandResponse.Token, callbackMessage)

				go jobs.SteamSalesFeed(ctx, session, commandResponse.GuildID)
				session.Jobs.SteamSales = append(session.Jobs.SteamSales, commandResponse.GuildID)
			}
		}
	}
}
