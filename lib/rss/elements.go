package rss

import (
	"encoding/xml"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item,omitempty"`
}

type Item struct {
	Title       string `xml:"title,omitempty"`
	Description string `xml:"description,omitempty"`
	Author      string `xml:"author,omitempty"`
	Category    string `xml:"category,omitempty"`
	Comments    string `xml:"comments,omitempty"`
	Enclosure   string `xml:"enclosure,omitempty"`
	Guid        string `xml:"guid,omitempty"`
	Link        string `xml:"link,omitempty"`
	PubDate     string `xml:"pubDate,omitempty"`
	Source      string `xml:"source,omitempty"`
}
