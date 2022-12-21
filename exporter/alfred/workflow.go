package alfred

import (
	"github.com/deanishe/awgo"
	"github.com/hyponet/outline-workflow/outline"
)

const (
	Type = "alfred"
)

type Exporter struct {
	*aw.Workflow

	resources []outline.Resource
}

func (e *Exporter) AddResource(resource ...outline.Resource) {
	e.resources = append(e.resources, resource...)
}

func (e *Exporter) Flush() error {
	e.Workflow.Run(func() {
		for _, res := range e.resources {
			e.addItem(res)
		}
		e.SendFeedback()
	})
	return nil
}

func (e *Exporter) addItem(res outline.Resource) {
	item := e.NewItem(res.Name).Subtitle(res.Describe).Autocomplete(res.Name)
	if res.URL != "" {
		item.ActionForType(aw.TypeURL, res.URL)
	}
}

func NewAlfredExporter() *Exporter {
	return &Exporter{Workflow: aw.New()}
}
