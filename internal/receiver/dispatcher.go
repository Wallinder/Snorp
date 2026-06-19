package receiver

import (
	"context"
	"encoding/json"
	"snorp/internal/program"
	"snorp/pkg/discord"
	"sync"
)

func StartDispatchReader(ctx context.Context, app *program.Application, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-app.Discord.DispatchChan:
				go dispatchReader(ctx, app, message)
			}
		}
	})
}

func dispatchReader(ctx context.Context, app *program.Application, message discord.DispatchMessage) {
	switch message.Type {

	case "GUILD_CREATE":
		var guild discord.Guild
		if err := json.Unmarshal(message.Data, &guild); err != nil {
			app.ErrorChan <- program.Errors{Origin: "dispatcher", Err: err, Fatal: false}
			return
		}

	case "INTERACTION_CREATE":
		var interaction discord.Interaction
		if err := json.Unmarshal(message.Data, &interaction); err != nil {
			app.ErrorChan <- program.Errors{Origin: "dispatcher", Err: err, Fatal: false}
			return
		}
		interactionHandler(ctx, app, interaction)
	}
}
