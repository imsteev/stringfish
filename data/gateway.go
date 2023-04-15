package data

// TODO: configure persistence
type Gateway struct{}

func (g Gateway) GetAllSubscriptions() []Subscription {
	var subs []Subscription
	for _, v := range Subscriptions {
		subs = append(subs, v)
	}
	return subs
}

func (g Gateway) AddSubscription(source string, sourceType SourceType) {
	Subscriptions[source] = Subscription{Source: source, Type: sourceType}
}
