# User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object. | [optional] [default to null]
**Name** | **string** | The name of this user, it is migrated from Slack or Google Workspace when the SSO connection is made but can be changed if neccessary. | [optional] [default to null]
**AvatarUrl** | **string** | The URL for the image associated with this user, it will be displayed in the application UI and email notifications. | [optional] [default to null]
**Email** | **string** | The email associated with this user, it is migrated from Slack or Google Workspace when the SSO connection is made but can be changed if neccessary. | [optional] [default to null]
**IsAdmin** | **bool** | Whether this user has admin permissions. | [optional] [default to null]
**IsSuspended** | **bool** | Whether this user has been suspended. | [optional] [default to null]
**LastActiveAt** | **string** | The last time this user made an API request, this value is updated at most every 5 minutes. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time that this user first signed in or was invited as a guest. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

