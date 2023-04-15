package feed

import (
	"fmt"
	hackernews "stringfish/clients"
	"stringfish/lib/rss"
)

type HackerNewsFeed struct {
	Username string
	Client   hackernews.HackerNewsClient
}

func (h HackerNewsFeed) GenerateRss() (*rss.Rss, error) {
	user, err := h.Client.GetUser(h.Username)
	if err != nil {
		return nil, err
	}

	profileLink := fmt.Sprintf("https://news.ycombinator.com/user?id=%s", user.Id)

	r := rss.Rss{
		Version: "2.0",
		Channel: rss.Channel{
			Title:       user.Id,
			Description: user.About,
			Link:        profileLink,
		},
	}

	for _, submittedID := range user.Submitted {
		item, err := h.Client.GetItem(submittedID)
		if err != nil {
			return nil, err
		}
		r.Channel.Items = append(r.Channel.Items, makeRssItem(item))
	}

	return &r, nil
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