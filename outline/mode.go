package outline

const (
	ResourceCollection = "collection"
	ResourceDocument   = "document"
	ResourceTemplate   = "template"
	ResourceEvent      = "event"
)

type Resource struct {
	ID       string
	Name     string
	Describe string
	Type     string
	URL      string
}
