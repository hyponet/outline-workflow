# EventsListBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sort** | **string** |  | [optional] [default to null]
**Direction** | **string** |  | [optional] [default to null]
**Offset** | **float64** |  | [optional] [default to null]
**Limit** | **float64** |  | [optional] [default to null]
**Name** | **string** | Filter to a specific event, e.g. \&quot;collections.create\&quot;. Event names are in the format \&quot;objects.verb\&quot; | [optional] [default to null]
**ActorId** | **string** | Filter to events performed by the selected user | [optional] [default to null]
**DocumentId** | **string** | Filter to events performed in the selected document | [optional] [default to null]
**CollectionId** | **string** | Filter to events performed in the selected collection | [optional] [default to null]
**AuditLog** | **bool** | Whether to return detailed events suitable for an audit log. Without this flag less detailed event types will be returned. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

