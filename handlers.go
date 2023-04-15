package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	hackernews "stringfish/clients"
	"stringfish/data"
	"stringfish/feed"
	"time"
)

// TODO: parametrized route for GET + POST
func HandleSubscriptions(w http.ResponseWriter, r *http.Request) {
	var s data.Subscription
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		reject(w, http.StatusInternalServerError, err)
		return
	}

	g := data.Gateway{}

	if r.Method == http.MethodPost {
		if s.Source != string(data.Hackernews) {
			reject(w, http.StatusBadRequest, fmt.Errorf("%s source type not supported", s.Type))
			return
		}
		g.AddSubscription(s.Source, data.Hackernews)
		s.Type = data.Hackernews
		respondJson(w, s)
		return
	}

	respondJson(w, g.GetAllSubscriptions())
}

// TODO: parametrized route
func HandleRSS(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	source, sourceType := params.Get("source"), params.Get("type")

	if sourceType != string(data.Hackernews) {
		reject(w, http.StatusBadRequest, fmt.Errorf("unsupported source type: %s", sourceType))
		return
	}

	feed := feed.HackerNewsFeed{
		Username: source,
		Client: hackernews.HackerNewsClient{
			Timeout: 120 * time.Second,
		},
	}

	rss, err := feed.GenerateRss()
	if err != nil {
		reject(w, http.StatusInternalServerError, fmt.Errorf("could not generate RSS feed"))
		return
	}

	rssXml, err := xml.Marshal(rss)
	if err != nil {
		reject(w, http.StatusInternalServerError, fmt.Errorf("problem marshalling RSS feed"))
		return
	}

	w.Header().Add("Content-Type", "application/rss+xml")
	w.Write([]byte(xml.Header + string(rssXml)))
}

func respondJson(w http.ResponseWriter, body any) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(b)
}

func reject(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}
