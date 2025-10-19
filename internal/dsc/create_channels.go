package dsc

import (
	"snorp/internal/api"
	"snorp/internal/state"
)

func CreateAdminChannel(session *state.SessionState, guildID string) error {
	adminChannel := &api.GuildChannels{
		Name: "snorp-admin",
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
		Topic:    "snorp:admin",
		Position: 0,
	}

	superUser := api.GuildChannelsPermissions{
		ID:    session.Config.Bot.SuperuserID,
		Type:  1,
		Allow: "1024",
		Deny:  "0",
	}
	adminChannel.Permissions = append(adminChannel.Permissions, superUser)

	_, err := api.CreateGuildChannel(session, guildID, adminChannel)
	if err != nil {
		return err
	}

	return nil
}
