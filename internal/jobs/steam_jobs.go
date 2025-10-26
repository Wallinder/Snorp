package jobs

import (
	"context"
	"fmt"
	"log"
	"snorp/internal/api"
	"snorp/internal/state"
	"snorp/pkg/steam"
	"time"
)

func SteamSales(ctx context.Context, session *state.SessionState, guild api.Guild) {
	var salesChannelID, newsChannelID string
	var salesChannel, newsChannel bool

	for _, channel := range guild.Channels {
		if channel.Topic == "snorp:steamsales" {
			salesChannel = true
			salesChannelID = channel.ID
		}
		if channel.Topic == "snorp:steamnews" {
			newsChannel = true
			newsChannelID = channel.ID
		}
	}

	if !salesChannel {
		steamSalesChannel := &api.GuildChannels{
			Name:  "steam-sales",
			Type:  api.GUILD_TEXT,
			Topic: "snorp:steamsales",
		}
		salesChannel, err := api.CreateGuildChannel(session, guild.ID, steamSalesChannel)
		if err != nil {
			log.Println(err)
		}
		salesChannelID = salesChannel.ID
	}

	if !newsChannel {
		steamNewsChannel := &api.GuildChannels{
			Name:  "steam-news",
			Type:  api.GUILD_TEXT,
			Topic: "snorp:steamnews",
		}
		newsChannel, err := api.CreateGuildChannel(session, guild.ID, steamNewsChannel)
		if err != nil {
			log.Println(err)
		}
		newsChannelID = newsChannel.ID
	}

	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	var lastRun time.Time
	for {
		select {

		case <-ctx.Done():
			return

		case <-ticker.C:
			sales, err := steam.GetSalesData()
			if err != nil {
				log.Println(err)
			}

			for _, item := range sales.Channel.Item {
				pubDate, err := time.Parse(time.RFC1123, item.PubDate)
				if err != nil {
					log.Println("Error parsing time:", err)
				}
				if pubDate.After(lastRun) {
					newsMessage := api.Message{
						Content: fmt.Sprintf("%s\n%s", item.Title, item.Link),
					}
					_, err := api.CreateMessage(session, salesChannelID, newsMessage)
					if err != nil {
						log.Println(err)
						return
					}
					time.Sleep(3 * time.Second)
				}
			}

			news, err := steam.GetNewsData()
			if err != nil {
				log.Println(err)
			}

			for _, item := range news.Channel.Item {
				pubDate, err := time.Parse(time.RFC1123, item.PubDate)
				if err != nil {
					log.Println("Error parsing time:", err)
				}
				if pubDate.After(lastRun) {
					newsMessage := api.Message{
						Content: fmt.Sprintf("%s\n%s\n", item.Title, item.Link),
					}
					_, err := api.CreateMessage(session, newsChannelID, newsMessage)
					if err != nil {
						log.Println(err)
						return
					}
					time.Sleep(3 * time.Second)
				}
			}
			lastRun = time.Now()
		}
	}
}
