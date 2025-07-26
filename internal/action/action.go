package action

import (
	"encoding/json"
	"log"
	"menial/internal/state"
)

func DispatchHandler(sessionState *state.SessionState, action string, dispatchMessage json.RawMessage) {
	switch action {

	case "READY":
		var readyData state.ReadyData
		err := json.Unmarshal(dispatchMessage, &readyData)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}
		sessionState.ReadyData = readyData

	case "RESUMED":
		log.Println("Connection successfully resumed..")
	}
}
