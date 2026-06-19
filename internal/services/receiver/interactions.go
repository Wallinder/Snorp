package receiver

import (
	"bytes"
	"context"
	"encoding/json"
	"snorp/pkg/discord"
)

func (s *DispatcherService) interactionHandler(_ context.Context, interaction discord.Interaction) {
	switch interaction.Data.Name {
	}
}

func (s *DispatcherService) interactionCallback(interaction discord.Interaction, callback discord.InteractionCallback) {
	uri := "/interactions/" + interaction.ID + "/" + interaction.Token + "/callback"

	data, err := json.Marshal(callback)
	if err != nil {
		s.ErrChan <- err
	}

	_, err = s.Discord.NewDiscordRequest("POST", uri, bytes.NewReader(data))
	if err != nil {
		s.ErrChan <- err
	}
}
