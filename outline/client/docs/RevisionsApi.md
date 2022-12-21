# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RevisionsInfoPost**](RevisionsApi.md#RevisionsInfoPost) | **Post** /revisions.info | Retrieve a revision
[**RevisionsListPost**](RevisionsApi.md#RevisionsListPost) | **Post** /revisions.list | List all revisions

# **RevisionsInfoPost**
> InlineResponse20028 RevisionsInfoPost(ctx, optional)
Retrieve a revision

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RevisionsApiRevisionsInfoPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RevisionsApiRevisionsInfoPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RevisionsInfoBody**](RevisionsInfoBody.md)|  | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevisionsListPost**
> InlineResponse20029 RevisionsListPost(ctx, optional)
List all revisions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RevisionsApiRevisionsListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RevisionsApiRevisionsListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RevisionsListBody**](RevisionsListBody.md)|  | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

