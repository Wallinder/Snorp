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

func FindOrCreateChannel(session *state.SessionState, guildID string, topic, name string) (string, error) {
	channels, err := api.GetGuildChannels(session, guildID)
	if err != nil {
		return "", err
	}

	for _, channel := range channels {
		if channel.Topic == topic {
			return channel.ID, nil
		}
	}

	newChannel := &api.GuildChannels{
		Name:  name,
		Type:  api.GUILD_TEXT,
		Topic: topic,
	}
	created, err := api.CreateGuildChannel(session, guildID, newChannel)
	if err != nil {
		return "", fmt.Errorf("failed to create channel %s: %w", name, err)
	}
	return created.ID, nil
}

func ProcessFeedItems(session *state.SessionState, channelID string, items []steam.Item, lastRun time.Time) error {
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
			return err
		}

		time.Sleep(3 * time.Second)
	}
	return nil
}

func SteamNewsFeed(ctx context.Context, session *state.SessionState, guildID string) {
	session.Jobs.SteamNews[guildID] = true
	defer func() {
		session.Jobs.SteamNews[guildID] = false
	}()

	newsChannelID, err := FindOrCreateChannel(session, guildID, "snorp:steamnews", "steam-news")
	if err != nil {
		log.Println(err)
		return
	}

	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	var lastRun time.Time

	for {
		select {

		case <-ctx.Done():
			return

		case <-ticker.C:
			news, err := steam.GetNewsData()
			if err != nil {
				log.Printf("Error fetching news data: %v\n", err)
			} else {
				err := ProcessFeedItems(session, newsChannelID, news.Channel.Item, lastRun)
				if err != nil {
					return
				}
			}
			lastRun = time.Now()
		}
	}
}

func SteamSalesFeed(ctx context.Context, session *state.SessionState, guildID string) {
	session.Jobs.SteamSales[guildID] = true
	defer func() {
		session.Jobs.SteamSales[guildID] = false
	}()

	salesChannelID, err := FindOrCreateChannel(session, guildID, "snorp:steamsales", "steam-sales")
	if err != nil {
		log.Println(err)
		return
	}

	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	var lastRun time.Time

	for {
		select {

		case <-ctx.Done():
			return

		case <-ticker.C:
			sales, err := steam.GetSalesData()
			if err != nil {
			} else {
				err := ProcessFeedItems(session, salesChannelID, sales.Channel.Item, lastRun)
				if err != nil {
					return
				}
			}
			lastRun = time.Now()
		}
	}
}
