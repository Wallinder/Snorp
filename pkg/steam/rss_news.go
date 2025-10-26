package steam

import (
	"encoding/xml"
)

const STEAM_NEWS = "https://store.steampowered.com/feeds/news/collection/steam"

type SteamNewsRSS struct {
	XMLName xml.Name    `xml:"rss"`
	Text    string      `xml:",chardata"`
	Version string      `xml:"version,attr"`
	Channel NewsChannel `xml:"channel"`
}

type NewsLink struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type NewsDescription struct {
	Text string `xml:",chardata"`
	P    string `xml:"p"`
}

type NewsGuid struct {
	Text        string `xml:",chardata"`
	IsPermaLink string `xml:"isPermaLink,attr"`
}

type NewsEnclosure struct {
	Text   string `xml:",chardata"`
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type NewsItem struct {
	Text        string          `xml:",chardata"`
	Title       string          `xml:"title"`
	Description NewsDescription `xml:"description"`
	Link        string          `xml:"link"`
	PubDate     string          `xml:"pubDate"`
	Guid        NewsGuid        `xml:"guid"`
	Enclosure   NewsEnclosure   `xml:"enclosure"`
}

type NewsChannel struct {
	Text        string     `xml:",chardata"`
	Link        NewsLink   `xml:"link"`
	Title       string     `xml:"title"`
	Description string     `xml:"description"`
	Language    string     `xml:"language"`
	Generator   string     `xml:"generator"`
	Item        []NewsItem `xml:"item"`
}
