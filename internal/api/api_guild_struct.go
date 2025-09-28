package api

import "time"

type Guild struct {
	Region                      string          `json:"region"`
	HomeHeader                  any             `json:"home_header"`
	PremiumSubscriptionCount    int             `json:"premium_subscription_count"`
	DiscoverySplash             any             `json:"discovery_splash"`
	ModeratorReporting          any             `json:"moderator_reporting"`
	SoundboardSounds            []any           `json:"soundboard_sounds"`
	Unavailable                 bool            `json:"unavailable"`
	DefaultMessageNotifications int             `json:"default_message_notifications"`
	Profile                     any             `json:"profile"`
	Splash                      any             `json:"splash"`
	Banner                      any             `json:"banner"`
	SafetyAlertsChannelID       any             `json:"safety_alerts_channel_id"`
	Presences                   []any           `json:"presences"`
	MaxVideoChannelUsers        int             `json:"max_video_channel_users"`
	PreferredLocale             string          `json:"preferred_locale"`
	Large                       bool            `json:"large"`
	JoinedAt                    time.Time       `json:"joined_at"`
	PremiumFeatures             any             `json:"premium_features"`
	SystemChannelID             string          `json:"system_channel_id"`
	HubType                     any             `json:"hub_type"`
	ID                          string          `json:"id"`
	VanityURLCode               any             `json:"vanity_url_code"`
	PremiumProgressBarEnabled   bool            `json:"premium_progress_bar_enabled"`
	VerificationLevel           int             `json:"verification_level"`
	Stickers                    []any           `json:"stickers"`
	SystemChannelFlags          int             `json:"system_channel_flags"`
	Name                        string          `json:"name"`
	EmbeddedActivities          []any           `json:"embedded_activities"`
	MfaLevel                    int             `json:"mfa_level"`
	Members                     []GuildMembers  `json:"members"`
	PublicUpdatesChannelID      any             `json:"public_updates_channel_id"`
	Threads                     []any           `json:"threads"`
	GuildScheduledEvents        []any           `json:"guild_scheduled_events"`
	ApplicationID               any             `json:"application_id"`
	AfkChannelID                any             `json:"afk_channel_id"`
	ActivityInstances           []any           `json:"activity_instances"`
	Description                 any             `json:"description"`
	Channels                    []GuildChannels `json:"channels"`
	Features                    []string        `json:"features"`
	InventorySettings           any             `json:"inventory_settings"`
	OwnerID                     string          `json:"owner_id"`
	AfkTimeout                  int             `json:"afk_timeout"`
	MaxStageVideoChannelUsers   int             `json:"max_stage_video_channel_users"`
	OwnerConfiguredContentLevel int             `json:"owner_configured_content_level"`
	Nsfw                        bool            `json:"nsfw"`
	NsfwLevel                   int             `json:"nsfw_level"`
	VoiceStates                 []any           `json:"voice_states"`
	MaxMembers                  int             `json:"max_members"`
	Emojis                      []any           `json:"emojis"`
	StageInstances              []any           `json:"stage_instances"`
	Lazy                        bool            `json:"lazy"`
	LatestOnboardingQuestionID  any             `json:"latest_onboarding_question_id"`
	IncidentsData               any             `json:"incidents_data"`
	MemberCount                 int             `json:"member_count"`
	Icon                        any             `json:"icon"`
	RulesChannelID              any             `json:"rules_channel_id"`
	Roles                       []GuildRoles    `json:"roles"`
	Version                     int64           `json:"version"`
	PremiumTier                 int             `json:"premium_tier"`
	ExplicitContentFilter       int             `json:"explicit_content_filter"`
}

type GuildChannels struct {
	Version          int64                      `json:"version,omitzero"`
	Type             int                        `json:"type,omitzero"`
	Position         int                        `json:"position,omitzero"`
	Permissions      []GuildChannelsPermissions `json:"permission_overwrites,omitzero"`
	Name             string                     `json:"name,omitzero"`
	ID               string                     `json:"id,omitzero"`
	Flags            int                        `json:"flags,omitzero"`
	Topic            any                        `json:"topic,omitzero"`
	RateLimitPerUser int                        `json:"rate_limit_per_user,omitzero"`
	ParentID         string                     `json:"parent_id,omitzero"`
	LastMessageID    string                     `json:"last_message_id,omitzero"`
	UserLimit        int                        `json:"user_limit,omitzero"`
	RtcRegion        any                        `json:"rtc_region,omitzero"`
	Bitrate          int                        `json:"bitrate,omitzero"`
}

type GuildChannelsPermissions struct {
	ID    string `json:"id"`
	Type  int    `json:"type"`
	Allow string `json:"allow"`
	Deny  string `json:"deny"`
}

type GuildMembers struct {
	User                       GuildUser `json:"user"`
	Roles                      []string  `json:"roles"`
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

type GuildUser struct {
	ID                   string `json:"id"`
	Username             string `json:"username"`
	Discriminator        string `json:"discriminator"`
	GlobalName           any    `json:"global_name"`
	Avatar               string `json:"avatar"`
	Bot                  bool   `json:"bot"`
	System               bool   `json:"system"`
	MfaEnabled           bool   `json:"mfa_enabled"`
	PublicFlags          int    `json:"public_flags"`
	PrimaryGuild         any    `json:"primary_guild"`
	DisplayNameStyles    any    `json:"display_name_styles"`
	DisplayName          any    `json:"display_name"`
	Collectibles         any    `json:"collectibles"`
	AvatarDecorationData any    `json:"avatar_decoration_data"`
}

type GuildRoles struct {
	Version      int64  `json:"version"`
	UnicodeEmoji any    `json:"unicode_emoji"`
	Position     int    `json:"position"`
	Permissions  string `json:"permissions"`
	Name         string `json:"name"`
	Mentionable  bool   `json:"mentionable"`
	Managed      bool   `json:"managed"`
	ID           string `json:"id"`
	Icon         any    `json:"icon"`
	Hoist        bool   `json:"hoist"`
	Flags        int    `json:"flags"`
	Color        int    `json:"color"`
}
