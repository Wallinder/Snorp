package receiver

import (
	"context"
	"encoding/json"
	"snorp/pkg/discord"
	"sync"
)

type DispatcherService struct {
	Discord *discord.DiscordService
	ErrChan chan error
}

func (s *DispatcherService) Name() string {
	return "dispatcher"
}

func (s *DispatcherService) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-s.Discord.DispatchChan:
				go s.dispatchReader(ctx, message)
			}
		}
	})
}

func (s *DispatcherService) dispatchReader(ctx context.Context, message discord.DispatchMessage) {
	switch message.Type {

	case "GUILD_CREATE":
		var guild discord.Guild
		if err := json.Unmarshal(message.Data, &guild); err != nil {
			s.ErrChan <- err
			return
		}

	case "INTERACTION_CREATE":
		var interaction discord.Interaction
		if err := json.Unmarshal(message.Data, &interaction); err != nil {
			s.ErrChan <- err
			return
		}
		s.interactionHandler(ctx, interaction)
	}
}
