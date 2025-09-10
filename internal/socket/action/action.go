package action

import (
	"encoding/json"
	"log"
	"snorp/internal/state"
)

func DispatchHandler(session *state.SessionState, action string, dispatchMessage json.RawMessage) {
	switch action {

	case "READY":
		log.Println("Handshake completed..")
		var readyData state.ReadyData
		err := json.Unmarshal(dispatchMessage, &readyData)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		session.ReadyData = readyData

	case "RESUMED":
		log.Println("Connection successfully resumed..")
	}
}
