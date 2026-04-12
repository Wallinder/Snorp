package models

type Role struct {
	Version      int64  `json:"version"`
	UnicodeEmoji any    `json:"unicode_emoji"`
	Tags         string `json:"tags"`
	Position     int    `json:"position"`
	Permissions  string `json:"permissions"`
	Name         string `json:"name"`
	Mentionable  bool   `json:"mentionable"`
	Managed      bool   `json:"managed"`
	ID           string `json:"id"`
	Icon         any    `json:"icon"`
	Hoist        bool   `json:"hoist"`
	Flags        int    `json:"flags"`
	Colors       Colors `json:"colors"`
	Color        int    `json:"color"`
}

type Colors struct {
	TertiaryColor  any `json:"tertiary_color"`
	SecondaryColor any `json:"secondary_color"`
	PrimaryColor   int `json:"primary_color"`
}
