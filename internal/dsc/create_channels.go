package dsc

import (
	"log"
	"snorp/internal/api"
	"snorp/internal/state"
)

func CreateDesiredGuildChannels(session *state.SessionState, guildID string, guildOwnerID string) *api.Message {
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
		log.Printf("Error creating category channel: %s\n", err)
		return nil
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

	newAdminChannel, err := api.CreateGuildChannel(session, guildID, adminChannel)
	if err != nil {
		log.Printf("Error creating admin channel: %s\n", err)
		return nil
	}

	message := api.Message{
		Content: "TEST",
	}

	newMessage, err := api.CreateMessage(session, newAdminChannel.ID, message)
	if err != nil {
		log.Printf("Error creating message: %s\n", err)
		return nil
	}

	return newMessage
}
