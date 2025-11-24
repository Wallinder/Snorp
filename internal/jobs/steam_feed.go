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

const MAX_RETRIES = 3

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

		attempts := 0
		for {
			if attempts >= MAX_RETRIES {
				return fmt.Errorf("exceeded retry limit, when sending message to channel: %v", channelID)
			}
			if _, err := api.CreateMessage(session, channelID, message); err != nil {
				log.Printf("Error creating message: %v\n", err)
				attempts++
			}
			break
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

	newChannel := &api.GuildChannels{
		Name:  "steam-news",
		Type:  api.GUILD_TEXT,
		Topic: "snorp:steamnews",
	}

	newsChannelID, err := api.FindOrCreateChannel(session, newChannel, guildID)
	if err != nil {
		log.Println(err)
		return
	}

	ticker := time.NewTicker(30 * time.Minute)
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

	newChannel := &api.GuildChannels{
		Name:  "steam-sales",
		Type:  api.GUILD_TEXT,
		Topic: "snorp:steamsales",
	}

	salesChannelID, err := api.FindOrCreateChannel(session, newChannel, guildID)
	if err != nil {
		log.Println(err)
		return
	}

	ticker := time.NewTicker(30 * time.Minute)
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
