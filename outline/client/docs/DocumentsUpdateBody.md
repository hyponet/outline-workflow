# DocumentsUpdateBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the document. Either the UUID or the urlId is acceptable. | [default to null]
**Title** | **string** | The title of the document. | [optional] [default to null]
**Text** | **string** | The body of the document, may contain markdown formatting. | [optional] [default to null]
**Append** | **bool** | If true the text field will be appended to the end of the existing document, rather than the default behavior of replacing it. This is potentially useful for things like logging into a document. | [optional] [default to null]
**Publish** | **bool** | Whether this document should be published and made visible to other team members, if a draft | [optional] [default to null]
**Done** | **bool** | Whether the editing session has finished, this will trigger any notifications. This property will soon be deprecated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

