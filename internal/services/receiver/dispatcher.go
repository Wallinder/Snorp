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

	case "GUILD_MEMBER_ADD":
		var member *discord.Member
		if err := json.Unmarshal(message.Data, &member); err != nil {
			s.ErrChan <- err
			return
		}
		if err := s.Storage.SaveMember(member, member.GuildID); err != nil {
			s.ErrChan <- err
		}

	case "GUILD_ROLE_CREATE":
		var data discord.GuildRoleCreate
		if err := json.Unmarshal(message.Data, &data); err != nil {
			s.ErrChan <- err
			return
		}
		if err := s.Storage.SaveRole(data.Role, data.GuildID); err != nil {
			s.ErrChan <- err
		}

	case "GUILD_ROLE_DELETE":
		var role map[string]string
		if err := json.Unmarshal(message.Data, &role); err != nil {

		}
		if err := s.Storage.DeleteRole(role["role_id"], role["guild_id"]); err != nil {
			s.ErrChan <- err
		}

	case "CHANNEL_CREATE":
		var channel *discord.Channel
		if err := json.Unmarshal(message.Data, &channel); err != nil {
			s.ErrChan <- err
			return
		}
		if err := s.Storage.SaveChannel(channel, channel.GuildID); err != nil {
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
