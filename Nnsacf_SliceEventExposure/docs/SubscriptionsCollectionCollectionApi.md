# \SubscriptionsCollectionCollectionApi

All URIs are relative to *https://example.com/nnsacf-slice-ee/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsCollectionCollectionApi.md#CreateSubscription) | **Post** /subscriptions | Nnsacf_SliceEventExposure Subscribe service Operation



## CreateSubscription

> CreatedSACEventSubscription CreateSubscription(ctx).SACEventSubscription(sACEventSubscription).Execute()

Nnsacf_SliceEventExposure Subscribe service Operation

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
    sACEventSubscription := *openapiclient.NewSACEventSubscription(*openapiclient.NewSACEvent(*openapiclient.NewSACEventType(), []openapiclient.Snssai{*openapiclient.NewSnssai(int32(123))}), "EventNotifyUri_example", "NfId_example") // SACEventSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionCollectionApi.CreateSubscription(context.Background()).SACEventSubscription(sACEventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionCollectionApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: CreatedSACEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionCollectionApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sACEventSubscription** | [**SACEventSubscription**](SACEventSubscription.md) |  | 

### Return type

[**CreatedSACEventSubscription**](CreatedSACEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

