package feed

import (
	"encoding/xml"
	"fmt"
	hackernews "rss/clients"
	"rss/lib/rss"
)

type HackerNewsFeed struct {
	Username string
}

func (h HackerNewsFeed) GenerateXml() string {
	c := hackernews.HackerNewsClient{}
	u, _ := c.GetUser(h.Username)

	r := rss.Rss{
		Version: "2.0",
		Channel: rss.Channel{
			Title:       u.Id,
			Description: u.About,
			Link:        fmt.Sprintf("https://news.ycombinator.com/user?id=%s", u.Id),
		},
	}

	for _, submittedID := range u.Submitted {
		item, _ := c.GetItem(submittedID)
		r.Channel.Items = append(r.Channel.Items, makeRssItem(item))
	}

	bytes, _ := xml.Marshal(r)

	return xml.Header + string(bytes)
}

func makeRssItem(i hackernews.Item) rss.Item {
	var r rss.Item
	switch i.Type {
	case "story":
		r = rss.Item{
			Title:   i.Title,
			Author:  i.By,
			Link:    i.Url,
			PubDate: i.Time,
		}
	case "comment":
		r = rss.Item{
			Description: i.Text,
			Author:      i.By,
			PubDate:     i.Time,
		}
	case "ask":
		r = rss.Item{
			Title:       i.Title,
			Description: i.Text,
			Author:      i.By,
			PubDate:     i.Time,
		}
	case "job":
		r = rss.Item{
			Title:       i.Title,
			Description: i.Text,
			Author:      i.By,
			Link:        i.Url,
			PubDate:     i.Time,
		}
	case "poll":
		r = rss.Item{
			Title:       i.Title,
			Description: i.Text,
			Link:        i.Url,
			PubDate:     i.Time,
		}
	case "pollopt":
		r = rss.Item{
			Description: i.Text,
			Author:      i.By,
			PubDate:     i.Time,
		}
	default:
		return rss.Item{}
	}
	return r

}
