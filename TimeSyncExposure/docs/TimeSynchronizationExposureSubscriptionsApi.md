# \TimeSynchronizationExposureSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-time-sync/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewSubscription**](TimeSynchronizationExposureSubscriptionsApi.md#CreateNewSubscription) | **Post** /{afId}/subscriptions | Creates a new subscription resource
[**ReadAllSubscriptions**](TimeSynchronizationExposureSubscriptionsApi.md#ReadAllSubscriptions) | **Get** /{afId}/subscriptions | read all of the active subscriptions for the AF



## CreateNewSubscription

> TimeSyncExposureSubsc CreateNewSubscription(ctx, afId).TimeSyncExposureSubsc(timeSyncExposureSubsc).Execute()

Creates a new subscription resource

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
    afId := "afId_example" // string | Identifier of the AF
    timeSyncExposureSubsc := *openapiclient.NewTimeSyncExposureSubsc("SubsNotifId_example", "SubsNotifUri_example") // TimeSyncExposureSubsc | new subscription creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimeSynchronizationExposureSubscriptionsApi.CreateNewSubscription(context.Background(), afId).TimeSyncExposureSubsc(timeSyncExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeSynchronizationExposureSubscriptionsApi.CreateNewSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewSubscription`: TimeSyncExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `TimeSynchronizationExposureSubscriptionsApi.CreateNewSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeSyncExposureSubsc** | [**TimeSyncExposureSubsc**](TimeSyncExposureSubsc.md) | new subscription creation | 

### Return type

[**TimeSyncExposureSubsc**](TimeSyncExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllSubscriptions

> []TimeSyncExposureSubsc ReadAllSubscriptions(ctx, afId).Execute()

read all of the active subscriptions for the AF

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
    afId := "afId_example" // string | Identifier of the AF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimeSynchronizationExposureSubscriptionsApi.ReadAllSubscriptions(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeSynchronizationExposureSubscriptionsApi.ReadAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSubscriptions`: []TimeSyncExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `TimeSynchronizationExposureSubscriptionsApi.ReadAllSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TimeSyncExposureSubsc**](TimeSyncExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

