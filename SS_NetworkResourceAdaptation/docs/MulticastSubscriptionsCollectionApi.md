# \MulticastSubscriptionsCollectionApi

All URIs are relative to *https://example.com/ss-nra/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMulticastSubscription**](MulticastSubscriptionsCollectionApi.md#CreateMulticastSubscription) | **Post** /multicast-subscriptions | Creates a new Individual Multicast Subscription resource



## CreateMulticastSubscription

> MulticastSubscription CreateMulticastSubscription(ctx).MulticastSubscription(multicastSubscription).Execute()

Creates a new Individual Multicast Subscription resource

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
    multicastSubscription := *openapiclient.NewMulticastSubscription("ValGroupId_example", *openapiclient.NewServiceAnnoucementMode(), "MultiQosReq_example", "NotifUri_example") // MulticastSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MulticastSubscriptionsCollectionApi.CreateMulticastSubscription(context.Background()).MulticastSubscription(multicastSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MulticastSubscriptionsCollectionApi.CreateMulticastSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMulticastSubscription`: MulticastSubscription
    fmt.Fprintf(os.Stdout, "Response from `MulticastSubscriptionsCollectionApi.CreateMulticastSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMulticastSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multicastSubscription** | [**MulticastSubscription**](MulticastSubscription.md) |  | 

### Return type

[**MulticastSubscription**](MulticastSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

