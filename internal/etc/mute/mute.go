package mute

import (
	"slices"
	"snorp/internal/api"
	"snorp/internal/state"
)

func Messages(session *state.SessionState, user string, channelID string, messageID string) {
	if slices.Contains(session.Config.Bot.Mute.Users, user) {
		api.DeleteMessage(session, channelID, messageID)
	}
}
