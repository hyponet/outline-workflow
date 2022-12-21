# Team

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object. | [optional] [default to null]
**Name** | **string** | The name of this team, it is usually auto-generated when the first SSO connection is made but can be changed if neccessary. | [optional] [default to null]
**AvatarUrl** | **string** | The URL for the image associated with this team, it will be displayed in the team switcher and in the top left of the knowledge base along with the name. | [optional] [default to null]
**Sharing** | **bool** | Whether this team has share links globally enabled. If this value is false then all sharing UI and APIs are disabled. | [optional] [default to null]
**DefaultCollectionId** | **string** | If set then the referenced collection is where users will be redirected to after signing in instead of the Home screen | [optional] [default to null]
**DefaultUserRole** | **string** | If set then this role will be used as the default for users that signup via SSO | [optional] [default to null]
**MemberCollectionCreate** | **bool** | Whether members are allowed to create new collections. If false then only admins can create collections. | [optional] [default to null]
**DocumentEmbeds** | **bool** | Whether this team has embeds in documents globally enabled. It can be disabled to reduce potential data leakage to third parties. | [optional] [default to null]
**CollaborativeEditing** | **bool** | Whether this team has collaborative editing in documents globally enabled. | [optional] [default to null]
**InviteRequired** | **bool** | Whether an invite is required to join this team, if false users may join with a linked SSO provider. | [optional] [default to null]
**AllowedDomains** | **[]string** |  | [optional] [default to null]
**GuestSignin** | **bool** | Whether this team has guest signin enabled. Guests can signin with an email address and are not required to have a Google Workspace/Slack SSO account once invited. | [optional] [default to null]
**Subdomain** | **string** | Represents the subdomain at which this team&#x27;s knowledge base can be accessed. | [optional] [default to null]
**Url** | **string** | The fully qualified URL at which this team&#x27;s knowledge base can be accessed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

