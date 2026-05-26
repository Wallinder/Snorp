package discord

import "time"

type ComponentType uint

const (
	ActionsRowComponent            ComponentType = 1
	ButtonComponent                ComponentType = 2
	SelectMenuComponent            ComponentType = 3
	TextInputComponent             ComponentType = 4
	UserSelectMenuComponent        ComponentType = 5
	RoleSelectMenuComponent        ComponentType = 6
	MentionableSelectMenuComponent ComponentType = 7
	ChannelSelectMenuComponent     ComponentType = 8
	SectionComponent               ComponentType = 9
	TextDisplayComponent           ComponentType = 10
	ThumbnailComponent             ComponentType = 11
	MediaGalleryComponent          ComponentType = 12
	FileComponentType              ComponentType = 13
	SeparatorComponent             ComponentType = 14
	ContainerComponent             ComponentType = 17
	LabelComponent                 ComponentType = 18
	FileUploadComponent            ComponentType = 19
)

type EntitlementType int

type Entitlement struct {
	ID             string          `json:"id"`
	SKUID          string          `json:"sku_id"`
	ApplicationID  string          `json:"application_id"`
	UserID         string          `json:"user_id,omitempty"`
	Type           EntitlementType `json:"type"`
	Deleted        bool            `json:"deleted"`
	StartsAt       *time.Time      `json:"starts_at,omitempty"`
	EndsAt         *time.Time      `json:"ends_at,omitempty"`
	GuildID        string          `json:"guild_id,omitempty"`
	Consumed       *bool           `json:"consumed,omitempty"`
	SubscriptionID string          `json:"subscription_id,omitempty"`
}

const (
	RoleMentions     = "roles"
	UserMentions     = "users"
	EveryoneMentions = "everyone"
)

type AllowedMentions struct {
	Parse []string `json:"parse,omitempty"`
	Users []string `json:"users,omitempty"`
}

type Poll struct {
	Question         PollMedia
	Answer           []PollAnswer
	Duration         int  `json:"duration,omitempty"`
	AllowMultiselect bool `json:"allow_multiselect,omitempty"`
	LayoutType       int  `json:"layout_type,omitempty"`
}

type PollAnswer struct {
	AnswerID  int       `json:"answer_id,omitempty"`
	PollMedia PollMedia `json:"poll_media"`
}

type PollMedia struct {
	Text  string `json:"text,omitempty"`
	Emoji *Emoji `json:"emoji,omitempty"`
}
