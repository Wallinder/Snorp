package models

type Channel struct {
	Version              int64     `json:"version"`
	Type                 int       `json:"type"`
	Position             int       `json:"position"`
	PermissionOverwrites []*any    `json:"permission_overwrites"`
	Name                 string    `json:"name"`
	ID                   string    `json:"id"`
	Flags                int       `json:"flags"`
	Topic                any       `json:"topic,omitempty"`
	RateLimitPerUser     int       `json:"rate_limit_per_user,omitempty"`
	ParentID             string    `json:"parent_id,omitempty"`
	LastMessageID        string    `json:"last_message_id,omitempty"`
	IconEmoji            IconEmoji `json:"icon_emoji"`
	UserLimit            int       `json:"user_limit,omitempty"`
	RtcRegion            any       `json:"rtc_region,omitempty"`
	Bitrate              int       `json:"bitrate,omitempty"`
}

type IconEmoji struct {
	Name string `json:"name"`
	ID   any    `json:"id"`
}
