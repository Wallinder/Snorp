package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"snorp/internal/state"
	"time"
)

type Message struct {
	Type            int           `json:"type,omitzero"`
	Tts             bool          `json:"tts,omitzero"`
	Timestamp       time.Time     `json:"timestamp,omitzero"`
	Pinned          bool          `json:"pinned,omitzero"`
	Nonce           string        `json:"nonce,omitzero"`
	Mentions        []any         `json:"mentions,omitzero"`
	MentionRoles    []any         `json:"mention_roles,omitzero"`
	MentionEveryone bool          `json:"mention_everyone,omitzero"`
	Member          MessageMember `json:"member,omitzero"`
	ID              string        `json:"id,omitzero"`
	Flags           int           `json:"flags,omitzero"`
	Embeds          []any         `json:"embeds,omitzero"`
	EditedTimestamp any           `json:"edited_timestamp,omitzero"`
	Content         string        `json:"content,omitzero"`
	Components      []any         `json:"components,omitzero"`
	ChannelType     int           `json:"channel_type,omitzero"`
	ChannelID       string        `json:"channel_id,omitzero"`
	Author          MessageAuthor `json:"author,omitzero"`
	Attachments     []any         `json:"attachments,omitzero"`
	GuildID         string        `json:"guild_id,omitzero"`
}

type MessageMember struct {
	Roles                      []any     `json:"roles,omitzero"`
	PremiumSince               any       `json:"premium_since,omitzero"`
	Pending                    bool      `json:"pending,omitzero"`
	Nick                       any       `json:"nick,omitzero"`
	Mute                       bool      `json:"mute,omitzero"`
	JoinedAt                   time.Time `json:"joined_at,omitzero"`
	Flags                      int       `json:"flags,omitzero"`
	Deaf                       bool      `json:"deaf,omitzero"`
	CommunicationDisabledUntil any       `json:"communication_disabled_until,omitzero"`
	Banner                     any       `json:"banner,omitzero"`
	Avatar                     any       `json:"avatar,omitzero"`
}

type MessageAuthor struct {
	Username             string                    `json:"username,omitzero"`
	PublicFlags          int                       `json:"public_flags,omitzero"`
	PrimaryGuild         MessageAuthorPrimaryGuild `json:"primary_guild,omitzero"`
	ID                   string                    `json:"id,omitzero"`
	GlobalName           string                    `json:"global_name,omitzero"`
	DisplayNameStyles    any                       `json:"display_name_styles,omitzero"`
	Discriminator        string                    `json:"discriminator,omitzero"`
	Collectibles         any                       `json:"collectibles,omitzero"`
	Clan                 MessageAuthorClan         `json:"clan,omitzero"`
	AvatarDecorationData any                       `json:"avatar_decoration_data,omitzero"`
	Avatar               string                    `json:"avatar,omitzero"`
}

type MessageAuthorPrimaryGuild struct {
	Tag             string `json:"tag,omitzero"`
	IdentityGuildID string `json:"identity_guild_id,omitzero"`
	IdentityEnabled bool   `json:"identity_enabled,omitzero"`
	Badge           string `json:"badge,omitzero"`
}

type MessageAuthorClan struct {
	Tag             string `json:"tag,omitzero"`
	IdentityGuildID string `json:"identity_guild_id,omitzero"`
	IdentityEnabled bool   `json:"identity_enabled,omitzero"`
	Badge           string `json:"badge,omitzero"`
}

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
