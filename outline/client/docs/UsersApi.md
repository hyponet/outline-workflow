# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersActivatePost**](UsersApi.md#UsersActivatePost) | **Post** /users.activate | Activate a user
[**UsersDeletePost**](UsersApi.md#UsersDeletePost) | **Post** /users.delete | Delete a user
[**UsersDemotePost**](UsersApi.md#UsersDemotePost) | **Post** /users.demote | Demote a user
[**UsersInfoPost**](UsersApi.md#UsersInfoPost) | **Post** /users.info | Retrieve a user
[**UsersInvitePost**](UsersApi.md#UsersInvitePost) | **Post** /users.invite | Invite users
[**UsersListPost**](UsersApi.md#UsersListPost) | **Post** /users.list | List all users
[**UsersPromotePost**](UsersApi.md#UsersPromotePost) | **Post** /users.promote | Promote a user
[**UsersSuspendPost**](UsersApi.md#UsersSuspendPost) | **Post** /users.suspend | Suspend a user
[**UsersUpdatePost**](UsersApi.md#UsersUpdatePost) | **Post** /users.update | Update a user

# **UsersActivatePost**
> InlineResponse20033 UsersActivatePost(ctx, optional)
Activate a user

Activating a previously suspended user allows them to signin again. Users that are activated will cause billing totals to be re-calculated in the hosted version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersActivatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersActivatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersActivateBody**](UsersActivateBody.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersDeletePost**
> InlineResponse2001 UsersDeletePost(ctx, optional)
Delete a user

Deleting a user removes the object entirely. In almost every circumstance it is preferable to suspend a user, as a deleted user can be recreated by signing in with SSO again.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersDeleteBody**](UsersDeleteBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersDemotePost**
> InlineResponse20033 UsersDemotePost(ctx, optional)
Demote a user

Demote a team admin to regular user permissions. This endpoint is only available for admin users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersDemotePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersDemotePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersDemoteBody**](UsersDemoteBody.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersInfoPost**
> InlineResponse20033 UsersInfoPost(ctx, optional)
Retrieve a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersInfoPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersInfoPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersInfoBody**](UsersInfoBody.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersInvitePost**
> InlineResponse20032 UsersInvitePost(ctx, optional)
Invite users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersInvitePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersInvitePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersInviteBody**](UsersInviteBody.md)|  | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersListPost**
> InlineResponse20034 UsersListPost(ctx, optional)
List all users

List and filter all the users in the team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersListBody**](UsersListBody.md)|  | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPromotePost**
> InlineResponse20033 UsersPromotePost(ctx, optional)
Promote a user

Promote a user to be a team admin. This endpoint is only available for admin users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersPromotePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPromotePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersPromoteBody**](UsersPromoteBody.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSuspendPost**
> InlineResponse20033 UsersSuspendPost(ctx, optional)
Suspend a user

Suspending a user prevents the user from signing in. Users that are suspended are also not counted against billing totals in the hosted version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersSuspendPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersSuspendPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersSuspendBody**](UsersSuspendBody.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUpdatePost**
> InlineResponse20033 UsersUpdatePost(ctx, optional)
Update a user

Update a users name or avatar. No `id` is required as it is only possible to update the current user at this time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersUpdatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUpdatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersUpdateBody**](UsersUpdateBody.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

