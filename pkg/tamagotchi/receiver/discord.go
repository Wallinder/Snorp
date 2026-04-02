package receiver

type DiscordMessage struct {
	ChannelID string
	Embeds    []Embed `json:"embeds"`
}

type Embed struct {
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"decription"`
	Image       Image  `json:"image"`
}

type Image struct {
	Url string `json:"url"`
}

func (d *DiscordMessage) Notify() error {
	return nil
}
