# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumentsArchivePost**](DocumentsApi.md#DocumentsArchivePost) | **Post** /documents.archive | Archive a document
[**DocumentsCreatePost**](DocumentsApi.md#DocumentsCreatePost) | **Post** /documents.create | Create a document
[**DocumentsDeletePost**](DocumentsApi.md#DocumentsDeletePost) | **Post** /documents.delete | Delete a document
[**DocumentsDraftsPost**](DocumentsApi.md#DocumentsDraftsPost) | **Post** /documents.drafts | List all draft documents
[**DocumentsExportPost**](DocumentsApi.md#DocumentsExportPost) | **Post** /documents.export | Export a document as markdown
[**DocumentsImportPost**](DocumentsApi.md#DocumentsImportPost) | **Post** /documents.import | Import a file as a document
[**DocumentsInfoPost**](DocumentsApi.md#DocumentsInfoPost) | **Post** /documents.info | Retrieve a document
[**DocumentsListPost**](DocumentsApi.md#DocumentsListPost) | **Post** /documents.list | List all documents
[**DocumentsMovePost**](DocumentsApi.md#DocumentsMovePost) | **Post** /documents.move | Move a document
[**DocumentsRestorePost**](DocumentsApi.md#DocumentsRestorePost) | **Post** /documents.restore | Restore a document
[**DocumentsSearchPost**](DocumentsApi.md#DocumentsSearchPost) | **Post** /documents.search | Search all documents
[**DocumentsStarPost**](DocumentsApi.md#DocumentsStarPost) | **Post** /documents.star | Star a document
[**DocumentsTemplatizePost**](DocumentsApi.md#DocumentsTemplatizePost) | **Post** /documents.templatize | Create a template from a document
[**DocumentsUnpublishPost**](DocumentsApi.md#DocumentsUnpublishPost) | **Post** /documents.unpublish | Unpublish a document
[**DocumentsUnstarPost**](DocumentsApi.md#DocumentsUnstarPost) | **Post** /documents.unstar | Unstar a document
[**DocumentsUpdatePost**](DocumentsApi.md#DocumentsUpdatePost) | **Post** /documents.update | Update a document
[**DocumentsViewedPost**](DocumentsApi.md#DocumentsViewedPost) | **Post** /documents.viewed | List all recently viewed documents

# **DocumentsArchivePost**
> InlineResponse20016 DocumentsArchivePost(ctx, optional)
Archive a document

Archiving a document allows outdated information to be moved out of sight whilst retaining the ability to optionally search and restore it later.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsArchivePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsArchivePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsArchiveBody**](DocumentsArchiveBody.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsCreatePost**
> InlineResponse20016 DocumentsCreatePost(ctx, optional)
Create a document

This method allows you to create or publish a new document. By default a document is set to the collection root. If you want to create a nested/child document, you should pass parentDocumentId to set the parent document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsCreateBody**](DocumentsCreateBody.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsDeletePost**
> InlineResponse2001 DocumentsDeletePost(ctx, optional)
Delete a document

Deleting a document moves it to the trash. If not restored within 30 days it is permenantly deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsDeleteBody**](DocumentsDeleteBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsDraftsPost**
> InlineResponse20014 DocumentsDraftsPost(ctx, optional)
List all draft documents

This method will list all draft documents belonging to the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsDraftsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsDraftsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsDraftsBody**](DocumentsDraftsBody.md)|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsExportPost**
> InlineResponse20013 DocumentsExportPost(ctx, optional)
Export a document as markdown

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsExportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsExportPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsExportBody**](DocumentsExportBody.md)|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsImportPost**
> InlineResponse20012 DocumentsImportPost(ctx, optional)
Import a file as a document

This method allows you to create a new document by importing an existing file. By default a document is set to the collection root. If you want to create a nested/child document, you should pass parentDocumentId to set the parent document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsImportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsImportPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | [**optional.Interface of interface{}**](.md)|  | 
 **collectionId** | [**optional.Interface of string**](.md)|  | 
 **parentDocumentId** | [**optional.Interface of string**](.md)|  | 
 **template** | **optional.**|  | 
 **publish** | **optional.**|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsInfoPost**
> InlineResponse20012 DocumentsInfoPost(ctx, optional)
Retrieve a document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsInfoPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsInfoPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsInfoBody**](DocumentsInfoBody.md)|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsListPost**
> InlineResponse20014 DocumentsListPost(ctx, optional)
List all documents

This method will list all published documents and draft documents belonging to the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsListBody**](DocumentsListBody.md)|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsMovePost**
> InlineResponse20018 DocumentsMovePost(ctx, optional)
Move a document

Move a document to a new location or collection. If no parent document is provided, the document will be moved to the collection root.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsMovePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsMovePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsMoveBody**](DocumentsMoveBody.md)|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsRestorePost**
> InlineResponse20016 DocumentsRestorePost(ctx, optional)
Restore a document

If a document has been archived or deleted, it can be restored. Optionally a revision can be passed to restore the document to a previous point in time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsRestorePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsRestorePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsRestoreBody**](DocumentsRestoreBody.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsSearchPost**
> InlineResponse20015 DocumentsSearchPost(ctx, optional)
Search all documents

This methods allows you to search your teams documents with keywords. Note that search results will be restricted to those accessible by the current access token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsSearchPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsSearchBody**](DocumentsSearchBody.md)|  | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsStarPost**
> InlineResponse20017 DocumentsStarPost(ctx, optional)
Star a document

Starring a document gives it extra priority in the UI and makes it easier to find important information later.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsStarPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsStarPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsStarBody**](DocumentsStarBody.md)|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsTemplatizePost**
> InlineResponse20016 DocumentsTemplatizePost(ctx, optional)
Create a template from a document

This method allows you to createa new template using an existing document as the basis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsTemplatizePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsTemplatizePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsTemplatizeBody**](DocumentsTemplatizeBody.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsUnpublishPost**
> InlineResponse20016 DocumentsUnpublishPost(ctx, optional)
Unpublish a document

Unpublishing a document moves it back to a draft status and out of the collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsUnpublishPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsUnpublishPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsUnpublishBody**](DocumentsUnpublishBody.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsUnstarPost**
> InlineResponse20017 DocumentsUnstarPost(ctx, optional)
Unstar a document

Starring a document gives it extra priority in the UI and makes it easier to find important information later.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsUnstarPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsUnstarPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsUnstarBody**](DocumentsUnstarBody.md)|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsUpdatePost**
> InlineResponse20016 DocumentsUpdatePost(ctx, optional)
Update a document

This method allows you to modify an already created document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsUpdatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsUpdatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsUpdateBody**](DocumentsUpdateBody.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocumentsViewedPost**
> InlineResponse20014 DocumentsViewedPost(ctx, optional)
List all recently viewed documents

This method will list all documents recently viewed by the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiDocumentsViewedPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiDocumentsViewedPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DocumentsViewedBody**](DocumentsViewedBody.md)|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

