package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"snorp/internal/state"
)

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

	CHAT_INPUT          = 1 //Slash commands; a text-based command that shows up when a user types /
	USER_COMMAND        = 2 //A UI-based command that shows up when you right click or tap on a user
	MESSAGE             = 3 //A UI-based command that shows up when you right click or tap on a message
	PRIMARY_ENTRY_POINT = 4 //A UI-based command that represents the primary way to invoke an app's Activity

	GUILD_INSTALL = 0 //App is installable to servers
	USER_INSTALL  = 1 //App is installable to users

	GUILD           = 0 //Interaction can be used within servers
	BOT_DM          = 1 //Interaction can be used within DMs with the app's bot user
	PRIVATE_CHANNEL = 2 //Interaction can be used within Group DMs and DMs other than the app's bot user

	APP_HANDLER             = 1 //The app handles the interaction using an interaction token
	DISCORD_LAUNCH_ACTIVITY = 2 //Discord handles the interaction by launching an Activity and sending a follow-up message without coordinating with the app
)

func GetGlobalCommand(session *state.SessionState) []string {
	request := state.HttpRequest{
		Method: "GET",
		Uri:    fmt.Sprintf("/applications/%s/commands", session.ReadyData.User.ID),
		Body:   nil,
	}

	response, err := session.SendRequest(request)
	if err != nil {
		log.Printf("Error fetching commands: %v\n", err)
	}
	defer response.Body.Close()

	var commands []string

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &commands)
	if err != nil {
		log.Fatalf("Error unmarshaling json: %v", err)
	}

	return commands
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

func RegisterGlobalCommand(session *state.SessionState, command *ApplicationCommand) {
	jsonData, err := json.Marshal(command)
	if err != nil {
		log.Printf("Error marshaling command: %v\n", err)
	}

	reader := bytes.NewReader(jsonData)

	request := state.HttpRequest{
		Method: "POST",
		Uri:    fmt.Sprintf("/applications/%s/commands", session.ReadyData.User.ID),
		Body:   reader,
	}

	_, err = session.SendRequest(request)
	if err != nil {
		log.Printf("Error creating command: %v\n", err)
	}
}
