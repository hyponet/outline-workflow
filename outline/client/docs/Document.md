# Document

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object. | [optional] [default to null]
**CollectionId** | **string** | Identifier for the associated collection. | [optional] [default to null]
**ParentDocumentId** | **string** | Identifier for the document this is a child of, if any. | [optional] [default to null]
**Title** | **string** | The title of the document. | [optional] [default to null]
**FullWidth** | **bool** | Whether this document should be displayed in a full-width view. | [optional] [default to null]
**Emoji** | **string** | An emoji associated with the document. | [optional] [default to null]
**Text** | **string** | The text content of the document, contains markdown formatting | [optional] [default to null]
**UrlId** | **string** | A short unique ID that can be used to identify the document as an alternative to the UUID | [optional] [default to null]
**Collaborators** | [**[]User**](User.md) |  | [optional] [default to null]
**Pinned** | **bool** | Whether this document is pinned in the collection | [optional] [default to null]
**Template** | **bool** | Whether this document is a template | [optional] [default to null]
**TemplateId** | **string** | Unique identifier for the template this document was created from, if any | [optional] [default to null]
**Revision** | **float64** | A number that is auto incrementing with every revision of the document that is saved | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time that this object was created | [optional] [default to null]
**CreatedBy** | [***User**](User.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time that this object was last changed | [optional] [default to null]
**UpdatedBy** | [***User**](User.md) |  | [optional] [default to null]
**PublishedAt** | [**time.Time**](time.Time.md) | The date and time that this object was published | [optional] [default to null]
**ArchivedAt** | [**time.Time**](time.Time.md) | The date and time that this object was archived | [optional] [default to null]
**DeletedAt** | [**time.Time**](time.Time.md) | The date and time that this object was deleted | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

