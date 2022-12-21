# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionsAddGroupPost**](CollectionsApi.md#CollectionsAddGroupPost) | **Post** /collections.add_group | Add a group to a collection
[**CollectionsAddUserPost**](CollectionsApi.md#CollectionsAddUserPost) | **Post** /collections.add_user | Add a collection user
[**CollectionsCreatePost**](CollectionsApi.md#CollectionsCreatePost) | **Post** /collections.create | Create a collection
[**CollectionsDeletePost**](CollectionsApi.md#CollectionsDeletePost) | **Post** /collections.delete | Delete a collection
[**CollectionsExportAllPost**](CollectionsApi.md#CollectionsExportAllPost) | **Post** /collections.export_all | Export all collections
[**CollectionsExportPost**](CollectionsApi.md#CollectionsExportPost) | **Post** /collections.export | Export a collection
[**CollectionsGroupMembershipsPost**](CollectionsApi.md#CollectionsGroupMembershipsPost) | **Post** /collections.group_memberships | List all collection group members
[**CollectionsInfoPost**](CollectionsApi.md#CollectionsInfoPost) | **Post** /collections.info | Retrieve a collection
[**CollectionsListPost**](CollectionsApi.md#CollectionsListPost) | **Post** /collections.list | List all collections
[**CollectionsMembershipsPost**](CollectionsApi.md#CollectionsMembershipsPost) | **Post** /collections.memberships | List all collection memberships
[**CollectionsRemoveGroupPost**](CollectionsApi.md#CollectionsRemoveGroupPost) | **Post** /collections.remove_group | Remove a collection group
[**CollectionsRemoveUserPost**](CollectionsApi.md#CollectionsRemoveUserPost) | **Post** /collections.remove_user | Remove a collection user
[**CollectionsUpdatePost**](CollectionsApi.md#CollectionsUpdatePost) | **Post** /collections.update | Update a collection

# **CollectionsAddGroupPost**
> InlineResponse2009 CollectionsAddGroupPost(ctx, optional)
Add a group to a collection

This method allows you to give all members in a group access to a collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsAddGroupPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsAddGroupPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsAddGroupBody**](CollectionsAddGroupBody.md)|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsAddUserPost**
> InlineResponse2007 CollectionsAddUserPost(ctx, optional)
Add a collection user

This method allows you to add a user membership to the specified collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsAddUserPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsAddUserPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsAddUserBody**](CollectionsAddUserBody.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsCreatePost**
> InlineResponse2006 CollectionsCreatePost(ctx, optional)
Create a collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsCreateBody**](CollectionsCreateBody.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsDeletePost**
> InlineResponse2001 CollectionsDeletePost(ctx, optional)
Delete a collection

Delete a collection and all of its documents. This action can’t be undone so please be careful.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsDeleteBody**](CollectionsDeleteBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsExportAllPost**
> InlineResponse20011 CollectionsExportAllPost(ctx, )
Export all collections

Triggers a bulk export of all documents in markdown format and their attachments. If documents are nested then they will be nested in folders inside the zip file. The endpoint returns a `FileOperation` that can be queried to track the progress of the export and get the url for the final file.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsExportPost**
> InlineResponse20011 CollectionsExportPost(ctx, optional)
Export a collection

Triggers a bulk export of the collection in markdown format and their attachments. If documents are nested then they will be nested in folders inside the zip file. The endpoint returns a `FileOperation` that can be queried to track the progress of the export and get the url for the final file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsExportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsExportPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsExportBody**](CollectionsExportBody.md)|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsGroupMembershipsPost**
> InlineResponse20010 CollectionsGroupMembershipsPost(ctx, optional)
List all collection group members

This method allows you to list a collections group memberships. This is the list of groups that have been given access to the collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsGroupMembershipsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsGroupMembershipsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsGroupMembershipsBody**](CollectionsGroupMembershipsBody.md)|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsInfoPost**
> InlineResponse2004 CollectionsInfoPost(ctx, optional)
Retrieve a collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsInfoPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsInfoPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsInfoBody**](CollectionsInfoBody.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsListPost**
> InlineResponse2005 CollectionsListPost(ctx, optional)
List all collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Pagination**](Pagination.md)|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsMembershipsPost**
> InlineResponse2008 CollectionsMembershipsPost(ctx, optional)
List all collection memberships

This method allows you to list a collections individual memberships. It's important to note that memberships returned from this endpoint do not include group memberships.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsMembershipsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsMembershipsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsMembershipsBody**](CollectionsMembershipsBody.md)|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsRemoveGroupPost**
> InlineResponse2001 CollectionsRemoveGroupPost(ctx, optional)
Remove a collection group

This method allows you to revoke all members in a group access to a collection. Note that members of the group may still retain access through other groups or individual memberships.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsRemoveGroupPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsRemoveGroupPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsRemoveGroupBody**](CollectionsRemoveGroupBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsRemoveUserPost**
> InlineResponse2001 CollectionsRemoveUserPost(ctx, optional)
Remove a collection user

This method allows you to remove a user from the specified collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsRemoveUserPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsRemoveUserPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsRemoveUserBody**](CollectionsRemoveUserBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsUpdatePost**
> InlineResponse2006 CollectionsUpdatePost(ctx, optional)
Update a collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsUpdatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsUpdatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CollectionsUpdateBody**](CollectionsUpdateBody.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

