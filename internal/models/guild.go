package models

import "time"

type Guild struct {
	MfaLevel                               int         `json:"mfa_level"`
	PremiumSubscriptionCount               int         `json:"premium_subscription_count"`
	Channels                               []*Channel  `json:"channels"`
	SafetyAlertsChannelID                  any         `json:"safety_alerts_channel_id"`
	Profile                                any         `json:"profile"`
	Experiments                            Experiments `json:"experiments"`
	HomeHeader                             any         `json:"home_header"`
	DefaultMessageNotifications            int         `json:"default_message_notifications"`
	VerificationLevel                      int         `json:"verification_level"`
	NsfwLevel                              int         `json:"nsfw_level"`
	PremiumTier                            int         `json:"premium_tier"`
	VoiceStates                            []*any      `json:"voice_states"`
	SoundboardSounds                       []*any      `json:"soundboard_sounds"`
	Region                                 string      `json:"region"`
	Splash                                 any         `json:"splash"`
	ID                                     string      `json:"id"`
	MaxStageVideoChannelUsers              int         `json:"max_stage_video_channel_users"`
	StageInstances                         []*any      `json:"stage_instances"`
	Nsfw                                   bool        `json:"nsfw"`
	InventorySettings                      any         `json:"inventory_settings"`
	Name                                   string      `json:"name"`
	Icon                                   any         `json:"icon"`
	IncidentsData                          any         `json:"incidents_data"`
	Members                                []*Member   `json:"members"`
	AfkTimeout                             int         `json:"afk_timeout"`
	PreferredLocale                        string      `json:"preferred_locale"`
	Threads                                []*any      `json:"threads"`
	MaxMembers                             int         `json:"max_members"`
	GuildScheduledEvents                   []*any      `json:"guild_scheduled_events"`
	DiscoverySplash                        any         `json:"discovery_splash"`
	HubType                                any         `json:"hub_type"`
	Emojis                                 []*Emoji    `json:"emojis"`
	MemberCount                            int         `json:"member_count"`
	Description                            any         `json:"description"`
	Roles                                  any         `json:"roles,omitempty"`
	ModeratorReporting                     any         `json:"moderator_reporting"`
	LatestOnboardingQuestionID             any         `json:"latest_onboarding_question_id"`
	SystemChannelID                        string      `json:"system_channel_id"`
	Banner                                 any         `json:"banner"`
	PublicUpdatesChannelID                 any         `json:"public_updates_channel_id"`
	ApplicationCommandCounts               any         `json:"application_command_counts"`
	Large                                  bool        `json:"large"`
	PremiumProgressBarEnabledUserUpdatedAt any         `json:"premium_progress_bar_enabled_user_updated_at"`
	PremiumProgressBarEnabled              bool        `json:"premium_progress_bar_enabled"`
	PremiumFeatures                        any         `json:"premium_features"`
	Features                               []string    `json:"features"`
	VanityURLCode                          any         `json:"vanity_url_code"`
	SystemChannelFlags                     int         `json:"system_channel_flags"`
	Presences                              []*Presence `json:"presences"`
	Unavailable                            bool        `json:"unavailable"`
	AfkChannelID                           any         `json:"afk_channel_id"`
	Stickers                               []*any      `json:"stickers"`
	EmbeddedActivities                     []*any      `json:"embedded_activities"`
	MaxVideoChannelUsers                   int         `json:"max_video_channel_users"`
	ApplicationID                          any         `json:"application_id"`
	OwnerID                                string      `json:"owner_id"`
	OwnerConfiguredContentLevel            int         `json:"owner_configured_content_level"`
	RulesChannelID                         any         `json:"rules_channel_id"`
	Version                                int64       `json:"version"`
	ActivityInstances                      []*any      `json:"activity_instances"`
	Lazy                                   bool        `json:"lazy"`
	JoinedAt                               time.Time   `json:"joined_at"`
	ExplicitContentFilter                  int         `json:"explicit_content_filter"`
	GameApplicationIds                     any         `json:"game_application_ids"`
}

type Experiments struct {
	EvaluationID any   `json:"evaluation_id"`
	Assignments  []any `json:"assignments"`
}
