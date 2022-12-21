# Share

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object. | [optional] [default to null]
**DocumentTitle** | **string** | Title of the shared document. | [optional] [default to null]
**DocumentUrl** | **string** | URL of the original document. | [optional] [default to null]
**Url** | **string** | URL of the publicly shared document. | [optional] [default to null]
**Published** | **bool** | If true the share can be loaded without a user account. | [optional] [default to null]
**IncludeChildDocuments** | **bool** | If to also give permission to view documents nested beneath this one. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Date and time when this share was created | [optional] [default to null]
**CreatedBy** | [***User**](User.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Date and time when this share was edited | [optional] [default to null]
**LastAccessedAt** | [**time.Time**](time.Time.md) | Date and time when this share was last viewed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

