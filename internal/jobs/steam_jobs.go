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

func FindOrCreateChannel(session *state.SessionState, guild api.Guild, topic, name string) (string, error) {
	for _, channel := range guild.Channels {
		if channel.Topic == topic {
			return channel.ID, nil
		}
	}

	newChannel := &api.GuildChannels{
		Name:  name,
		Type:  api.GUILD_TEXT,
		Topic: topic,
	}
	created, err := api.CreateGuildChannel(session, guild.ID, newChannel)
	if err != nil {
		return "", fmt.Errorf("failed to create channel %s: %w", name, err)
	}
	return created.ID, nil
}

func ProcessFeedItems(session *state.SessionState, channelID string, items []steam.Item, lastRun time.Time) {
	for _, item := range items {
		pubDate, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			log.Printf("Error parsing time: %v", err)
			continue
		}
		if !pubDate.After(lastRun) {
			continue
		}

		message := api.Message{
			Content: fmt.Sprintf("%s\n%s", item.Title, item.Link),
		}
		if _, err := api.CreateMessage(session, channelID, message); err != nil {
			log.Printf("Error creating message: %v", err)
			continue
		}

		time.Sleep(3 * time.Second)
	}
}

func SteamFeed(ctx context.Context, session *state.SessionState, guild api.Guild) {
	salesChannelID, err := FindOrCreateChannel(session, guild, "snorp:steamsales", "steam-sales")
	if err != nil {
		log.Println(err)
		return
	}
	newsChannelID, err := FindOrCreateChannel(session, guild, "snorp:steamnews", "steam-news")
	if err != nil {
		log.Println(err)
		return
	}

	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	var lastRun = session.StartTime
	for {
		select {

		case <-ctx.Done():
			return

		case <-ticker.C:
			sales, err := steam.GetSalesData()
			if err != nil {
				log.Printf("Error fetching sales data: %v", err)
			} else {
				ProcessFeedItems(session, salesChannelID, sales.Channel.Item, lastRun)
			}

			news, err := steam.GetNewsData()
			if err != nil {
				log.Printf("Error fetching news data: %v", err)
			} else {
				ProcessFeedItems(session, newsChannelID, news.Channel.Item, lastRun)
			}

			lastRun = time.Now()
		}
	}
}
