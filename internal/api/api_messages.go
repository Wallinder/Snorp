package api

import (
	"fmt"
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
