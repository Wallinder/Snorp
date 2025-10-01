package command

type ApplicationCommand struct {
	ID                       string                     `json:"id"`
	Type                     string                     `json:"type"`
	ApplicationID            string                     `json:"application_id,omitzero"`
	GuildID                  string                     `json:"guild_id,omitzero"`
	Name                     string                     `json:"name,omitzero"`
	NameLocalizations        []string                   `json:"name_localizations,omitzero"`
	Description              string                     `json:"description,omitzero"`
	DescriptionLocalizations []string                   `json:"description_localizations,omitzero"`
	Options                  []ApplicationCommandOption `json:"options,omitzero"`
	DefaultMemberPermissions string                     `json:"default_member_permissions,omitzero"`
	DmPermissions            bool                       `json:"dm_permission,omitzero"`
	DefaultPermissions       bool                       `json:"default_permission,omitzero"`
	Nsfw                     bool                       `json:"nsfw,omitzero"`
	IntegrationTypes         []int                      `json:"integration_types,omitzero"`
	Contexts                 []int                      `json:"contexts,omitzero"`
	Version                  string                     `json:"version,omitzero"`
	Handler                  []int                      `json:"handler,omitzero"`
}

type ApplicationCommandOption struct {
	Type                     int      `json:"type"`
	Name                     string   `json:"name"`
	NameLocalizations        []string `json:"name_localizations"`
	Description              string   `json:"description"`
	DescriptionLocalizations []string `json:"description_localizations"`
	Required                 bool     `json:"required"`
	Choices                  []string `json:"choices"`
	Options                  []int    `json:"options"`
	ChannelTypes             []int    `json:"channel_types"`
	MinValue                 int      `json:"min_value"`
	MaxValue                 int      `json:"max_value"`
	MinLength                int      `json:"min_length"`
	MaxLength                int      `json:"max_length"`
	Autocomplete             bool     `json:"autocomplete"`
}
