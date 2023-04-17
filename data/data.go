package data

type SourceType string

// Supported source types
const (
	Hackernews SourceType = "hackernews"
)

type Subscription struct {
	// Type represents where an RSS feed comes from.
	Type SourceType

	// Source is contextualized by Type. For example, if Type is "hackernews",
	// it should be a HackerNews username.
	Source string
}

var subscriptions map[string]Subscription

// TODO: persistence
func init() {
	subscriptions = make(map[string]Subscription)
}
