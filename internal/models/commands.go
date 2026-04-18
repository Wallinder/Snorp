package models

type ApplicationCommandType uint8

const (
	ChatApplicationCommand    ApplicationCommandType = 1
	UserApplicationCommand    ApplicationCommandType = 2
	MessageApplicationCommand ApplicationCommandType = 3
	EntryApplicationCommand   ApplicationCommandType = 4
)

type InteractionContextType uint

const (
	InteractionContextGuild          InteractionContextType = 0
	InteractionContextBotDM          InteractionContextType = 1
	InteractionContextPrivateChannel InteractionContextType = 2
)

type ApplicationIntegrationType uint

const (
	ApplicationIntegrationGuildInstall ApplicationIntegrationType = 0
	ApplicationIntegrationUserInstall  ApplicationIntegrationType = 1
)

type ApplicationCommand struct {
	ID                       string                        `json:"id,omitempty"`
	ApplicationID            string                        `json:"application_id,omitempty"`
	GuildID                  string                        `json:"guild_id,omitempty"`
	Version                  string                        `json:"version,omitempty"`
	Type                     ApplicationCommandType        `json:"type,omitempty"`
	Name                     string                        `json:"name"`
	NameLocalizations        any                           `json:"name_localizations,omitempty"`
	DefaultPermission        *bool                         `json:"default_permission,omitempty"`
	DefaultMemberPermissions *int64                        `json:"default_member_permissions,string,omitempty"`
	NSFW                     *bool                         `json:"nsfw,omitempty"`
	DMPermission             *bool                         `json:"dm_permission,omitempty"`
	Contexts                 *[]InteractionContextType     `json:"contexts,omitempty"`
	IntegrationTypes         *[]ApplicationIntegrationType `json:"integration_types,omitempty"`
	Description              string                        `json:"description,omitempty"`
	DescriptionLocalizations any                           `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption   `json:"options"`
}

type ApplicationCommandOptionType uint8

const (
	ApplicationCommandOptionSubCommand      ApplicationCommandOptionType = 1
	ApplicationCommandOptionSubCommandGroup ApplicationCommandOptionType = 2
	ApplicationCommandOptionString          ApplicationCommandOptionType = 3
	ApplicationCommandOptionInteger         ApplicationCommandOptionType = 4
	ApplicationCommandOptionBoolean         ApplicationCommandOptionType = 5
	ApplicationCommandOptionUser            ApplicationCommandOptionType = 6
	ApplicationCommandOptionChannel         ApplicationCommandOptionType = 7
	ApplicationCommandOptionRole            ApplicationCommandOptionType = 8
	ApplicationCommandOptionMentionable     ApplicationCommandOptionType = 9
	ApplicationCommandOptionNumber          ApplicationCommandOptionType = 10
	ApplicationCommandOptionAttachment      ApplicationCommandOptionType = 11
)

type ApplicationCommandOption struct {
	Type                     ApplicationCommandOptionType      `json:"type"`
	Name                     string                            `json:"name"`
	NameLocalizations        any                               `json:"name_localizations,omitempty"`
	Description              string                            `json:"description,omitempty"`
	DescriptionLocalizations any                               `json:"description_localizations,omitempty"`
	ChannelTypes             []ChannelType                     `json:"channel_types"`
	Required                 bool                              `json:"required"`
	Options                  []*ApplicationCommandOption       `json:"options"`
	Autocomplete             bool                              `json:"autocomplete"`
	Choices                  []*ApplicationCommandOptionChoice `json:"choices"`
	MinValue                 *float64                          `json:"min_value,omitempty"`
	MaxValue                 float64                           `json:"max_value,omitempty"`
	MinLength                *int                              `json:"min_length,omitempty"`
	MaxLength                int                               `json:"max_length,omitempty"`
}

type ApplicationCommandOptionChoice struct {
	Name              string `json:"name"`
	NameLocalizations any    `json:"name_localizations,omitempty"`
	Value             any    `json:"value"`
}
