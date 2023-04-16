package feed

import (
	"fmt"
	"math"
	"sort"
	hackernews "stringfish/clients"
	"stringfish/lib/rss"
	"time"
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

	itemChunks := chunks(user.Submitted, 50)

	workers := make(chan []int, len(itemChunks))
	results := make(chan []hackernews.Item, len(itemChunks))

	for i, chunk := range itemChunks {
		workers <- chunk
		go func(i int) {
			ids := <-workers
			items := make([]hackernews.Item, len(ids))

			for _, submittedID := range ids {
				item, err := h.Client.GetItem(submittedID)
				// TODO: how to handle failure?
				if err != nil || item.Deleted || item.Id == 0 {
					continue
				}
				items = append(items, item)
			}
			results <- items
		}(i)
	}

	var items []hackernews.Item
	for range itemChunks {
		items = append(items, <-results...)
	}

	sort.Slice(items, func(i int, j int) bool {
		return items[j].Time < items[i].Time
	})

	for _, item := range items {
		r.Channel.Items = append(r.Channel.Items, makeRssItem(item))
	}

	return &r, nil
}

func chunks(arr []int, chunkSize int) [][]int {
	numChunks := int(math.Ceil(float64(len(arr)) / (float64(chunkSize))))

	chunks := make([][]int, numChunks)
	i := 0
	for i < len(arr) {
		chunkIdx := i / chunkSize
		chunks[chunkIdx] = append(chunks[chunkIdx], arr[i])
		i++
	}

	return chunks
}

func makeRssItem(i hackernews.Item) rss.Item {
	var r rss.Item

	switch i.Type {
	case "story":
		r = rss.Item{
			Title:  i.Title,
			Author: i.By,
			Link:   i.Url,
		}
	case "comment":
		r = rss.Item{
			Description: i.Text,
			Author:      i.By,
		}
	case "ask":
		r = rss.Item{
			Title:       i.Title,
			Description: i.Text,
			Author:      i.By,
		}
	case "job":
		r = rss.Item{
			Title:       i.Title,
			Description: i.Text,
			Author:      i.By,
			Link:        i.Url,
		}
	case "poll":
		r = rss.Item{
			Title:       i.Title,
			Description: i.Text,
			Link:        i.Url,
		}
	case "pollopt":
		r = rss.Item{
			Description: i.Text,
			Author:      i.By,
		}
	default:
		return rss.Item{}
	}

	r.PubDate = time.Unix(int64(i.Time), 0).String()
	r.Guid = fmt.Sprintf("%s-%d", i.Type, i.Id)
	return r

}
