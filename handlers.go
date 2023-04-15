package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss/data"
	"rss/lib/feed"
)

// TODO: parametrized route for GET + POST
func HandleSubscriptions(w http.ResponseWriter, r *http.Request) {
	var s data.Subscription
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		fmt.Printf("%s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		if s.Source != "" {
			data.AddSubscription(s.Source, data.Hackernews)
			s.Type = data.Hackernews
			respond(w, s)
		}
	} else {
		var subs []string
		for k := range data.Subscriptions {
			subs = append(subs, k)
		}
		respond(w, subs)
	}
}

// TODO: parametrized route
// TODO: content-type
func HandleRSS(w http.ResponseWriter, r *http.Request) {
	for _, s := range data.Subscriptions {
		switch s.Type {
		case "hackernews":
			feed := feed.HackerNewsFeed{Username: s.Source}
			w.Write([]byte(feed.GenerateXml()))
		}
	}
}

func respond(w http.ResponseWriter, body any) {
	b, err := json.Marshal(body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(b)
}
