package models

type InteractionType uint8

type InteractionData interface {
	Type() InteractionType
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
