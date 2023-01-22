# \SubscriptionsCollectionApi

All URIs are relative to *https://example.com/nnssf-nssaiavailability/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NSSAIAvailabilityPost**](SubscriptionsCollectionApi.md#NSSAIAvailabilityPost) | **Post** /nssai-availability/subscriptions | Creates subscriptions for notification about updates to NSSAI availability information



## NSSAIAvailabilityPost

> NssfEventSubscriptionCreatedData NSSAIAvailabilityPost(ctx).NssfEventSubscriptionCreateData(nssfEventSubscriptionCreateData).ContentEncoding(contentEncoding).Execute()

Creates subscriptions for notification about updates to NSSAI availability information

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
    nssfEventSubscriptionCreateData := *openapiclient.NewNssfEventSubscriptionCreateData("NfNssaiAvailabilityUri_example", []openapiclient.Tai{*openapiclient.NewTai(*openapiclient.NewPlmnId("Mcc_example", "Mnc_example"), "Tac_example")}, *openapiclient.NewNssfEventType()) // NssfEventSubscriptionCreateData | Subscription for notification about updates to NSSAI availability information
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionApi.NSSAIAvailabilityPost(context.Background()).NssfEventSubscriptionCreateData(nssfEventSubscriptionCreateData).ContentEncoding(contentEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionApi.NSSAIAvailabilityPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NSSAIAvailabilityPost`: NssfEventSubscriptionCreatedData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionApi.NSSAIAvailabilityPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNSSAIAvailabilityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nssfEventSubscriptionCreateData** | [**NssfEventSubscriptionCreateData**](NssfEventSubscriptionCreateData.md) | Subscription for notification about updates to NSSAI availability information | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 

### Return type

[**NssfEventSubscriptionCreatedData**](NssfEventSubscriptionCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

