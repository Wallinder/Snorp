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
)

func RegisterCommands(ctx context.Context, session *state.SessionState) {
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
				Name:        "regnummer",
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
}
