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

var subscriptions map[string]Subscription

// TODO: persistence
func init() {
	subscriptions = make(map[string]Subscription)
}
