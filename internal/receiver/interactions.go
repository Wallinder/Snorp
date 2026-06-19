package receiver

import (
	"bytes"
	"context"
	"encoding/json"
	"snorp/internal/program"
	"snorp/pkg/discord"
)

func interactionHandler(_ context.Context, app *program.Application, interaction discord.Interaction) {
	switch interaction.Data.Name {
	}
}

func interactionCallback(app *program.Application, interaction discord.Interaction, callback discord.InteractionCallback) {
	uri := "/interactions/" + interaction.ID + "/" + interaction.Token + "/callback"

	data, err := json.Marshal(callback)
	if err != nil {
		app.ErrorChan <- program.Errors{Origin: "interaction", Err: err, Fatal: false}
	}

	_, err = app.Discord.NewDiscordRequest("POST", uri, bytes.NewReader(data))
	if err != nil {
		app.ErrorChan <- program.Errors{Origin: "interaction", Err: err, Fatal: false}
	}
}
