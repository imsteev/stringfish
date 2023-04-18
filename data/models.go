package data

type SourceType string

// Supported source types
const (
	Hackernews SourceType = "hackernews"
	XmlLink    SourceType = "xml_link"
)

type Subscription struct {
	// Type represents where an RSS feed comes from.
	Type SourceType

	// Source is contextualized by Type. For example, if Type is "hackernews",
	// it should be a HackerNews username.
	Source string
}
