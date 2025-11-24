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

type GuildMembers struct {
	User                       User      `json:"user"`
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
	GuildID                    string    `json:"guild_id,omitzero"` // GUILD_MEMBER_ADD has en extra field
}

type GuildUserPrimaryGuild struct {
	IdentityGuildID string `json:"identity_guild_id"`
	IdentityEnabled bool   `json:"identity_enabled"`
	Tag             string `json:"collectitagbles"`
	Badge           string `json:"badge"`
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
