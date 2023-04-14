package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss/data"
	"rss/lib"
)

var Decoder = json.NewDecoder

func HandleSubscriptions(w http.ResponseWriter, r *http.Request) {
	var s data.Subscription
	if err := Decoder(r.Body).Decode(&s); err != nil {
		fmt.Printf("%s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		if s.Source != "" {
			data.AddSubscription(s.Source, data.Hackernews)
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

func HandleRSS(w http.ResponseWriter, r *http.Request) {

	for _, s := range data.Subscriptions {
		switch s.Type {
		case "hackernews":
			hnRss := lib.HackerNewsRss{
				Username: s.Source,
			}
			w.Write([]byte(hnRss.GenerateXml()))
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
