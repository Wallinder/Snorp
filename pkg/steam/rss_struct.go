package steam

import (
	"encoding/xml"
)

type SteamRSS struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Link struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Description struct {
	Text string `xml:",chardata"`
	P    string `xml:"p"`
}

type Guid struct {
	Text        string `xml:",chardata"`
	IsPermaLink string `xml:"isPermaLink,attr"`
}

type Enclosure struct {
	Text   string `xml:",chardata"`
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Item struct {
	Text        string      `xml:",chardata"`
	Title       string      `xml:"title"`
	Description Description `xml:"description"`
	Link        string      `xml:"link"`
	PubDate     string      `xml:"pubDate"`
	Guid        Guid        `xml:"guid"`
	Enclosure   Enclosure   `xml:"enclosure"`
}

type Channel struct {
	Text        string `xml:",chardata"`
	Link        Link   `xml:"link"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	Generator   string `xml:"generator"`
	Item        []Item `xml:"item"`
}
