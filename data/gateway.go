package data

type Gateway struct{}

func (g Gateway) GetAllSubscriptions() []Subscription {
	var subs []Subscription
	for _, v := range subscriptions {
		subs = append(subs, v)
	}
	return subs
}

func (g Gateway) AddSubscription(source string, sourceType SourceType) {
	subscriptions[source] = Subscription{Source: source, Type: sourceType}
}
