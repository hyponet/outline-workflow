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
	item := e.NewItem(res.Name).Subtitle(res.Describe).UID(res.ID).Valid(true)

	switch res.Type {
	case outline.ResourceDocument:
		item.Arg(res.URL)
		item.ActionForType(aw.TypeURL, res.URL)
	case outline.ResourceCollection, outline.ResourceTemplate:
		item.Arg(res.ID)
	}
}

func NewAlfredExporter() *Exporter {
	return &Exporter{Workflow: aw.New()}
}
