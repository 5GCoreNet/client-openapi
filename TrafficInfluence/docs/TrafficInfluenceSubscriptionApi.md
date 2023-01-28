# \TrafficInfluenceSubscriptionApi

All URIs are relative to *https://example.com/3gpp-traffic-influence/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewSubscription**](TrafficInfluenceSubscriptionApi.md#CreateNewSubscription) | **Post** /{afId}/subscriptions | Creates a new subscription resource
[**ReadAllSubscriptions**](TrafficInfluenceSubscriptionApi.md#ReadAllSubscriptions) | **Get** /{afId}/subscriptions | read all of the active subscriptions for the AF



## CreateNewSubscription

> TrafficInfluSub CreateNewSubscription(ctx, afId).TrafficInfluSub(trafficInfluSub).Execute()

Creates a new subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/TrafficInfluence"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    trafficInfluSub := *openapiclient.NewTrafficInfluSub() // TrafficInfluSub | Request to create a new subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrafficInfluenceSubscriptionApi.CreateNewSubscription(context.Background(), afId).TrafficInfluSub(trafficInfluSub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficInfluenceSubscriptionApi.CreateNewSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewSubscription`: TrafficInfluSub
    fmt.Fprintf(os.Stdout, "Response from `TrafficInfluenceSubscriptionApi.CreateNewSubscription`: %v\n", resp)
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

 **trafficInfluSub** | [**TrafficInfluSub**](TrafficInfluSub.md) | Request to create a new subscription resource | 

### Return type

[**TrafficInfluSub**](TrafficInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllSubscriptions

> []TrafficInfluSub ReadAllSubscriptions(ctx, afId).Execute()

read all of the active subscriptions for the AF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/TrafficInfluence"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrafficInfluenceSubscriptionApi.ReadAllSubscriptions(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficInfluenceSubscriptionApi.ReadAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSubscriptions`: []TrafficInfluSub
    fmt.Fprintf(os.Stdout, "Response from `TrafficInfluenceSubscriptionApi.ReadAllSubscriptions`: %v\n", resp)
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

[**[]TrafficInfluSub**](TrafficInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

