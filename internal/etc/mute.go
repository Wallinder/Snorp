package etc

import (
	"slices"
	"snorp/internal/api"
	"snorp/internal/state"
)

func Messages(session *state.SessionState, message api.Message) {
	if slices.Contains(session.Config.Bot.Mute.Users, message.Author.Username) {
		api.DeleteMessage(session, message.ChannelID, message.ID)
	}
}
