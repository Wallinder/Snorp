package channel

import (
	"encoding/json"
	"io"
	"log"
	"snorp/internal/api"
	"snorp/internal/state"
)

const (
	GUILD_TEXT          = 0  //a text channel within a server
	DM                  = 1  //a direct message between users
	GUILD_VOICE         = 2  //a voice channel within a server
	GROUP_DM            = 3  //a direct message between multiple users
	GUILD_CATEGORY      = 4  //an organizational category that contains up to 50 channels
	GUILD_ANNOUNCEMENT  = 5  //a channel that users can follow and crosspost into their own server (formerly news channels)
	ANNOUNCEMENT_THREAD = 10 //a temporary sub-channel within a GUILD_ANNOUNCEMENT channel
	PUBLIC_THREAD       = 11 //a temporary sub-channel within a GUILD_TEXT or GUILD_FORUM channel
	PRIVATE_THREAD      = 12 //a temporary sub-channel within a GUILD_TEXT channel that is only viewable by those invited and those with the MANAGE_THREADS permission
	GUILD_STAGE_VOICE   = 13 //a voice channel for hosting events with an audience
	GUILD_DIRECTORY     = 14 //the channel in a hub containing the listed servers
	GUILD_FORUM         = 15 //Channel that can only contain threads
	GUILD_MEDIA         = 16 //Channel that can only contain threads, similar to GUILD_FORUM channels
)

func Create(session *state.SessionState, guild api.Guild) {
	for _, channel := range guild.Channels {
		if channel.Name == "Snorp" && channel.Type == 4 {
			return
		}
	}

	categoryChannel := &api.GuildChannels{
		Name: "Snorp",
		Type: GUILD_CATEGORY,
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
		Type: GUILD_TEXT,
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
