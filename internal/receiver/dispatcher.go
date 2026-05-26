package receiver

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/internal/state"
	"snorp/pkg/discord"
	"sync"
)

func StartDispatchReader(ctx context.Context, session *state.SessionState, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-session.Discord.DispatchChan:
				go dispatchReader(ctx, session, message)
			}
		}
	})
}

func dispatchReader(ctx context.Context, session *state.SessionState, message discord.DispatchMessage) {
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
