# \IndividualTimeSynchronizationExposureSubscriptionApi

All URIs are relative to *https://example.com/3gpp-time-sync/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnSubscription**](IndividualTimeSynchronizationExposureSubscriptionApi.md#DeleteAnSubscription) | **Delete** /{afId}/subscriptions/{subscriptionId} | Deletes an already existing subscription
[**FullyUpdateAnSubscription**](IndividualTimeSynchronizationExposureSubscriptionApi.md#FullyUpdateAnSubscription) | **Put** /{afId}/subscriptions/{subscriptionId} | Fully updates/replaces an existing subscription resource
[**ReadAnSubscription**](IndividualTimeSynchronizationExposureSubscriptionApi.md#ReadAnSubscription) | **Get** /{afId}/subscriptions/{subscriptionId} | read an active subscription for the AF and the subscription Id
[**ReadTimeSynSubscription**](IndividualTimeSynchronizationExposureSubscriptionApi.md#ReadTimeSynSubscription) | **Get** /{afId}/subscriptions/{subscriptionId}/configurations/{instanceReference} | read an active subscription for the AF and the subscription Id



## DeleteAnSubscription

> DeleteAnSubscription(ctx, afId, subscriptionId).Execute()

Deletes an already existing subscription

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionApi.DeleteAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureSubscriptionApi.DeleteAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FullyUpdateAnSubscription

> TimeSyncExposureSubsc FullyUpdateAnSubscription(ctx, afId, subscriptionId).TimeSyncExposureSubsc(timeSyncExposureSubsc).Execute()

Fully updates/replaces an existing subscription resource

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    timeSyncExposureSubsc := *openapiclient.NewTimeSyncExposureSubsc("SubsNotifId_example", "SubsNotifUri_example") // TimeSyncExposureSubsc | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionApi.FullyUpdateAnSubscription(context.Background(), afId, subscriptionId).TimeSyncExposureSubsc(timeSyncExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureSubscriptionApi.FullyUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnSubscription`: TimeSyncExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualTimeSynchronizationExposureSubscriptionApi.FullyUpdateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeSyncExposureSubsc** | [**TimeSyncExposureSubsc**](TimeSyncExposureSubsc.md) | Parameters to update/replace the existing subscription | 

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


## ReadAnSubscription

> TimeSyncExposureSubsc ReadAnSubscription(ctx, afId, subscriptionId).Execute()

read an active subscription for the AF and the subscription Id

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionApi.ReadAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureSubscriptionApi.ReadAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnSubscription`: TimeSyncExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualTimeSynchronizationExposureSubscriptionApi.ReadAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnSubscriptionRequest struct via the builder pattern


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


## ReadTimeSynSubscription

> TimeSyncExposureConfig ReadTimeSynSubscription(ctx, afId, subscriptionId, instanceReference).Execute()

read an active subscription for the AF and the subscription Id

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    instanceReference := "instanceReference_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionApi.ReadTimeSynSubscription(context.Background(), afId, subscriptionId, instanceReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureSubscriptionApi.ReadTimeSynSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTimeSynSubscription`: TimeSyncExposureConfig
    fmt.Fprintf(os.Stdout, "Response from `IndividualTimeSynchronizationExposureSubscriptionApi.ReadTimeSynSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 
**instanceReference** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTimeSynSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TimeSyncExposureConfig**](TimeSyncExposureConfig.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

