# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesCreatePost**](SharesApi.md#SharesCreatePost) | **Post** /shares.create | Create a share
[**SharesInfoPost**](SharesApi.md#SharesInfoPost) | **Post** /shares.info | Retrieve a share object
[**SharesListPost**](SharesApi.md#SharesListPost) | **Post** /shares.list | List all shares
[**SharesRevokePost**](SharesApi.md#SharesRevokePost) | **Post** /shares.revoke | Revoke a share
[**SharesUpdatePost**](SharesApi.md#SharesUpdatePost) | **Post** /shares.update | Update a share

# **SharesCreatePost**
> InlineResponse20030 SharesCreatePost(ctx, optional)
Create a share

Creates a new share link that can be used by to access a document. If you request multiple shares for the same document with the same API key, the same share object will be returned. By default all shares are unpublished.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharesApiSharesCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiSharesCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SharesCreateBody**](SharesCreateBody.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SharesInfoPost**
> InlineResponse20030 SharesInfoPost(ctx, optional)
Retrieve a share object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharesApiSharesInfoPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiSharesInfoPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SharesInfoBody**](SharesInfoBody.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SharesListPost**
> InlineResponse20031 SharesListPost(ctx, optional)
List all shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharesApiSharesListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiSharesListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SharesListBody**](SharesListBody.md)|  | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SharesRevokePost**
> InlineResponse2001 SharesRevokePost(ctx, optional)
Revoke a share

Makes the share link inactive so that it can no longer be used to access the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharesApiSharesRevokePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiSharesRevokePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SharesRevokeBody**](SharesRevokeBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SharesUpdatePost**
> InlineResponse20030 SharesUpdatePost(ctx, optional)
Update a share

Allows changing an existing shares published status, which removes authentication and makes it available to anyone with the link.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharesApiSharesUpdatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiSharesUpdatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SharesUpdateBody**](SharesUpdateBody.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

