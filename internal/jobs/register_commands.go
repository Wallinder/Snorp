package jobs

import (
	"context"
	"log"
	"snorp/internal/api"
	"snorp/internal/state"
)

const (
	ARCHIVE_MESSAGE = "Archive Message"
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
}
