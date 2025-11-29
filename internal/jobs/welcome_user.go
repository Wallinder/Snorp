package jobs

import (
	"fmt"
	"log"
	"snorp/internal/api"
	"snorp/internal/state"
)

func WelcomeUser(session *state.SessionState, member api.GuildMembers) {
	if session.Jobs.Welcome[member.GuildID] == "" {
		fmt.Println("test")
		return
	}

	newMessage := api.Message{
		Content: fmt.Sprintf("Welcome to %s", member.User.DisplayName),
	}

	_, err := api.CreateMessage(session, session.Jobs.Welcome[member.GuildID], newMessage)
	if err != nil {
		log.Println(err)
		return
	}
}
