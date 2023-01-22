# \DCCFAnalyticsSubscriptionsCollectionApi

All URIs are relative to *https://example.com/ndccf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDCCFAnalyticsSubscription**](DCCFAnalyticsSubscriptionsCollectionApi.md#CreateDCCFAnalyticsSubscription) | **Post** /analytics-subscriptions | Creates a new Individual DCCF Analytics Subscription resource.



## CreateDCCFAnalyticsSubscription

> NdccfAnalyticsSubscription CreateDCCFAnalyticsSubscription(ctx).NdccfAnalyticsSubscription(ndccfAnalyticsSubscription).Execute()

Creates a new Individual DCCF Analytics Subscription resource.

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
    ndccfAnalyticsSubscription := *openapiclient.NewNdccfAnalyticsSubscription(*openapiclient.NewNnwdafEventsSubscription([]openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewNwdafEvent())}), "AnaNotifUri_example", "AnaNotifCorrId_example") // NdccfAnalyticsSubscription | Contains the information for the creation the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DCCFAnalyticsSubscriptionsCollectionApi.CreateDCCFAnalyticsSubscription(context.Background()).NdccfAnalyticsSubscription(ndccfAnalyticsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DCCFAnalyticsSubscriptionsCollectionApi.CreateDCCFAnalyticsSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDCCFAnalyticsSubscription`: NdccfAnalyticsSubscription
    fmt.Fprintf(os.Stdout, "Response from `DCCFAnalyticsSubscriptionsCollectionApi.CreateDCCFAnalyticsSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDCCFAnalyticsSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndccfAnalyticsSubscription** | [**NdccfAnalyticsSubscription**](NdccfAnalyticsSubscription.md) | Contains the information for the creation the resource. | 

### Return type

[**NdccfAnalyticsSubscription**](NdccfAnalyticsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

