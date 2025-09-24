package action

import "time"

type Message struct {
	Type            int           `json:"type"`
	Tts             bool          `json:"tts"`
	Timestamp       time.Time     `json:"timestamp"`
	Pinned          bool          `json:"pinned"`
	Nonce           string        `json:"nonce"`
	Mentions        []any         `json:"mentions"`
	MentionRoles    []any         `json:"mention_roles"`
	MentionEveryone bool          `json:"mention_everyone"`
	Member          MessageMember `json:"member"`
	ID              string        `json:"id"`
	Flags           int           `json:"flags"`
	Embeds          []any         `json:"embeds"`
	EditedTimestamp any           `json:"edited_timestamp"`
	Content         string        `json:"content"`
	Components      []any         `json:"components"`
	ChannelType     int           `json:"channel_type"`
	ChannelID       string        `json:"channel_id"`
	Author          MessageAuthor `json:"author"`
	Attachments     []any         `json:"attachments"`
	GuildID         string        `json:"guild_id"`
}

type MessageMember struct {
	Roles                      []any     `json:"roles"`
	PremiumSince               any       `json:"premium_since"`
	Pending                    bool      `json:"pending"`
	Nick                       any       `json:"nick"`
	Mute                       bool      `json:"mute"`
	JoinedAt                   time.Time `json:"joined_at"`
	Flags                      int       `json:"flags"`
	Deaf                       bool      `json:"deaf"`
	CommunicationDisabledUntil any       `json:"communication_disabled_until"`
	Banner                     any       `json:"banner"`
	Avatar                     any       `json:"avatar"`
}

type MessageAuthor struct {
	Username             string                    `json:"username"`
	PublicFlags          int                       `json:"public_flags"`
	PrimaryGuild         MessageAuthorPrimaryGuild `json:"primary_guild"`
	ID                   string                    `json:"id"`
	GlobalName           string                    `json:"global_name"`
	DisplayNameStyles    any                       `json:"display_name_styles"`
	Discriminator        string                    `json:"discriminator"`
	Collectibles         any                       `json:"collectibles"`
	Clan                 MessageAuthorClan         `json:"clan"`
	AvatarDecorationData any                       `json:"avatar_decoration_data"`
	Avatar               string                    `json:"avatar"`
}

type MessageAuthorPrimaryGuild struct {
	Tag             string `json:"tag"`
	IdentityGuildID string `json:"identity_guild_id"`
	IdentityEnabled bool   `json:"identity_enabled"`
	Badge           string `json:"badge"`
}

type MessageAuthorClan struct {
	Tag             string `json:"tag"`
	IdentityGuildID string `json:"identity_guild_id"`
	IdentityEnabled bool   `json:"identity_enabled"`
	Badge           string `json:"badge"`
}
