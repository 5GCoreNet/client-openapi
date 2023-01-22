# \EventsSubscriptionDocumentApi

All URIs are relative to *https://example.com/ntsctsf-qos-tscai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEventsSubsc**](EventsSubscriptionDocumentApi.md#DeleteEventsSubsc) | **Delete** /tsc-app-sessions/{appSessionId}/events-subscription | Deletes the Events Subscription subresource.
[**PutEventsSubsc**](EventsSubscriptionDocumentApi.md#PutEventsSubsc) | **Put** /tsc-app-sessions/{appSessionId}/events-subscription | Creates or modifies an Events Subscription subresource



## DeleteEventsSubsc

> DeleteEventsSubsc(ctx, appSessionId).Execute()

Deletes the Events Subscription subresource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appSessionId := "appSessionId_example" // string | String identifying the Individual TSC Application Session Context resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsSubscriptionDocumentApi.DeleteEventsSubsc(context.Background(), appSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsSubscriptionDocumentApi.DeleteEventsSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSessionId** | **string** | String identifying the Individual TSC Application Session Context resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventsSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEventsSubsc

> EventsSubscReqData PutEventsSubsc(ctx, appSessionId).EventsSubscReqData(eventsSubscReqData).Execute()

Creates or modifies an Events Subscription subresource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appSessionId := "appSessionId_example" // string | String identifying the Events Subscription resource
    eventsSubscReqData := *openapiclient.NewEventsSubscReqData([]openapiclient.TscEvent{*openapiclient.NewTscEvent()}, "NotifUri_example", "NotifCorreId_example") // EventsSubscReqData | Creation or modification of an Events Subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsSubscriptionDocumentApi.PutEventsSubsc(context.Background(), appSessionId).EventsSubscReqData(eventsSubscReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsSubscriptionDocumentApi.PutEventsSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutEventsSubsc`: EventsSubscReqData
    fmt.Fprintf(os.Stdout, "Response from `EventsSubscriptionDocumentApi.PutEventsSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSessionId** | **string** | String identifying the Events Subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEventsSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventsSubscReqData** | [**EventsSubscReqData**](EventsSubscReqData.md) | Creation or modification of an Events Subscription resource. | 

### Return type

[**EventsSubscReqData**](EventsSubscReqData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

