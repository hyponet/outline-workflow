# Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object. | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**ModelId** | **string** | Identifier for the object this event is associated with when it is not one of document, collection, or user. | [optional] [default to null]
**ActorId** | **string** | The user that performed the action. | [optional] [default to null]
**ActorIpAddress** | **string** | The ip address the action was performed from. This field is only returned when the &#x60;auditLog&#x60; boolean is true. | [optional] [default to null]
**CollectionId** | **string** | Identifier for the associated collection, if any | [optional] [default to null]
**DocumentId** | **string** | Identifier for the associated document, if any | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time that this event was created | [optional] [default to null]
**Data** | [***interface{}**](interface{}.md) | Additional unstructured data associated with the event | [optional] [default to null]
**Actor** | [***User**](User.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

