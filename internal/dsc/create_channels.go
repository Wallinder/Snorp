package dsc

import (
	"snorp/internal/api"
	"snorp/internal/state"
)

func CreateDesiredGuildChannels(session *state.SessionState, guildID string, guildOwnerID string) error {
	categoryChannel := &api.GuildChannels{
		Name: "Snorp",
		Type: api.GUILD_CATEGORY,
		Permissions: []api.GuildChannelsPermissions{
			{
				ID:    guildID,
				Type:  0,
				Allow: "1024",
				Deny:  "0",
			},
		},
		Position: 0,
	}

	newCategoryChannel, err := api.CreateGuildChannel(session, guildID, categoryChannel)
	if err != nil {
		return err
	}

	adminChannel := &api.GuildChannels{
		Name: "Admin",
		Type: api.GUILD_TEXT,
		Permissions: []api.GuildChannelsPermissions{
			{
				ID:    session.ReadyData.User.ID,
				Type:  1,
				Allow: "1024",
				Deny:  "0",
			},
			{
				ID:    guildID,
				Type:  0,
				Allow: "0",
				Deny:  "1024",
			},
		},
		Position: 0,
		ParentID: newCategoryChannel.ID,
	}

	if session.Config.Bot.SuperuserID != guildOwnerID {
		superUser := api.GuildChannelsPermissions{
			ID:    session.Config.Bot.SuperuserID,
			Type:  1,
			Allow: "1024",
			Deny:  "0",
		}
		adminChannel.Permissions = append(adminChannel.Permissions, superUser)
	}

	_, err = api.CreateGuildChannel(session, guildID, adminChannel)
	if err != nil {
		return err
	}

	return nil
}
