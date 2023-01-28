# \IndividualMessageDeliverySubscriptionDocumentApi

All URIs are relative to *https://example.com/vae-message-delivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMessageDeliverySubscription**](IndividualMessageDeliverySubscriptionDocumentApi.md#DeleteMessageDeliverySubscription) | **Delete** /subscriptions/{subscriptionId} | Delete an individual Message Delivery Subscription resource
[**ReadIndividualMessageDeliverySubscription**](IndividualMessageDeliverySubscriptionDocumentApi.md#ReadIndividualMessageDeliverySubscription) | **Get** /subscriptions/{subscriptionId} | Get an existing individual Message Delivery Subscription resource



## DeleteMessageDeliverySubscription

> DeleteMessageDeliverySubscription(ctx, subscriptionId).Execute()

Delete an individual Message Delivery Subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_MessageDelivery"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Individual Message Delivery Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMessageDeliverySubscriptionDocumentApi.DeleteMessageDeliverySubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMessageDeliverySubscriptionDocumentApi.DeleteMessageDeliverySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Individual Message Delivery Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageDeliverySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIndividualMessageDeliverySubscription

> MessageDeliverySubscriptionData ReadIndividualMessageDeliverySubscription(ctx, subscriptionId).Execute()

Get an existing individual Message Delivery Subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_MessageDelivery"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Individual Message Delivery Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMessageDeliverySubscriptionDocumentApi.ReadIndividualMessageDeliverySubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMessageDeliverySubscriptionDocumentApi.ReadIndividualMessageDeliverySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualMessageDeliverySubscription`: MessageDeliverySubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `IndividualMessageDeliverySubscriptionDocumentApi.ReadIndividualMessageDeliverySubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Individual Message Delivery Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualMessageDeliverySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageDeliverySubscriptionData**](MessageDeliverySubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

