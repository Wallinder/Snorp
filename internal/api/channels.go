package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func CreateGuildChannel(session *state.SessionState, guildID string, channel *GuildChannels) (GuildChannels, error) {
	var newChannel GuildChannels

	body, err := json.Marshal(channel)
	if err != nil {
		return newChannel, err
	}

	request := state.HttpRequest{
		Method: "POST",
		Uri:    fmt.Sprintf("/guilds/%s/channels", guildID),
		Body:   bytes.NewBuffer(body),
	}

	response, err := session.SendRequest(request)
	if err != nil {
		return newChannel, err
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return newChannel, err
	}
	defer response.Body.Close()

	err = json.Unmarshal(body, &newChannel)
	if err != nil {
		return newChannel, err
	}

	return newChannel, nil
}

func GetGuildChannels(session *state.SessionState, guildID string) ([]GuildChannels, error) {
	request := state.HttpRequest{
		Method: "GET",
		Uri:    fmt.Sprintf("/guilds/%s/channels", guildID),
		Body:   nil,
	}

	response, err := session.SendRequest(request)
	if err != nil {
		return nil, err
	}

	var channels []GuildChannels

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.Unmarshal(body, &channels)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
