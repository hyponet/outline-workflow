# {{classname}}

All URIs are relative to *https://app.getoutline.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsListPost**](EventsApi.md#EventsListPost) | **Post** /events.list | List all events

# **EventsListPost**
> InlineResponse20019 EventsListPost(ctx, optional)
List all events

Events are an audit trail of important events that happen in the knowledge base.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventsApiEventsListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiEventsListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EventsListBody**](EventsListBody.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[http](../README.md#http)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

