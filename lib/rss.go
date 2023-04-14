package lib

import (
	"fmt"
	"rss/clients"
	"rss/data"
)

type HackerNewsRss struct {
	Username string
}

func (h HackerNewsRss) GenerateXml() string {
	c := clients.HackerNewsClient{}

	sub, ok := data.Subscriptions[h.Username]
	if !ok || sub.Type != "hackernews" {
		return ""
	}

	// Get hackernews user
	u, _ := c.GetUser(sub.Source)

	items := ""
	for _, submittedID := range u.Submitted {
		item, _ := c.GetItem(submittedID)
		items += h.makeRssItem(item).generateXml()
	}

	userLink := fmt.Sprintf("https://news.ycombinator.com/user?id=%s", u.Id)
	return fmt.Sprintf(`<rss version="2.0"><channel><title>%s</title><description>%s</description><link>%s</link>%s</channel>`,
		u.Id, u.About, userLink, items)
}

func (h HackerNewsRss) makeRssItem(i clients.Item) rssItem {
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
	Author      string
	Category    string
	Comments    string
	Description string
	Enclosure   string
	Guid        string
	Link        string
	PubDate     string
	Source      string
	Title       string
}

func (r rssItem) generateXml() string {
	xmlStr := "<item>"
	if r.Author != "" {
		xmlStr += makeRssElem("author", r.Author)
	}
	if r.Category != "" {
		xmlStr += makeRssElem("cateogry", r.Category)
	}
	if r.Comments != "" {
		xmlStr += makeRssElem("comments", r.Comments)
	}
	if r.Description != "" {
		xmlStr += makeRssElem("description", r.Description)
	}
	if r.Enclosure != "" {
		xmlStr += makeRssElem("enclosure", r.Enclosure)
	}
	if r.Guid != "" {
		xmlStr += makeRssElem("guid", r.Guid)
	}
	if r.Link != "" {
		xmlStr += makeRssElem("link", r.Link)
	}
	if r.PubDate != "" {
		xmlStr += makeRssElem("pubDate", r.PubDate)
	}
	if r.Source != "" {
		xmlStr += makeRssElem("source", r.Source)
	}
	if r.Title != "" {
		xmlStr += makeRssElem("title", r.Title)
	}
	xmlStr += "</item>"
	return xmlStr
}

func makeRssElem(tag string, val string) string {
	return fmt.Sprintf(`<%s>%s</%s>`, tag, val, tag)
}
