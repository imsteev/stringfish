package hackernews

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseUrl string = "https://hacker-news.firebaseio.com/v0"

type HackerNewsClient struct {
	Timeout time.Duration
}

type User struct {
	Id        string
	Created   string
	Karma     int
	About     string
	Submitted []int
}

type Item struct {
	Id          string
	Deleted     bool
	Type        string
	By          string
	Time        string
	Text        string
	Dead        bool
	Parent      string
	Poll        any
	Kids        []int
	Url         string
	Score       int
	Title       string
	Parts       []any
	Descendants int
}

func (h *HackerNewsClient) GetUser(username string) (User, error) {
	client := http.Client{Timeout: h.Timeout}
	r, err := client.Get(fmt.Sprintf("%s/user/%s.json?print=pretty", baseUrl, username))
	if err != nil {
		return User{}, err
	}
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	return u, nil
}

func (h *HackerNewsClient) GetItem(id int) (Item, error) {
	client := http.Client{}
	r, err := client.Get(fmt.Sprintf("%s/item/%d.json?print=pretty", baseUrl, id))
	if err != nil {
		return Item{}, err
	}
	var i Item
	json.NewDecoder(r.Body).Decode(&i)
	return i, nil
}
