package api

type ApplicationCommand struct {
	ID                       string                     `json:"id,omitzero"`
	Type                     int                        `json:"type,omitzero"`
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
	IntegrationTypes         [2]int                     `json:"integration_types,omitzero"`
	Contexts                 [3]int                     `json:"contexts,omitzero"`
	Version                  string                     `json:"version,omitzero"`
	Handler                  int                        `json:"handler,omitzero"`
}

type ApplicationCommandOption struct {
	Type                     int      `json:"type,omitzero"`
	Name                     string   `json:"name,omitzero"`
	NameLocalizations        []string `json:"name_localizations,omitzero"`
	Description              string   `json:"description,omitzero"`
	DescriptionLocalizations []string `json:"description_localizations,omitzero"`
	Required                 bool     `json:"required,omitzero"`
	Choices                  []string `json:"choices,omitzero"`
	Options                  []int    `json:"options,omitzero"`
	ChannelTypes             []int    `json:"channel_types,omitzero"`
	MinValue                 int      `json:"min_value,omitzero"`
	MaxValue                 int      `json:"max_value,omitzero"`
	MinLength                int      `json:"min_length,omitzero"`
	MaxLength                int      `json:"max_length,omitzero"`
	Autocomplete             bool     `json:"autocomplete,omitzero"`
}
