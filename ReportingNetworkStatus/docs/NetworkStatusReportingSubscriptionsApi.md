# \NetworkStatusReportingSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-net-stat-report/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNwStatusReportSubscription**](NetworkStatusReportingSubscriptionsApi.md#CreateNwStatusReportSubscription) | **Post** /{scsAsId}/subscriptions | Create a new network status reporting subscription resource.
[**FetchAllNwStatusReportSubscriptions**](NetworkStatusReportingSubscriptionsApi.md#FetchAllNwStatusReportSubscriptions) | **Get** /{scsAsId}/subscriptions | Read all network status reporting subscription resources for a given SCS/AS.



## CreateNwStatusReportSubscription

> NetworkStatusReportingSubscription CreateNwStatusReportSubscription(ctx, scsAsId).NetworkStatusReportingSubscription(networkStatusReportingSubscription).Execute()

Create a new network status reporting subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ReportingNetworkStatus"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    networkStatusReportingSubscription := *openapiclient.NewNetworkStatusReportingSubscription("NotificationDestination_example", *openapiclient.NewLocationArea()) // NetworkStatusReportingSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkStatusReportingSubscriptionsApi.CreateNwStatusReportSubscription(context.Background(), scsAsId).NetworkStatusReportingSubscription(networkStatusReportingSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkStatusReportingSubscriptionsApi.CreateNwStatusReportSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNwStatusReportSubscription`: NetworkStatusReportingSubscription
    fmt.Fprintf(os.Stdout, "Response from `NetworkStatusReportingSubscriptionsApi.CreateNwStatusReportSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNwStatusReportSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkStatusReportingSubscription** | [**NetworkStatusReportingSubscription**](NetworkStatusReportingSubscription.md) |  | 

### Return type

[**NetworkStatusReportingSubscription**](NetworkStatusReportingSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllNwStatusReportSubscriptions

> []NetworkStatusReportingSubscription FetchAllNwStatusReportSubscriptions(ctx, scsAsId).Execute()

Read all network status reporting subscription resources for a given SCS/AS.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ReportingNetworkStatus"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkStatusReportingSubscriptionsApi.FetchAllNwStatusReportSubscriptions(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkStatusReportingSubscriptionsApi.FetchAllNwStatusReportSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllNwStatusReportSubscriptions`: []NetworkStatusReportingSubscription
    fmt.Fprintf(os.Stdout, "Response from `NetworkStatusReportingSubscriptionsApi.FetchAllNwStatusReportSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllNwStatusReportSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NetworkStatusReportingSubscription**](NetworkStatusReportingSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

