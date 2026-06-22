package receiver

import (
	"context"
	"encoding/json"
	"snorp/internal/storage"
	"snorp/pkg/discord"
	"sync"
)

type DispatcherService struct {
	Discord *discord.Discord
	Storage storage.Storage
	ErrChan chan error
}

func NewDispatchService(discord *discord.Discord, storage storage.Storage) *DispatcherService {
	return &DispatcherService{
		Discord: discord,
		Storage: storage,
		ErrChan: make(chan error),
	}
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
		if err := s.Storage.SaveGuild(guild); err != nil {
			s.ErrChan <- err
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
