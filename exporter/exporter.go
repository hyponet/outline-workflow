package exporter

import (
	"github.com/hyponet/outline-workflow/exporter/alfred"
	"github.com/hyponet/outline-workflow/exporter/jsons"
	"github.com/hyponet/outline-workflow/outline"
)

var (
	Type = alfred.Type
)

type Exporter interface {
	AddResource(resource ...outline.Resource)
	Flush() error
}

func New() Exporter {
	switch Type {
	case alfred.Type:
		return alfred.NewAlfredExporter()
	case jsons.Type:
		return jsons.NewJsonPrinter()
	}
	return nil
}
