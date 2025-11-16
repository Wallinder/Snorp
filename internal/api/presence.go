package api

type Presence struct {
	User                 User         `json:"user"`
	GuildID              string       `json:"guild_id"`
	Status               string       `json:"status"`
	ProcessedAtTimestamp int64        `json:"processed_at_timestamp"`
	Activities           []Activity   `json:"activities,omitzero"`
	ClientStatus         ClientStatus `json:"client_status"`
}

type ClientStatus struct {
	Desktop string `json:"desktop,omitzero"`
	Mobile  string `json:"mobile,omitzero"`
	Web     string `json:"web,omitzero"`
}

type Activity struct {
	Name              string             `json:"name"`
	Type              int                `json:"type"`
	Url               string             `json:"url,omitzero"`
	CreatedAt         int                `json:"created_at"`
	Timestamps        ActivityTimestamps `json:"timestamps,omitzero"`
	ApplicationID     string             `json:"application_id,omitzero"`
	StatusDisplayType int                `json:"status_display_type,omitzero"`
	Details           string             `json:"details,omitzero"`
	DetailsUrl        string             `json:"details_url,omitzero"`
	State             string             `json:"state,omitzero"`
	StateUrl          string             `json:"state_url,omitzero"`
	Emoji             ActivityEmoji      `json:"emoji,omitzero"`
	Party             ActivityParty      `json:"party,omitzero"`
	Assets            ActivityAssets     `json:"assets,omitzero"`
	Secrets           ActivitySecrets    `json:"secrets,omitzero"`
	Instance          bool               `json:"instance,omitzero"`
	Flags             int                `json:"flags,omitzero"`
	Buttons           []ActivityButtons  `json:"buttons,omitzero"`
}

type ActivityTimestamps struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type ActivityEmoji struct {
	Name     string `json:"name"`
	ID       string `json:"id,omitzero"`
	Animated bool   `json:"details,omitzero"`
}

type ActivityParty struct {
	ID   int `json:"id"`
	Size int `json:"size"`
}

type ActivityAssets struct {
	LargeImage       string `json:"large_image,omitzero"`
	LargeText        string `json:"large_text,omitzero"`
	LargeUrl         string `json:"large_url,omitzero"`
	SmallImage       string `json:"small_image,omitzero"`
	SmallText        string `json:"small_text,omitzero"`
	SmallUrl         string `json:"small_url,omitzero"`
	InviteCoverImage string `json:"invite_cover_image,omitzero"`
}

type ActivitySecrets struct {
	Join     string `json:"join"`
	Spectate string `json:"spectate"`
	Match    string `json:"match"`
}

type ActivityButtons struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}
