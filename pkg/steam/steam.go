package steam

import (
	"encoding/xml"
	"io"
	"net/http"
)

const (
	STEAM_NEWS  = "https://store.steampowered.com/feeds/news/collection/steam"
	STEAM_SALES = "https://store.steampowered.com/feeds/news/collection/sales"
)

func GetSalesData() (*SteamRSS, error) {
	var rss *SteamRSS

	response, err := http.Get(STEAM_SALES)
	if err != nil {
		return rss, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return rss, err
	}

	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return rss, err
	}
	return rss, nil
}

func GetNewsData() (*SteamRSS, error) {
	var rss *SteamRSS

	response, err := http.Get(STEAM_NEWS)
	if err != nil {
		return rss, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return rss, err
	}

	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return rss, err
	}
	return rss, nil
}
