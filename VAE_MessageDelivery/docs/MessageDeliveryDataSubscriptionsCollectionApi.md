# \MessageDeliveryDataSubscriptionsCollectionApi

All URIs are relative to *https://example.com/vae-message-delivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualMessageDeliveryDataSubscription**](MessageDeliveryDataSubscriptionsCollectionApi.md#CreateIndividualMessageDeliveryDataSubscription) | **Post** /subscriptions | Create a new Individual Message Delivery Data Subscription resource



## CreateIndividualMessageDeliveryDataSubscription

> MessageDeliverySubscriptionData CreateIndividualMessageDeliveryDataSubscription(ctx).MessageDeliverySubscriptionData(messageDeliverySubscriptionData).Execute()

Create a new Individual Message Delivery Data Subscription resource

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
    messageDeliverySubscriptionData := *openapiclient.NewMessageDeliverySubscriptionData("AppSerId_example", "ServiceId_example", "NotifUri_example") // MessageDeliverySubscriptionData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageDeliveryDataSubscriptionsCollectionApi.CreateIndividualMessageDeliveryDataSubscription(context.Background()).MessageDeliverySubscriptionData(messageDeliverySubscriptionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageDeliveryDataSubscriptionsCollectionApi.CreateIndividualMessageDeliveryDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualMessageDeliveryDataSubscription`: MessageDeliverySubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `MessageDeliveryDataSubscriptionsCollectionApi.CreateIndividualMessageDeliveryDataSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualMessageDeliveryDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageDeliverySubscriptionData** | [**MessageDeliverySubscriptionData**](MessageDeliverySubscriptionData.md) |  | 

### Return type

[**MessageDeliverySubscriptionData**](MessageDeliverySubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

