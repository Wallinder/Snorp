package models

import "time"

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
