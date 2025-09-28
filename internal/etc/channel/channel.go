package channel

import (
	"encoding/json"
	"io"
	"log"
	"snorp/internal/api"
	"snorp/internal/state"
)

func Create(session *state.SessionState, guild api.Guild) {
	for _, channel := range guild.Channels {
		if channel.Name == "Snorp" && channel.Type == 4 {
			return
		}
	}

	categoryChannel := &api.GuildChannels{
		Name: "Snorp",
		Type: 4,
		Permissions: []api.GuildChannelsPermissions{
			{
				ID:    guild.ID,
				Type:  0,
				Allow: "1024",
				Deny:  "0",
			},
		},
		Position: 0,
	}

	response, err := categoryChannel.CreateChannel(session, guild.ID)
	if err != nil {
		log.Printf("Error creating category channel: %s\n", err)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	var channel api.GuildChannels

	err = json.Unmarshal(body, &channel)
	if err != nil {
		log.Println(err)
		return
	}

	adminChannel := &api.GuildChannels{
		Name: "Admin",
		Type: 0,
		Permissions: []api.GuildChannelsPermissions{
			{
				ID:    session.ReadyData.User.ID,
				Type:  1,
				Allow: "1024",
				Deny:  "0",
			},
			{
				ID:    guild.ID,
				Type:  0,
				Allow: "0",
				Deny:  "1024",
			},
		},
		Position: 0,
		ParentID: channel.ID,
	}

	if session.Config.Bot.SuperuserID != guild.OwnerID {
		superUser := api.GuildChannelsPermissions{
			ID:    session.Config.Bot.SuperuserID,
			Type:  1,
			Allow: "1024",
			Deny:  "0",
		}
		adminChannel.Permissions = append(adminChannel.Permissions, superUser)
	}

	_, err = adminChannel.CreateChannel(session, guild.ID)
	if err != nil {
		log.Printf("Error creating admin channel: %s\n", err)
	}
}
