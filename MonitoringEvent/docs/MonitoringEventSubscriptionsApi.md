# \MonitoringEventSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-monitoring-event/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMonitoringEventSubscription**](MonitoringEventSubscriptionsApi.md#CreateMonitoringEventSubscription) | **Post** /{scsAsId}/subscriptions | Creates a new subscription resource for monitoring event notification.
[**FetchAllMonitoringEventSubscriptions**](MonitoringEventSubscriptionsApi.md#FetchAllMonitoringEventSubscriptions) | **Get** /{scsAsId}/subscriptions | Read all or queried active subscriptions for the SCS/AS.



## CreateMonitoringEventSubscription

> CreateMonitoringEventSubscription200Response CreateMonitoringEventSubscription(ctx, scsAsId).MonitoringEventSubscription(monitoringEventSubscription).Execute()

Creates a new subscription resource for monitoring event notification.

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
    monitoringEventSubscription := *openapiclient.NewMonitoringEventSubscription("NotificationDestination_example", *openapiclient.NewMonitoringType()) // MonitoringEventSubscription | Subscription for notification about monitoring event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringEventSubscriptionsApi.CreateMonitoringEventSubscription(context.Background(), scsAsId).MonitoringEventSubscription(monitoringEventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringEventSubscriptionsApi.CreateMonitoringEventSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitoringEventSubscription`: CreateMonitoringEventSubscription200Response
    fmt.Fprintf(os.Stdout, "Response from `MonitoringEventSubscriptionsApi.CreateMonitoringEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitoringEventSubscription** | [**MonitoringEventSubscription**](MonitoringEventSubscription.md) | Subscription for notification about monitoring event | 

### Return type

[**CreateMonitoringEventSubscription200Response**](CreateMonitoringEventSubscription200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllMonitoringEventSubscriptions

> []MonitoringEventSubscription FetchAllMonitoringEventSubscriptions(ctx, scsAsId).IpAddrs(ipAddrs).IpDomain(ipDomain).MacAddrs(macAddrs).Execute()

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
    resp, r, err := apiClient.MonitoringEventSubscriptionsApi.FetchAllMonitoringEventSubscriptions(context.Background(), scsAsId).IpAddrs(ipAddrs).IpDomain(ipDomain).MacAddrs(macAddrs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringEventSubscriptionsApi.FetchAllMonitoringEventSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllMonitoringEventSubscriptions`: []MonitoringEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `MonitoringEventSubscriptionsApi.FetchAllMonitoringEventSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllMonitoringEventSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipAddrs** | [**[]IpAddr**](IpAddr.md) | The IP address(es) of the requested UE(s). | 
 **ipDomain** | **string** | The IPv4 address domain identifier. The attribute may only be provided if IPv4 address is included in the ip-addrs query parameter. | 
 **macAddrs** | **[]string** | The MAC address(es) of the requested UE(s). | 

### Return type

[**[]MonitoringEventSubscription**](MonitoringEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

