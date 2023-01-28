# \SubscriptionsCollectionApi

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsCollectionApi.md#CreateSubscription) | **Post** /subscriptions | Create a new subscription



## CreateSubscription

> SubscriptionData CreateSubscription(ctx).SubscriptionData(subscriptionData).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()

Create a new subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnrf_NFManagement"
)

func main() {
    subscriptionData := *openapiclient.NewSubscriptionData("NfStatusNotificationUri_example", "SubscriptionId_example") // SubscriptionData | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionApi.CreateSubscription(context.Background()).SubscriptionData(subscriptionData).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: SubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionData** | [**SubscriptionData**](SubscriptionData.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 

### Return type

[**SubscriptionData**](SubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

