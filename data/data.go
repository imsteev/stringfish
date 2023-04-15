package data

type SourceType string

// Supported source types
const (
	Hackernews SourceType = "hackernews"
)

type Subscription struct {
	Source string
	Type   SourceType
}

var Subscriptions map[string]Subscription

// TODO: persistence
func init() {
	Subscriptions = make(map[string]Subscription)
}
