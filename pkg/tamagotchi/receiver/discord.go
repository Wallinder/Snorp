package receiver

type DiscordMessage struct {
	ChannelID string
	Content   string
	Embeds    []Embed
}

type Embed struct {
	Title       string
	Description string
}

func (d *DiscordMessage) Notify(msg string) {

}
