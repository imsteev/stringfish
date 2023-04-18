package main

import (
	"log"
	"net/http"
)

const SERVER_ADDR = ":8080"

func main() {

	http.HandleFunc("/", LogHandler(HandleSubscriptions))
	http.HandleFunc("/subscriptions", LogHandler(HandleSubscriptions))
	http.HandleFunc("/rss", LogHandler(HandleRSS))

	log.Printf("listening on: http://%s\n", SERVER_ADDR)
	if err := http.ListenAndServe(SERVER_ADDR, nil); err != nil {
		log.Fatal(err)
	}

}
