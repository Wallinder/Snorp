package state

type ReadyData struct {
	V                    int         `json:"v"`
	UserSettings         any         `json:"user_settings"`
	User                 User        `json:"user"`
	SessionType          string      `json:"session_type"`
	SessionID            string      `json:"session_id"`
	ResumeGatewayURL     string      `json:"resume_gateway_url"`
	Relationships        any         `json:"relationships"`
	PrivateChannels      any         `json:"private_channels"`
	Presences            any         `json:"presences"`
	Guilds               any         `json:"guilds"`
	GuildJoinRequests    any         `json:"guild_join_requests"`
	GeoOrderedRtcRegions []string    `json:"geo_ordered_rtc_regions"`
	GameRelationships    any         `json:"game_relationships"`
	Auth                 any         `json:"auth"`
	Application          Application `json:"application"`
}

type User struct {
	Verified      bool   `json:"verified"`
	Username      string `json:"username"`
	PrimaryGuild  any    `json:"primary_guild"`
	MfaEnabled    bool   `json:"mfa_enabled"`
	ID            string `json:"id"`
	GlobalName    any    `json:"global_name"`
	Flags         int    `json:"flags"`
	Email         any    `json:"email"`
	Discriminator string `json:"discriminator"`
	Clan          any    `json:"clan"`
	Bot           bool   `json:"bot"`
	Avatar        any    `json:"avatar"`
}

type Application struct {
	ID    string `json:"id"`
	Flags int    `json:"flags"`
}
