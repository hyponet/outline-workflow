# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileOperationsInfoPost**](FileOperationsApi.md#FileOperationsInfoPost) | **Post** /fileOperations.info | Retrieve a file operation
[**FileOperationsListPost**](FileOperationsApi.md#FileOperationsListPost) | **Post** /fileOperations.list | List all file operations
[**FileOperationsRedirectPost**](FileOperationsApi.md#FileOperationsRedirectPost) | **Post** /fileOperations.redirect | Retrieve the file

# **FileOperationsInfoPost**
> InlineResponse20020 FileOperationsInfoPost(ctx, optional)
Retrieve a file operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileOperationsApiFileOperationsInfoPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileOperationsApiFileOperationsInfoPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FileOperationsInfoBody**](FileOperationsInfoBody.md)|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileOperationsListPost**
> InlineResponse20021 FileOperationsListPost(ctx, optional)
List all file operations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileOperationsApiFileOperationsListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileOperationsApiFileOperationsListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FileOperationsListBody**](FileOperationsListBody.md)|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileOperationsRedirectPost**
> *os.File FileOperationsRedirectPost(ctx, optional)
Retrieve the file

Load the resulting file from where it is stored based on the id. A temporary, signed url with embedded credentials is generated on demand.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileOperationsApiFileOperationsRedirectPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileOperationsApiFileOperationsRedirectPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FileOperationsRedirectBody**](FileOperationsRedirectBody.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

