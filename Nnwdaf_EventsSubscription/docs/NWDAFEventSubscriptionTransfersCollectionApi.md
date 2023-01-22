# \NWDAFEventSubscriptionTransfersCollectionApi

All URIs are relative to *https://example.com/nnwdaf-eventssubscription/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNWDAFEventSubscriptionTransfer**](NWDAFEventSubscriptionTransfersCollectionApi.md#CreateNWDAFEventSubscriptionTransfer) | **Post** /transfers | Provide information about requested analytics subscriptions transfer and potentially create a new Individual NWDAF Event Subscription Transfer resource.



## CreateNWDAFEventSubscriptionTransfer

> CreateNWDAFEventSubscriptionTransfer(ctx).AnalyticsSubscriptionsTransfer(analyticsSubscriptionsTransfer).Execute()

Provide information about requested analytics subscriptions transfer and potentially create a new Individual NWDAF Event Subscription Transfer resource.

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
    analyticsSubscriptionsTransfer := *openapiclient.NewAnalyticsSubscriptionsTransfer([]openapiclient.SubscriptionTransferInfo{*openapiclient.NewSubscriptionTransferInfo(*openapiclient.NewTransferRequestType(), *openapiclient.NewNnwdafEventsSubscription([]openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewNwdafEvent())}), "ConsumerId_example")}) // AnalyticsSubscriptionsTransfer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NWDAFEventSubscriptionTransfersCollectionApi.CreateNWDAFEventSubscriptionTransfer(context.Background()).AnalyticsSubscriptionsTransfer(analyticsSubscriptionsTransfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NWDAFEventSubscriptionTransfersCollectionApi.CreateNWDAFEventSubscriptionTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNWDAFEventSubscriptionTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsSubscriptionsTransfer** | [**AnalyticsSubscriptionsTransfer**](AnalyticsSubscriptionsTransfer.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

