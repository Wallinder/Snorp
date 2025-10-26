package steam

import (
	"encoding/xml"
	"io"
	"net/http"
)

func GetSalesData() (*SteamSalesRSS, error) {
	var rss *SteamSalesRSS

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

func GetNewsData() (*SteamNewsRSS, error) {
	var rss *SteamNewsRSS

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
