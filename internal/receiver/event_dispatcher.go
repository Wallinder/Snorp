package receiver

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/internal/state"
	"snorp/pkg/discord"
)

func dispatcher(ctx context.Context, session *state.SessionState) {
	message := <-session.Discord.DispatchChan

	switch message.Type {

	case "GUILD_CREATE":
		var guild discord.Guild
		if err := json.Unmarshal(message.Data, &guild); err != nil {
			slog.Info("failed to unmarshal json", "error", err)
			return
		}

	case "INTERACTION_CREATE":
		var interaction discord.Interaction
		if err := json.Unmarshal(message.Data, &interaction); err != nil {
			slog.Info("failed to unmarshal json", "error", err)
			return
		}
		interactionHandler(ctx, session, interaction)
	}
}
