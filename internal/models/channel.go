package models

type ChannelType int

const (
	ChannelTypeGuildText          ChannelType = 0
	ChannelTypeDM                 ChannelType = 1
	ChannelTypeGuildVoice         ChannelType = 2
	ChannelTypeGroupDM            ChannelType = 3
	ChannelTypeGuildCategory      ChannelType = 4
	ChannelTypeGuildNews          ChannelType = 5
	ChannelTypeGuildStore         ChannelType = 6
	ChannelTypeGuildNewsThread    ChannelType = 10
	ChannelTypeGuildPublicThread  ChannelType = 11
	ChannelTypeGuildPrivateThread ChannelType = 12
	ChannelTypeGuildStageVoice    ChannelType = 13
	ChannelTypeGuildDirectory     ChannelType = 14
	ChannelTypeGuildForum         ChannelType = 15
	ChannelTypeGuildMedia         ChannelType = 16
)

type Channel struct {
	Version              int64       `json:"version"`
	Type                 ChannelType `json:"type"`
	Position             int         `json:"position"`
	PermissionOverwrites []*any      `json:"permission_overwrites"`
	Name                 string      `json:"name"`
	ID                   string      `json:"id"`
	Flags                int         `json:"flags"`
	Topic                any         `json:"topic,omitempty"`
	RateLimitPerUser     int         `json:"rate_limit_per_user,omitempty"`
	ParentID             string      `json:"parent_id,omitempty"`
	LastMessageID        string      `json:"last_message_id,omitempty"`
	IconEmoji            IconEmoji   `json:"icon_emoji"`
	UserLimit            int         `json:"user_limit,omitempty"`
	RtcRegion            any         `json:"rtc_region,omitempty"`
	Bitrate              int         `json:"bitrate,omitempty"`
}

type IconEmoji struct {
	Name string `json:"name"`
	ID   any    `json:"id"`
}
