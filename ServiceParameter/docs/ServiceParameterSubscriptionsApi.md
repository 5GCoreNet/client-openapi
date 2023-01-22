# \ServiceParameterSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-service-parameter/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnSubscription**](ServiceParameterSubscriptionsApi.md#CreateAnSubscription) | **Post** /{afId}/subscriptions | Creates a new subscription resource
[**ReadAllSubscriptions**](ServiceParameterSubscriptionsApi.md#ReadAllSubscriptions) | **Get** /{afId}/subscriptions | read all of the active subscriptions for the AF



## CreateAnSubscription

> ServiceParameterData CreateAnSubscription(ctx, afId).ServiceParameterData(serviceParameterData).Execute()

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
    serviceParameterData := *openapiclient.NewServiceParameterData() // ServiceParameterData | Request to create a new subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceParameterSubscriptionsApi.CreateAnSubscription(context.Background(), afId).ServiceParameterData(serviceParameterData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceParameterSubscriptionsApi.CreateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnSubscription`: ServiceParameterData
    fmt.Fprintf(os.Stdout, "Response from `ServiceParameterSubscriptionsApi.CreateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceParameterData** | [**ServiceParameterData**](ServiceParameterData.md) | Request to create a new subscription resource | 

### Return type

[**ServiceParameterData**](ServiceParameterData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllSubscriptions

> []ServiceParameterData ReadAllSubscriptions(ctx, afId).Gpsis(gpsis).IpAddrs(ipAddrs).IpDomain(ipDomain).MacAddrs(macAddrs).Execute()

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
    gpsis := []string{"Inner_example"} // []string | The GPSI of the requested UE(s). (optional)
    ipAddrs := []openapiclient.IpAddr{openapiclient.IpAddr{Interface{}: new(interface{})}} // []IpAddr | The IP address(es) of the requested UE(s). (optional)
    ipDomain := "ipDomain_example" // string | The IPv4 address domain identifier. The attribute may only be provided if IPv4 address is included in the ip-addrs query parameter.  (optional)
    macAddrs := []string{"Inner_example"} // []string | The MAC address(es) of the requested UE(s). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceParameterSubscriptionsApi.ReadAllSubscriptions(context.Background(), afId).Gpsis(gpsis).IpAddrs(ipAddrs).IpDomain(ipDomain).MacAddrs(macAddrs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceParameterSubscriptionsApi.ReadAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSubscriptions`: []ServiceParameterData
    fmt.Fprintf(os.Stdout, "Response from `ServiceParameterSubscriptionsApi.ReadAllSubscriptions`: %v\n", resp)
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

 **gpsis** | **[]string** | The GPSI of the requested UE(s). | 
 **ipAddrs** | [**[]IpAddr**](IpAddr.md) | The IP address(es) of the requested UE(s). | 
 **ipDomain** | **string** | The IPv4 address domain identifier. The attribute may only be provided if IPv4 address is included in the ip-addrs query parameter.  | 
 **macAddrs** | **[]string** | The MAC address(es) of the requested UE(s). | 

### Return type

[**[]ServiceParameterData**](ServiceParameterData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

