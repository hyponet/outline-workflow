# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewsCreatePost**](ViewsApi.md#ViewsCreatePost) | **Post** /views.create | Create a view
[**ViewsListPost**](ViewsApi.md#ViewsListPost) | **Post** /views.list | List all views

# **ViewsCreatePost**
> InlineResponse20036 ViewsCreatePost(ctx, optional)
Create a view

Creates a new view for a document. This is documented in the interests of thoroughness however it is recommended that views are not created from outside of the Outline UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ViewsApiViewsCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiViewsCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ViewsCreateBody**](ViewsCreateBody.md)|  | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewsListPost**
> InlineResponse20035 ViewsListPost(ctx, optional)
List all views

List all users that have viewed a document and the overall view count.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ViewsApiViewsListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiViewsListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ViewsListBody**](ViewsListBody.md)|  | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

