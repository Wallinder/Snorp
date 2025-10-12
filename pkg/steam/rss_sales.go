package steam

import (
	"encoding/xml"
)

const STEAM_SALES = "https://store.steampowered.com/feeds/news/collection/sales/"

type SteamSalesRSS struct {
	XMLName xml.Name     `xml:"rss"`
	Text    string       `xml:",chardata"`
	Version string       `xml:"version,attr"`
	Channel SalesChannel `xml:"channel"`
}

type SalesLink struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type SalesDescription struct {
	Text string `xml:",chardata"`
	P    string `xml:"p"`
}

type SalesGuid struct {
	Text        string `xml:",chardata"`
	IsPermaLink string `xml:"isPermaLink,attr"`
}

type SalesEnclosure struct {
	Text   string `xml:",chardata"`
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type SalesItem struct {
	Text        string           `xml:",chardata"`
	Title       string           `xml:"title"`
	Description SalesDescription `xml:"description"`
	Link        string           `xml:"link"`
	PubDate     string           `xml:"pubDate"`
	Guid        SalesGuid        `xml:"guid"`
	Enclosure   SalesEnclosure   `xml:"enclosure"`
}

type SalesChannel struct {
	Text        string      `xml:",chardata"`
	Link        SalesLink   `xml:"link"`
	Title       string      `xml:"title"`
	Description string      `xml:"description"`
	Language    string      `xml:"language"`
	Generator   string      `xml:"generator"`
	Item        []SalesItem `xml:"item"`
}
