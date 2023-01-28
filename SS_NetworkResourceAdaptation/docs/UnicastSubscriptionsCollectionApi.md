# \UnicastSubscriptionsCollectionApi

All URIs are relative to *https://example.com/ss-nra/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUnicastSubscription**](UnicastSubscriptionsCollectionApi.md#CreateUnicastSubscription) | **Post** /unicast-subscriptions | Creates a new Individual Unicast Subscription resource



## CreateUnicastSubscription

> UnicastSubscription CreateUnicastSubscription(ctx).UnicastSubscription(unicastSubscription).Execute()

Creates a new Individual Unicast Subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_NetworkResourceAdaptation"
)

func main() {
    unicastSubscription := *openapiclient.NewUnicastSubscription(openapiclient.ValTargetUe{Interface{}: new(interface{})}, "NotifUri_example") // UnicastSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnicastSubscriptionsCollectionApi.CreateUnicastSubscription(context.Background()).UnicastSubscription(unicastSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnicastSubscriptionsCollectionApi.CreateUnicastSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUnicastSubscription`: UnicastSubscription
    fmt.Fprintf(os.Stdout, "Response from `UnicastSubscriptionsCollectionApi.CreateUnicastSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUnicastSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unicastSubscription** | [**UnicastSubscription**](UnicastSubscription.md) |  | 

### Return type

[**UnicastSubscription**](UnicastSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

