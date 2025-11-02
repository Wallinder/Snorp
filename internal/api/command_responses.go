package api

import (
	"encoding/json"
	"time"
)

type CommandResponse struct {
	ApplicationID  string                `json:"application_id"`
	ChannelID      string                `json:"channel_id"`
	Data           CommandResponseData   `json:"data"`
	GuildID        string                `json:"guild_id"`
	GuildLocale    string                `json:"guild_locale"`
	AppPermissions string                `json:"app_permissions"`
	ID             string                `json:"id"`
	Locale         string                `json:"locale"`
	Member         CommandResponseMember `json:"member"`
	Token          string                `json:"token"`
	Type           int                   `json:"type"`
	Version        int                   `json:"version"`
}

type CommandResponseData struct {
	ID       string                       `json:"id"`
	Name     string                       `json:"name"`
	TargetID string                       `json:"target_id,omitzero"`
	Resolved CommandResponseResolved      `json:"resolved,omitzero"`
	Options  []CommandResponseDataOptions `json:"options,omitzero"`
	Type     int                          `json:"type"`
}

type CommandResponseDataOptions struct {
	Type  int    `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CommandResponseMember struct {
	Avatar       any                 `json:"avatar"`
	Deaf         bool                `json:"deaf"`
	IsPending    bool                `json:"is_pending"`
	JoinedAt     time.Time           `json:"joined_at"`
	Mute         bool                `json:"mute"`
	Nick         any                 `json:"nick"`
	Pending      bool                `json:"pending"`
	Permissions  string              `json:"permissions"`
	PremiumSince any                 `json:"premium_since"`
	Roles        []string            `json:"roles"`
	User         CommandResponseUser `json:"user"`
}

type CommandResponseUser struct {
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	ID            string `json:"id"`
	PublicFlags   int    `json:"public_flags"`
	Username      string `json:"username"`
}

type CommandResponseResolved struct {
	Messages json.RawMessage `json:"messages"`
	Members  json.RawMessage `json:"members"`
	Users    json.RawMessage `json:"users"`
}

type CommandResponseAuthor struct {
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	ID            string `json:"id"`
	PublicFlags   int    `json:"public_flags"`
	Username      string `json:"username"`
}
