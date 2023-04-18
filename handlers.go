package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"os"
	hackernews "stringfish/clients"
	"stringfish/data"
	"stringfish/feed"
	"time"
)

// TODO: parametrized route for GET + POST
func HandleSubscriptions(w http.ResponseWriter, r *http.Request) {
	g := data.Gateway{}

	if r.Method == http.MethodPost {
		var s data.Subscription
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			reject(w, http.StatusInternalServerError, fmt.Errorf("problem decoding: %s", err))
			return
		}
		g.AddSubscription(s.Source, s.Type)
		respondJson(w, s)
		return
	}

	type subscription struct {
		data.Subscription
		GetRssLink string
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "http://localhost:8080"
	}

	protocol := []subscription{}
	for _, sub := range g.GetAllSubscriptions() {
		protocol = append(protocol, subscription{
			Subscription: sub,
			GetRssLink:   fmt.Sprintf("%s/rss?source=%s&type=%s", host, url.QueryEscape(sub.Source), sub.Type),
		})
	}
	respondJson(w, protocol)
}

// TODO: parametrized route
func HandleRSS(w http.ResponseWriter, r *http.Request) {
	var (
		params     = r.URL.Query()
		source     = params.Get("source")
		sourceType = params.Get("type")
	)

	switch sourceType {
	case string(data.Hackernews):
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
	case string(data.XmlLink):
		http.Redirect(w, r, source, http.StatusSeeOther)
	default:
		reject(w, http.StatusBadRequest, fmt.Errorf("unsupported source type: %s", sourceType))
		return
	}
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
