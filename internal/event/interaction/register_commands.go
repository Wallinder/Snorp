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

	STEAM       = "steam"
	STEAM_NEWS  = "news"
	STEAM_SALES = "sales"
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

	steamNews := &api.ApplicationCommand{
		Name:        STEAM,
		Type:        api.CHAT_INPUT,
		Description: "Steam Operations",
		Options: []api.ApplicationCommandOption{
			{
				Name:        STEAM_SALES,
				Description: "Creates a steam-sales channel",
				Type:        1,
				Required:    false,
			},
			{
				Name:        STEAM_NEWS,
				Description: "Creates a steam-news channel",
				Type:        1,
				Required:    false,
			},
		},
	}
	_, err = api.CreateGlobalCommand(session, steamNews)
	if err != nil {
		log.Println(err)
		return
	}
}
