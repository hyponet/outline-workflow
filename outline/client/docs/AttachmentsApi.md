# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachmentsCreatePost**](AttachmentsApi.md#AttachmentsCreatePost) | **Post** /attachments.create | Create an attachment
[**AttachmentsDeletePost**](AttachmentsApi.md#AttachmentsDeletePost) | **Post** /attachments.delete | Delete an attachment
[**AttachmentsRedirectPost**](AttachmentsApi.md#AttachmentsRedirectPost) | **Post** /attachments.redirect | Retrieve an attachment

# **AttachmentsCreatePost**
> InlineResponse200 AttachmentsCreatePost(ctx, optional)
Create an attachment

Creating an attachment object creates a database record and returns the inputs needed to generate a signed url and upload the file from the client to cloud storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AttachmentsApiAttachmentsCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AttachmentsApiAttachmentsCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AttachmentsCreateBody**](AttachmentsCreateBody.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachmentsDeletePost**
> InlineResponse2001 AttachmentsDeletePost(ctx, optional)
Delete an attachment

Deleting an attachment is permanant. It will not delete references or links to the attachment that may exist in your documents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AttachmentsApiAttachmentsDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AttachmentsApiAttachmentsDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AttachmentsDeleteBody**](AttachmentsDeleteBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachmentsRedirectPost**
> AttachmentsRedirectPost(ctx, optional)
Retrieve an attachment

Load an attachment from where it is stored based on the id. If the attachment is private then a temporary, signed url with embedded credentials is generated on demand.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AttachmentsApiAttachmentsRedirectPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AttachmentsApiAttachmentsRedirectPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AttachmentsRedirectBody**](AttachmentsRedirectBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

