package jsons

import (
	"encoding/json"
	"fmt"
	"github.com/hyponet/outline-workflow/outline"
)

const (
	Type = "json"
)

type JsonPrinter struct {
	resources []outline.Resource
}

func (j *JsonPrinter) AddResource(resource ...outline.Resource) {
	j.resources = append(j.resources, resource...)
}

func (j *JsonPrinter) Flush() error {
	data, err := json.MarshalIndent(j.resources, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func NewJsonPrinter() *JsonPrinter {
	return &JsonPrinter{resources: []outline.Resource{}}
}
