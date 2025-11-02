package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"snorp/internal/state"
)

// Interaction Callback Type
const (
	PONG                                    = 1  //ACK a Ping
	CHANNEL_MESSAGE_WITH_SOURCE             = 4  //Respond to an interaction with a message
	DEFERRED_CHANNEL_MESSAGE_WITH_SOURCE    = 5  //ACK an interaction and edit a response later, the user sees a loading state
	DEFERRED_UPDATE_MESSAGE                 = 6  //For components, ACK an interaction and edit the original message later; the user does not see a loading state
	UPDATE_MESSAGE                          = 7  //For components, edit the message the component was attached to
	APPLICATION_COMMAND_AUTOCOMPLETE_RESULT = 8  //Respond to an autocomplete interaction with suggested choices
	MODAL                                   = 9  //Respond to an interaction with a popup modal
	PREMIUM_REQUIRED                        = 10 //Deprecated; respond to an interaction with an upgrade button, only available for apps with monetization enabled
	LAUNCH_ACTIVITY                         = 12 //Launch the Activity associated with the app. Only available for apps with Activities enabled
)

type MessageCallback struct {
	Type int                 `json:"type"`
	Data MessageCallbackData `json:"data"`
}

type MessageCallbackData struct {
	Content string `json:"content"`
}

func InteractionMsgCallback(session *state.SessionState, interactionID string, interactionToken string, message MessageCallback) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(jsonData)

	request := state.HttpRequest{
		Method: "POST",
		Uri:    fmt.Sprintf("/interactions/%s/%s/callback", interactionID, interactionToken),
		Body:   reader,
	}

	_, err = session.SendRequest(request)
	if err != nil {
		return err
	}

	return nil
}
