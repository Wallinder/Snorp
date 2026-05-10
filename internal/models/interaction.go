package models

type InteractionType uint8

type InteractionData struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	CommandType int                      `json:"type"`
	Resolved    *InteractionDataResolved `json:"resolved"`
	Options     []*InteractionDataOption `json:"options"`
	TargetID    string                   `json:"target_id"`
}

type InteractionDataResolved struct {
	Users       map[string]*User       `json:"users"`
	Members     map[string]*Member     `json:"members"`
	Roles       map[string]*Role       `json:"roles"`
	Channels    map[string]*Channel    `json:"channels"`
	Messages    map[string]*Message    `json:"messages"`
	Attachments map[string]*Attachment `json:"attachments"`
}

type InteractionDataOption struct {
	Name    string                       `json:"name"`
	Type    ApplicationCommandOptionType `json:"type"`
	Value   any                          `json:"value,omitempty"`
	Options []*InteractionDataOption     `json:"options,omitempty"`
	Focused bool                         `json:"focused,omitempty"`
}

type Interaction struct {
	ID                           string                                `json:"id"`
	AppID                        string                                `json:"application_id"`
	Type                         InteractionType                       `json:"type"`
	Data                         InteractionData                       `json:"data"`
	GuildID                      string                                `json:"guild_id"`
	ChannelID                    string                                `json:"channel_id"`
	Message                      *Message                              `json:"message"`
	AppPermissions               int64                                 `json:"app_permissions,string"`
	Member                       *Member                               `json:"member"`
	User                         *User                                 `json:"user"`
	Locale                       any                                   `json:"locale"`
	GuildLocale                  *any                                  `json:"guild_locale"`
	Context                      InteractionContextType                `json:"context"`
	AuthorizingIntegrationOwners map[ApplicationIntegrationType]string `json:"authorizing_integration_owners"`
	Token                        string                                `json:"token"`
	Version                      int                                   `json:"version"`
	Entitlements                 []*Entitlement                        `json:"entitlements"`
}
