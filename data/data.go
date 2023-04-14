package data

type Subscription struct {
	Source string
	Type   string
}

var Subscriptions map[string]Subscription

func init() {
	Subscriptions = make(map[string]Subscription)
	AddSubscription("dfern")
}

type Data struct{}

func AddSubscription(source string) {
	Subscriptions[source] = Subscription{Source: source, Type: "hackernews"}
}

func GetSubscriptions(source string) Subscription {
	return Subscriptions[source]
}
