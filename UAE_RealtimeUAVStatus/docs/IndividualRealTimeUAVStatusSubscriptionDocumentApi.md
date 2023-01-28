# \IndividualRealTimeUAVStatusSubscriptionDocumentApi

All URIs are relative to *https://example.com/uae-uav-status/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRTUavStatusSubscription**](IndividualRealTimeUAVStatusSubscriptionDocumentApi.md#DeleteRTUavStatusSubscription) | **Delete** /subscriptions/{subscriptionId} | Request the deletion of an existing real-time UAV status subscription.
[**GetRTUavStatusSubscription**](IndividualRealTimeUAVStatusSubscriptionDocumentApi.md#GetRTUavStatusSubscription) | **Get** /subscriptions/{subscriptionId} | Retrieve a real-time UAV status subscription resource.
[**UpdateRTUavStatusSubscription**](IndividualRealTimeUAVStatusSubscriptionDocumentApi.md#UpdateRTUavStatusSubscription) | **Put** /subscriptions/{subscriptionId} | Request the update of an existing real-time UAV status subscription.



## DeleteRTUavStatusSubscription

> DeleteRTUavStatusSubscription(ctx, subscriptionId).Execute()

Request the deletion of an existing real-time UAV status subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/UAE_RealtimeUAVStatus"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Individual Real-time UAV Status Subscription identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualRealTimeUAVStatusSubscriptionDocumentApi.DeleteRTUavStatusSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualRealTimeUAVStatusSubscriptionDocumentApi.DeleteRTUavStatusSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Individual Real-time UAV Status Subscription identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRTUavStatusSubscriptionRequest struct via the builder pattern


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


## GetRTUavStatusSubscription

> RTUavStatusSubsc GetRTUavStatusSubscription(ctx, subscriptionId).Execute()

Retrieve a real-time UAV status subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/UAE_RealtimeUAVStatus"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Individual Real-time UAV Status Subscription identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualRealTimeUAVStatusSubscriptionDocumentApi.GetRTUavStatusSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualRealTimeUAVStatusSubscriptionDocumentApi.GetRTUavStatusSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRTUavStatusSubscription`: RTUavStatusSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualRealTimeUAVStatusSubscriptionDocumentApi.GetRTUavStatusSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Individual Real-time UAV Status Subscription identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRTUavStatusSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RTUavStatusSubsc**](RTUavStatusSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRTUavStatusSubscription

> RTUavStatusSubsc UpdateRTUavStatusSubscription(ctx, subscriptionId).RTUavStatusSubsc(rTUavStatusSubsc).Execute()

Request the update of an existing real-time UAV status subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/UAE_RealtimeUAVStatus"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Individual Real-time UAV Status Subscription identifier.
    rTUavStatusSubsc := *openapiclient.NewRTUavStatusSubsc("UassId_example", []openapiclient.UavId{*openapiclient.NewUavId()}, "NotificationUri_example") // RTUavStatusSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualRealTimeUAVStatusSubscriptionDocumentApi.UpdateRTUavStatusSubscription(context.Background(), subscriptionId).RTUavStatusSubsc(rTUavStatusSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualRealTimeUAVStatusSubscriptionDocumentApi.UpdateRTUavStatusSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRTUavStatusSubscription`: RTUavStatusSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualRealTimeUAVStatusSubscriptionDocumentApi.UpdateRTUavStatusSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Individual Real-time UAV Status Subscription identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRTUavStatusSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rTUavStatusSubsc** | [**RTUavStatusSubsc**](RTUavStatusSubsc.md) |  | 

### Return type

[**RTUavStatusSubsc**](RTUavStatusSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

