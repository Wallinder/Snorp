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

	categoryChannel := &api.NewChannel{
		Name:     "Snorp",
		Type:     4,
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

	adminChannel := &api.NewChannel{
		Name:     "Admin",
		Type:     0,
		Position: 0,
		ParentID: channel.ID,
	}

	_, err = adminChannel.CreateChannel(session, guild.ID)
	if err != nil {
		log.Printf("Error creating admin channel: %s\n", err)
	}
}
