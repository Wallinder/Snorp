package discord

import "time"

type Member struct {
	User                       User      `json:"user"`
	Roles                      []string  `json:"roles"` //ID of the roles, not the roles itself
	PremiumSince               string    `json:"premium_since,omitempty"`
	Nick                       string    `json:"nick,omitempty"`
	Mute                       bool      `json:"mute"`
	JoinedAt                   time.Time `json:"joined_at"`
	Flags                      int       `json:"flags"`
	Deaf                       bool      `json:"deaf"`
	CommunicationDisabledUntil string    `json:"communication_disabled_until"`
	Banner                     string    `json:"banner,omitempty"`
	Avatar                     string    `json:"avatar,omitempty"`
	Pending                    bool      `json:"pending,omitempty"`
	Permissions                string    `json:"permissions,omitempty"`
	GuildID                    string    `json:"guild_id,omitempty"` //Extra field added when called by "Guild Member Add"
}

type User struct {
	Username             string               `json:"username"`
	PublicFlags          int                  `json:"public_flags"`
	PrimaryGuild         UserPrimaryGuild     `json:"primary_guild"`
	ID                   string               `json:"id"`
	GlobalName           string               `json:"global_name"`
	DisplayName          string               `json:"display_name"`
	Discriminator        string               `json:"discriminator"`
	Collectibles         any                  `json:"collectibles"`
	Bot                  bool                 `json:"bot"`
	AvatarDecorationData AvatarDecorationData `json:"avatar_decoration_data"`
	Avatar               string               `json:"avatar"`
}

type AvatarDecorationData struct {
	Asset string `json:"asset"`
	SkuID string `json:"sku_id"`
}

type UserPrimaryGuild struct {
	IdentityGuildID string `json:"identity_guild_id"`
	IdentityEnabled bool   `json:"identity_enabled"`
	Tag             string `json:"tag"`
	Badge           string `json:"badge"`
}

type Role struct {
	Version      int64    `json:"version"`
	UnicodeEmoji string   `json:"unicode_emoji"`
	Tags         RoleTags `json:"tags"`
	Position     int      `json:"position"`
	Permissions  string   `json:"permissions"`
	Name         string   `json:"name"`
	Mentionable  bool     `json:"mentionable"`
	Managed      bool     `json:"managed"`
	ID           string   `json:"id"`
	Icon         string   `json:"icon"`
	Hoist        bool     `json:"hoist"`
	Flags        int      `json:"flags"`
	Colors       Colors   `json:"colors"`
	Color        int      `json:"color"`
}

type RoleTags struct {
	BotID                 string `json:"bot_id"`
	IntegrationID         string `json:"integration_id"`
	PremiumSubscriber     bool   `json:"premium_subscriber"`
	SubscriptionListingID string `json:"subscription_listing_id"`
	AvailableForPurchase  bool   `json:"available_for_purchase"`
	GuildConnections      bool   `json:"guild_connections"`
}

type Colors struct {
	TertiaryColor  any `json:"tertiary_color"`
	SecondaryColor any `json:"secondary_color"`
	PrimaryColor   int `json:"primary_color"`
}
