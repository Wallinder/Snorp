package models

import "time"

type Member struct {
	User                       User      `json:"user"`
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

type User struct {
	Username             string `json:"username"`
	PublicFlags          int    `json:"public_flags"`
	PrimaryGuild         any    `json:"primary_guild"`
	ID                   string `json:"id"`
	GlobalName           string `json:"global_name"`
	DisplayNameStyles    any    `json:"display_name_styles"`
	DisplayName          string `json:"display_name"`
	Discriminator        string `json:"discriminator"`
	Collectibles         any    `json:"collectibles"`
	Bot                  bool   `json:"bot"`
	AvatarDecorationData any    `json:"avatar_decoration_data"`
	Avatar               string `json:"avatar"`
}
