package models

import (
	"time"
)

type Message struct {
	Content      string        `json:"content,omitempty"`
	Nonce        string        `json:"nonce,omitempty"`
	TTS          bool          `json:"tts,omitempty"`
	Embeds       []*Embed      `json:"embeds"`
	StickerIDs   []*string     `json:"sticker_ids"`
	Attachments  []*Attachment `json:"attachments"`
	Flags        int           `json:"flags,omitempty"`
	EnforceNonce bool          `json:"enforce_nonce,omitempty"`
	PayloadJSON  string        `json:"payload_json,omitempty"`
}

type Attachment struct {
	ID          string `json:"id,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Description string `json:"description,omitempty"`
}

type Embed struct {
	Title       string         `json:"title,omitempty"`
	Type        string         `json:"type,omitempty"`
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitempty"`
	Timestamp   time.Time      `json:"timestamp"`
	Color       int            `json:"color,omitempty"`
	Footer      EmbedFooter    `json:"footer"`
	Image       EmbedImage     `json:"image"`
	Thumbnail   EmbedThumbnail `json:"thumbnail"`
	Video       EmbedVideo     `json:"video"`
	Provider    EmbedProvider  `json:"provider"`
	Author      EmbedAuthor    `json:"author"`
	Fields      []*EmbedField  `json:"fields"`
}

type EmbedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// message reaction add
type MessageReactionAdd struct {
	UserID          string `json:"user_id"`
	Type            int    `json:"type"`
	MessageID       string `json:"message_id"`
	MessageAuthorID string `json:"message_author_id"`
	Member          Member `json:"member"`
	Emoji           Emoji  `json:"emoji"`
	ChannelID       string `json:"channel_id"`
	Burst           bool   `json:"burst"`
	GuildID         string `json:"guild_id"`
}

type Emoji struct {
	Name string `json:"name"`
	ID   any    `json:"id"`
}
