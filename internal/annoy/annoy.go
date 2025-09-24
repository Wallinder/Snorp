package annoy

import (
	"slices"
)

func RemoveMessages(userlist []string, user string, guildID string, messageID string) {
	if slices.Contains(userlist, user) {
		//api.DeleteMessage(guildID, messageID)
	}
}
