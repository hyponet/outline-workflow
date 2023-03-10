package outline

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/hyponet/outline-workflow/outline/client"
	"strings"
)

type Outline struct {
	client *client.APIClient
}

func (o *Outline) ListCollections(ctx context.Context) ([]Resource, error) {
	collections, _, err := o.client.CollectionsApi.CollectionsListPost(ctx, nil)
	if err != nil {
		return nil, err
	}
	results := make([]Resource, 0, len(collections.Data))
	for _, col := range collections.Data {
		results = append(results, Resource{
			ID:       col.Id,
			Name:     col.Name,
			Describe: col.Description,
			Type:     ResourceCollection,
		})
	}
	return results, err
}

func (o *Outline) SearchDocuments(ctx context.Context, query string) ([]Resource, error) {
	filter := map[string]string{
		"query": query,
	}
	documents, _, err := o.client.DocumentsApi.DocumentsSearchPost(ctx, &client.DocumentsApiDocumentsSearchPostOpts{Body: optional.NewInterface(filter)})
	if err != nil {
		return nil, err
	}

	results := make([]Resource, 0, len(documents.Data))
	for _, doc := range documents.Data {
		results = append(results, Resource{
			ID:       doc.Document.Id,
			Name:     doc.Document.Title,
			Describe: documentDesc(doc.Document),
			Type:     ResourceDocument,
			URL:      documentURL(doc.Document),
		})
	}
	return results, err
}

func New(config ConnConfig) (*Outline, error) {
	if Host == "" {
		return nil, fmt.Errorf("Host is requested ")
	}
	if ApiToken == "" {
		return nil, fmt.Errorf("ApiToken is requested ")
	}
	o := &Outline{client: client.NewAPIClient(&client.Configuration{
		BasePath:      fmt.Sprintf("%s/api", Host),
		DefaultHeader: map[string]string{"Authorization": "Bearer " + ApiToken},
		UserAgent:     "outline-workflow",
		HTTPClient:    NewHttpClient(config),
	})}
	return o, nil
}

func documentURL(doc *client.Document) string {
	return fmt.Sprintf("%s/doc/%s-%s", Host, doc.Id, doc.UrlId)
}

func documentDesc(doc *client.Document) string {
	lines := strings.Split(doc.Text, "\n")
	for _, l := range lines {
		content := strings.TrimSpace(l)
		if strings.HasPrefix(content, "#") {
			continue
		}
		if len(content) == 0 {
			continue
		}
		return content
	}
	return doc.Text
}
