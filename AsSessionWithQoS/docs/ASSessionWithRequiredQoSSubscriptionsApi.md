# \ASSessionWithRequiredQoSSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-as-session-with-qos/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateASSessionWithQoSSubscription**](ASSessionWithRequiredQoSSubscriptionsApi.md#CreateASSessionWithQoSSubscription) | **Post** /{scsAsId}/subscriptions | Creates a new subscription resource.
[**FetchAllASSessionWithQoSSubscriptions**](ASSessionWithRequiredQoSSubscriptionsApi.md#FetchAllASSessionWithQoSSubscriptions) | **Get** /{scsAsId}/subscriptions | Read all or queried active subscriptions for the SCS/AS.



## CreateASSessionWithQoSSubscription

> AsSessionWithQoSSubscription CreateASSessionWithQoSSubscription(ctx, scsAsId).AsSessionWithQoSSubscription(asSessionWithQoSSubscription).Execute()

Creates a new subscription resource.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    asSessionWithQoSSubscription := *openapiclient.NewAsSessionWithQoSSubscription("NotificationDestination_example") // AsSessionWithQoSSubscription | Request to create a new subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASSessionWithRequiredQoSSubscriptionsApi.CreateASSessionWithQoSSubscription(context.Background(), scsAsId).AsSessionWithQoSSubscription(asSessionWithQoSSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASSessionWithRequiredQoSSubscriptionsApi.CreateASSessionWithQoSSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateASSessionWithQoSSubscription`: AsSessionWithQoSSubscription
    fmt.Fprintf(os.Stdout, "Response from `ASSessionWithRequiredQoSSubscriptionsApi.CreateASSessionWithQoSSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateASSessionWithQoSSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asSessionWithQoSSubscription** | [**AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md) | Request to create a new subscription resource | 

### Return type

[**AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllASSessionWithQoSSubscriptions

> []AsSessionWithQoSSubscription FetchAllASSessionWithQoSSubscriptions(ctx, scsAsId).IpAddrs(ipAddrs).IpDomain(ipDomain).MacAddrs(macAddrs).Execute()

Read all or queried active subscriptions for the SCS/AS.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    ipAddrs := []openapiclient.IpAddr{openapiclient.IpAddr{Interface{}: new(interface{})}} // []IpAddr | The IP address(es) of the requested UE(s). (optional)
    ipDomain := "ipDomain_example" // string | The IPv4 address domain identifier. The attribute may only be provided if IPv4 address is included in the ip-addrs query parameter. (optional)
    macAddrs := []string{"Inner_example"} // []string | The MAC address(es) of the requested UE(s). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASSessionWithRequiredQoSSubscriptionsApi.FetchAllASSessionWithQoSSubscriptions(context.Background(), scsAsId).IpAddrs(ipAddrs).IpDomain(ipDomain).MacAddrs(macAddrs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASSessionWithRequiredQoSSubscriptionsApi.FetchAllASSessionWithQoSSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllASSessionWithQoSSubscriptions`: []AsSessionWithQoSSubscription
    fmt.Fprintf(os.Stdout, "Response from `ASSessionWithRequiredQoSSubscriptionsApi.FetchAllASSessionWithQoSSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllASSessionWithQoSSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipAddrs** | [**[]IpAddr**](IpAddr.md) | The IP address(es) of the requested UE(s). | 
 **ipDomain** | **string** | The IPv4 address domain identifier. The attribute may only be provided if IPv4 address is included in the ip-addrs query parameter. | 
 **macAddrs** | **[]string** | The MAC address(es) of the requested UE(s). | 

### Return type

[**[]AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

