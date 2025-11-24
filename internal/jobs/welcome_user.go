package jobs

import (
	"fmt"
	"log"
	"snorp/internal/api"
	"snorp/internal/state"
)

func Welcome(session *state.SessionState, member api.GuildMembers) {
	channel := &api.GuildChannels{
		Name:  "new-phone-who-dis",
		Type:  api.GUILD_TEXT,
		Topic: "snorp:welcome",
	}

	channelID, err := api.FindOrCreateChannel(session, channel, member.GuildID)
	if err != nil {
		log.Println(err)
		return
	}

	newMessage := api.Message{
		Content: fmt.Sprintf("Welcome to %s", member.User.GlobalName),
	}

	_, err = api.CreateMessage(session, channelID, newMessage)
	if err != nil {
		log.Println(err)
		return
	}
}
