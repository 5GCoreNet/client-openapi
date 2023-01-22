# \RealTimeUAVStatusSubscriptionsCollectionApi

All URIs are relative to *https://example.com/uae-uav-status/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRTUavStatusSubsc**](RealTimeUAVStatusSubscriptionsCollectionApi.md#CreateRTUavStatusSubsc) | **Post** /subscriptions | Request the creation of a subscription to real-time UAV status reporting.
[**GetRTUavStatusSubscriptions**](RealTimeUAVStatusSubscriptionsCollectionApi.md#GetRTUavStatusSubscriptions) | **Get** /subscriptions | Retrieve all the active real-time UAV status subscriptions managed by the UAE Server.



## CreateRTUavStatusSubsc

> RTUavStatusSubsc CreateRTUavStatusSubsc(ctx).RTUavStatusSubsc(rTUavStatusSubsc).Execute()

Request the creation of a subscription to real-time UAV status reporting.

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
    rTUavStatusSubsc := *openapiclient.NewRTUavStatusSubsc("UassId_example", []openapiclient.UavId{*openapiclient.NewUavId()}, "NotificationUri_example") // RTUavStatusSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeUAVStatusSubscriptionsCollectionApi.CreateRTUavStatusSubsc(context.Background()).RTUavStatusSubsc(rTUavStatusSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeUAVStatusSubscriptionsCollectionApi.CreateRTUavStatusSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRTUavStatusSubsc`: RTUavStatusSubsc
    fmt.Fprintf(os.Stdout, "Response from `RealTimeUAVStatusSubscriptionsCollectionApi.CreateRTUavStatusSubsc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRTUavStatusSubscRequest struct via the builder pattern


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


## GetRTUavStatusSubscriptions

> []RTUavStatusSubsc GetRTUavStatusSubscriptions(ctx).Execute()

Retrieve all the active real-time UAV status subscriptions managed by the UAE Server.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeUAVStatusSubscriptionsCollectionApi.GetRTUavStatusSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeUAVStatusSubscriptionsCollectionApi.GetRTUavStatusSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRTUavStatusSubscriptions`: []RTUavStatusSubsc
    fmt.Fprintf(os.Stdout, "Response from `RealTimeUAVStatusSubscriptionsCollectionApi.GetRTUavStatusSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRTUavStatusSubscriptionsRequest struct via the builder pattern


### Return type

[**[]RTUavStatusSubsc**](RTUavStatusSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

