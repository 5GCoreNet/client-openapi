# \IndividualMessageDeliveryDocumentApi

All URIs are relative to *https://example.com/vae-message-delivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMessageDelivery**](IndividualMessageDeliveryDocumentApi.md#DeleteMessageDelivery) | **Delete** /subscriptions/{subscriptionId}/message-deliveries/{dlDeliveryId} | VAE Message delivery resource delete service Operation



## DeleteMessageDelivery

> DeleteMessageDelivery(ctx, subscriptionId, dlDeliveryId).Execute()

VAE Message delivery resource delete service Operation

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
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Individual Message Delivery Subscription
    dlDeliveryId := "dlDeliveryId_example" // string | Unique ID of the message delivery to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMessageDeliveryDocumentApi.DeleteMessageDelivery(context.Background(), subscriptionId, dlDeliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMessageDeliveryDocumentApi.DeleteMessageDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Individual Message Delivery Subscription | 
**dlDeliveryId** | **string** | Unique ID of the message delivery to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageDeliveryRequest struct via the builder pattern


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

