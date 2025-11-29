package interaction

import (
	"context"
	"log"
	"snorp/internal/api"
	"snorp/internal/state"
)

const (
	ARCHIVE_MESSAGE = "Archive Message"
)

const (
	SVV           = "svv"
	SVV_REGNUMMER = "regnummer"

	CREATE         = "create"
	CREATE_CHANNEL = "channel"

	CHANNEL_WELCOME     = "welcome-channel"
	CHANNEL_STEAM_NEWS  = "steam_news"
	CHANNEL_STEAM_SALES = "steam_sales"
)

func RegisterCommands(ctx context.Context, session *state.SessionState) {
	//commands, _ := api.GetGlobalCommand(session)
	//for _, command := range commands {
	//	fmt.Println(command)
	//}
	saveMessageCommand := &api.ApplicationCommand{
		Name: ARCHIVE_MESSAGE,
		Type: api.MESSAGE,
	}
	_, err := api.CreateGlobalCommand(session, saveMessageCommand)
	if err != nil {
		log.Println(err)
		return
	}

	svv := &api.ApplicationCommand{
		Name:        SVV,
		Type:        api.CHAT_INPUT,
		Description: "Statens Vegvesen",
		Options: []api.ApplicationCommandOption{
			{
				Name:        SVV_REGNUMMER,
				Description: "Hent Kjøretøydata",
				Type:        3,
				Required:    true,
			},
		},
	}
	_, err = api.CreateGlobalCommand(session, svv)
	if err != nil {
		log.Println(err)
		return
	}

	welcome := &api.ApplicationCommand{
		Name:        CREATE,
		Type:        api.CHAT_INPUT,
		Description: "Create resource",
		Options: []api.ApplicationCommandOption{
			{
				Name:        CREATE_CHANNEL,
				Description: "channel",
				Type:        3,
				Required:    false,
				Choices: []api.ApplicationCommandChoices{
					{
						Name:  "welcome",
						Value: CHANNEL_WELCOME,
					},
					{
						Name:  "steam-news",
						Value: CHANNEL_STEAM_NEWS,
					},
					{
						Name:  "steam-sales",
						Value: CHANNEL_STEAM_SALES,
					},
				},
			},
		},
	}
	_, err = api.CreateGlobalCommand(session, welcome)
	if err != nil {
		log.Println(err)
		return
	}
}
