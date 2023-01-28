# \TimeSynchronizationExposureSubscriptionsCollectionApi

All URIs are relative to *https://example.com/ntsctsf-time-sync/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TimeSynchronizationExposureSubscriptions**](TimeSynchronizationExposureSubscriptionsCollectionApi.md#TimeSynchronizationExposureSubscriptions) | **Post** /subscriptions | Creates a new subscription to notification of capability of time synchronization service resource



## TimeSynchronizationExposureSubscriptions

> TimeSyncExposureSubsc TimeSynchronizationExposureSubscriptions(ctx).TimeSyncExposureSubsc(timeSyncExposureSubsc).Execute()

Creates a new subscription to notification of capability of time synchronization service resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_TimeSynchronization"
)

func main() {
    timeSyncExposureSubsc := openapiclient.TimeSyncExposureSubsc{Interface{}: new(interface{})} // TimeSyncExposureSubsc | Contains the information for the creation the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimeSynchronizationExposureSubscriptionsCollectionApi.TimeSynchronizationExposureSubscriptions(context.Background()).TimeSyncExposureSubsc(timeSyncExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeSynchronizationExposureSubscriptionsCollectionApi.TimeSynchronizationExposureSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimeSynchronizationExposureSubscriptions`: TimeSyncExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `TimeSynchronizationExposureSubscriptionsCollectionApi.TimeSynchronizationExposureSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimeSynchronizationExposureSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeSyncExposureSubsc** | [**TimeSyncExposureSubsc**](TimeSyncExposureSubsc.md) | Contains the information for the creation the resource. | 

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

