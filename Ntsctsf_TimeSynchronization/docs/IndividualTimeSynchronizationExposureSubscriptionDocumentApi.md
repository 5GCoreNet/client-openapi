# \IndividualTimeSynchronizationExposureSubscriptionDocumentApi

All URIs are relative to *https://example.com/ntsctsf-time-sync/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualTimeSynchronizationExposureSubscription**](IndividualTimeSynchronizationExposureSubscriptionDocumentApi.md#DeleteIndividualTimeSynchronizationExposureSubscription) | **Delete** /subscriptions/{subscriptionId} | Delete an Individual TimeSynchronization Exposure Subscription
[**GetIndividualTimeSynchronizationExposureSubscription**](IndividualTimeSynchronizationExposureSubscriptionDocumentApi.md#GetIndividualTimeSynchronizationExposureSubscription) | **Get** /subscriptions/{subscriptionId} | Reads an existing Individual Time Synchronization Exposure Subscription
[**ReplaceIndividualTimeSynchronizationExposureSubscription**](IndividualTimeSynchronizationExposureSubscriptionDocumentApi.md#ReplaceIndividualTimeSynchronizationExposureSubscription) | **Put** /subscriptions/{subscriptionId} | Replace an individual Time Synchronization Exposure Subscription



## DeleteIndividualTimeSynchronizationExposureSubscription

> DeleteIndividualTimeSynchronizationExposureSubscription(ctx, subscriptionId).Execute()

Delete an Individual TimeSynchronization Exposure Subscription

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
    subscriptionId := "subscriptionId_example" // string | String identifying an Individual Time Synchronization Exposure Subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionDocumentApi.DeleteIndividualTimeSynchronizationExposureSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureSubscriptionDocumentApi.DeleteIndividualTimeSynchronizationExposureSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying an Individual Time Synchronization Exposure Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualTimeSynchronizationExposureSubscriptionRequest struct via the builder pattern


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


## GetIndividualTimeSynchronizationExposureSubscription

> TimeSyncExposureSubsc GetIndividualTimeSynchronizationExposureSubscription(ctx, subscriptionId).Execute()

Reads an existing Individual Time Synchronization Exposure Subscription

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
    subscriptionId := "subscriptionId_example" // string | String identifying an Individual Time Synchronization Exposure Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionDocumentApi.GetIndividualTimeSynchronizationExposureSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureSubscriptionDocumentApi.GetIndividualTimeSynchronizationExposureSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndividualTimeSynchronizationExposureSubscription`: TimeSyncExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualTimeSynchronizationExposureSubscriptionDocumentApi.GetIndividualTimeSynchronizationExposureSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying an Individual Time Synchronization Exposure Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualTimeSynchronizationExposureSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimeSyncExposureSubsc**](TimeSyncExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceIndividualTimeSynchronizationExposureSubscription

> TimeSyncExposureSubsc ReplaceIndividualTimeSynchronizationExposureSubscription(ctx, subscriptionId).TimeSyncExposureSubsc1(timeSyncExposureSubsc1).Execute()

Replace an individual Time Synchronization Exposure Subscription

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
    subscriptionId := "subscriptionId_example" // string | String identifying an Individual Time Synchronization Exposure Subscription.
    timeSyncExposureSubsc1 := *openapiclient.NewTimeSyncExposureSubsc1("SubsNotifId_example", "SubsNotifUri_example") // TimeSyncExposureSubsc1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionDocumentApi.ReplaceIndividualTimeSynchronizationExposureSubscription(context.Background(), subscriptionId).TimeSyncExposureSubsc1(timeSyncExposureSubsc1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureSubscriptionDocumentApi.ReplaceIndividualTimeSynchronizationExposureSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIndividualTimeSynchronizationExposureSubscription`: TimeSyncExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualTimeSynchronizationExposureSubscriptionDocumentApi.ReplaceIndividualTimeSynchronizationExposureSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying an Individual Time Synchronization Exposure Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIndividualTimeSynchronizationExposureSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeSyncExposureSubsc1** | [**TimeSyncExposureSubsc1**](TimeSyncExposureSubsc1.md) |  | 

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

