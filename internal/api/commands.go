package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"snorp/internal/state"
)

type ApplicationCommand struct {
	ID                       string                     `json:"id,omitzero"`
	Type                     int                        `json:"type,omitzero"`
	ApplicationID            string                     `json:"application_id,omitzero"`
	GuildID                  string                     `json:"guild_id,omitzero"`
	Name                     string                     `json:"name,omitzero"`
	NameLocalizations        []string                   `json:"name_localizations,omitzero"`
	Description              string                     `json:"description,omitzero"`
	DescriptionLocalizations []string                   `json:"description_localizations,omitzero"`
	Options                  []ApplicationCommandOption `json:"options,omitzero"`
	DefaultMemberPermissions string                     `json:"default_member_permissions,omitzero"`
	DmPermissions            bool                       `json:"dm_permission,omitzero"`
	DefaultPermissions       bool                       `json:"default_permission,omitzero"`
	Nsfw                     bool                       `json:"nsfw,omitzero"`
	IntegrationTypes         [2]int                     `json:"integration_types,omitzero"`
	Contexts                 [3]int                     `json:"contexts,omitzero"`
	Version                  string                     `json:"version,omitzero"`
	Handler                  int                        `json:"handler,omitzero"`
}

type ApplicationCommandOption struct {
	Type                     int      `json:"type,omitzero"`
	Name                     string   `json:"name,omitzero"`
	NameLocalizations        []string `json:"name_localizations,omitzero"`
	Description              string   `json:"description,omitzero"`
	DescriptionLocalizations []string `json:"description_localizations,omitzero"`
	Required                 bool     `json:"required,omitzero"`
	Choices                  []string `json:"choices,omitzero"`
	Options                  []int    `json:"options,omitzero"`
	ChannelTypes             []int    `json:"channel_types,omitzero"`
	MinValue                 int      `json:"min_value,omitzero"`
	MaxValue                 int      `json:"max_value,omitzero"`
	MinLength                int      `json:"min_length,omitzero"`
	MaxLength                int      `json:"max_length,omitzero"`
	Autocomplete             bool     `json:"autocomplete,omitzero"`
}

const (
	SUB_COMMAND       = 1
	SUB_COMMAND_GROUP = 2
	STRING            = 3
	INTEGER           = 4
	BOOLEAN           = 5
	USER              = 6
	CHANNEL           = 7
	ROLE              = 8
	MENTIONABLE       = 9
	NUMBER            = 10
	ATTACHMENT        = 11
)

const (
	CHAT_INPUT          = 1 //Slash commands; a text-based command that shows up when a user types /
	USER_COMMAND        = 2 //A UI-based command that shows up when you right click or tap on a user
	MESSAGE             = 3 //A UI-based command that shows up when you right click or tap on a message
	PRIMARY_ENTRY_POINT = 4 //A UI-based command that represents the primary way to invoke an app's Activity
)

const (
	GUILD_INSTALL = 0 //App is installable to servers
	USER_INSTALL  = 1 //App is installable to users

	GUILD           = 0 //Interaction can be used within servers
	BOT_DM          = 1 //Interaction can be used within DMs with the app's bot user
	PRIVATE_CHANNEL = 2 //Interaction can be used within Group DMs and DMs other than the app's bot user

	APP_HANDLER             = 1 //The app handles the interaction using an interaction token
	DISCORD_LAUNCH_ACTIVITY = 2 //Discord handles the interaction by launching an Activity and sending a follow-up message without coordinating with the app
)

func GetGlobalCommand(session *state.SessionState) ([]ApplicationCommand, error) {
	request := state.HttpRequest{
		Method: "GET",
		Uri:    fmt.Sprintf("/applications/%s/commands", session.ReadyData.User.ID),
		Body:   nil,
	}

	response, err := session.SendRequest(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var commands []ApplicationCommand

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &commands)
	if err != nil {
		return nil, err
	}

	return commands, nil
}

func DeleteGlobalCommand(session *state.SessionState, commandID string) {
	request := state.HttpRequest{
		Method: "DELETE",
		Uri:    fmt.Sprintf("/applications/%s/commands/%s", session.ReadyData.User.ID, commandID),
		Body:   nil,
	}

	_, err := session.SendRequest(request)
	if err != nil {
		log.Printf("Error deleting command: %v\n", err)
	}
}

func CreateGlobalCommand(session *state.SessionState, command *ApplicationCommand) (*ApplicationCommand, error) {
	jsonData, err := json.Marshal(command)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonData)

	request := state.HttpRequest{
		Method: "POST",
		Uri:    fmt.Sprintf("/applications/%s/commands", session.ReadyData.User.ID),
		Body:   reader,
	}

	response, err := session.SendRequest(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var newCommand *ApplicationCommand

	err = json.Unmarshal(body, &newCommand)
	if err != nil {
		return nil, err
	}

	return newCommand, nil
}
