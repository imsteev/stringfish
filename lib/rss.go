package lib

import (
	"encoding/xml"
	"fmt"
	"rss/clients"
)

type HackerNewsRss struct {
	Username string
}

func (h HackerNewsRss) GenerateXml() string {
	c := clients.HackerNewsClient{}
	u, _ := c.GetUser(h.Username)

	itemsXml := ""
	for _, submittedID := range u.Submitted {
		item, _ := c.GetItem(submittedID)
		itemXml, _ := xml.Marshal(makeRssItem(item))
		itemsXml += string(itemXml)
	}

	return fmt.Sprintf(`<rss version="2.0"><channel><title>%s</title><description>%s</description><link>%s</link>%s</channel></rss>`,
		u.Id, u.About, hnUserLink(u.Id), itemsXml)
}

func hnUserLink(id string) string {
	return fmt.Sprintf("https://news.ycombinator.com/user?id=%s", id)
}

func makeRssItem(i clients.Item) rssItem {
	var r rssItem
	switch i.Type {
	case "story":
		r = rssItem{
			Title:   i.Title,
			Author:  i.By,
			Link:    i.Url,
			PubDate: i.Time,
		}
	case "comment":
		r = rssItem{
			Description: i.Text,
			Author:      i.By,
			PubDate:     i.Time,
		}
	case "ask":
		r = rssItem{
			Title:       i.Title,
			Description: i.Text,
			Author:      i.By,
			PubDate:     i.Time,
		}
	case "job":
		r = rssItem{
			Title:       i.Title,
			Description: i.Text,
			Author:      i.By,
			Link:        i.Url,
			PubDate:     i.Time,
		}
	case "poll":
		r = rssItem{
			Title:       i.Title,
			Description: i.Text,
			Link:        i.Url,
			PubDate:     i.Time,
		}
	case "pollopt":
		r = rssItem{
			Description: i.Text,
			Author:      i.By,
			PubDate:     i.Time,
		}
	default:
		return rssItem{}
	}
	return r

}

type rssItem struct {
	Author      string `xml:"author"`
	Category    string `xml:"category"`
	Comments    string `xml:"comments"`
	Description string `xml:"description"`
	Enclosure   string `xml:"enclosure"`
	Guid        string `xml:"guid"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Source      string `xml:"source"`
	Title       string `xml:"title"`
}
