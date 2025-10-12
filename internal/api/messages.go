package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"snorp/internal/state"
)

func DeleteMessage(session *state.SessionState, channelID string, messageID string) {
	request := state.HttpRequest{
		Method: "DELETE",
		Uri:    fmt.Sprintf("/channels/%s/messages/%s", channelID, messageID),
		Body:   nil,
	}

	_, err := session.SendRequest(request)
	if err != nil {
		log.Printf("Error deleting message %s: %v\n", messageID, err)
	}
}

func CreateMessage(session *state.SessionState, channelID string, message Message) (*Message, error) {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonData)

	request := state.HttpRequest{
		Method: "POST",
		Uri:    fmt.Sprintf("/channels/%s/messages", channelID),
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

	var newMessage *Message

	err = json.Unmarshal(body, &newMessage)
	if err != nil {
		return nil, err
	}

	return newMessage, nil
}
