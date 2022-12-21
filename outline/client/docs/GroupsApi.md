# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsAddUserPost**](GroupsApi.md#GroupsAddUserPost) | **Post** /groups.add_user | Add a group member
[**GroupsCreatePost**](GroupsApi.md#GroupsCreatePost) | **Post** /groups.create | Create a group
[**GroupsDeletePost**](GroupsApi.md#GroupsDeletePost) | **Post** /groups.delete | Delete a group
[**GroupsInfoPost**](GroupsApi.md#GroupsInfoPost) | **Post** /groups.info | Retrieve a group
[**GroupsListPost**](GroupsApi.md#GroupsListPost) | **Post** /groups.list | List all groups
[**GroupsMembershipsPost**](GroupsApi.md#GroupsMembershipsPost) | **Post** /groups.memberships | List all group members
[**GroupsRemoveUserPost**](GroupsApi.md#GroupsRemoveUserPost) | **Post** /groups.remove_user | Remove a group member
[**GroupsUpdatePost**](GroupsApi.md#GroupsUpdatePost) | **Post** /groups.update | Update a group

# **GroupsAddUserPost**
> InlineResponse20026 GroupsAddUserPost(ctx, optional)
Add a group member

This method allows you to add a user to the specified group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsAddUserPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsAddUserPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsAddUserBody**](GroupsAddUserBody.md)|  | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsCreatePost**
> InlineResponse20024 GroupsCreatePost(ctx, optional)
Create a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsCreateBody**](GroupsCreateBody.md)|  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsDeletePost**
> InlineResponse2001 GroupsDeletePost(ctx, optional)
Delete a group

Deleting a group will cause all of its members to lose access to any collections the group has previously been added to. This action can’t be undone so please be careful.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsDeleteBody**](GroupsDeleteBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsInfoPost**
> InlineResponse20022 GroupsInfoPost(ctx, optional)
Retrieve a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsInfoPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsInfoPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsInfoBody**](GroupsInfoBody.md)|  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsListPost**
> InlineResponse20023 GroupsListPost(ctx, optional)
List all groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsListBody**](GroupsListBody.md)|  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsMembershipsPost**
> InlineResponse20025 GroupsMembershipsPost(ctx, optional)
List all group members

List and filter all the members in a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsMembershipsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsMembershipsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsMembershipsBody**](GroupsMembershipsBody.md)|  | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsRemoveUserPost**
> InlineResponse20027 GroupsRemoveUserPost(ctx, optional)
Remove a group member

This method allows you to remove a user from the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsRemoveUserPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsRemoveUserPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsRemoveUserBody**](GroupsRemoveUserBody.md)|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsUpdatePost**
> InlineResponse20024 GroupsUpdatePost(ctx, optional)
Update a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsUpdatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsUpdatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsUpdateBody**](GroupsUpdateBody.md)|  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

