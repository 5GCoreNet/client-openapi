# \MessageDeliveriesCollectionCollectionApi

All URIs are relative to *https://example.com/vae-message-delivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDownlinkMessageDelivery**](MessageDeliveriesCollectionCollectionApi.md#CreateDownlinkMessageDelivery) | **Post** /subscriptions/{subscriptionId}/message-deliveries | VAE Message delivery resource create service Operation



## CreateDownlinkMessageDelivery

> DownlinkMessageDeliveryData CreateDownlinkMessageDelivery(ctx, subscriptionId).DownlinkMessageDeliveryData(downlinkMessageDeliveryData).Execute()

VAE Message delivery resource create service Operation

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
    downlinkMessageDeliveryData := *openapiclient.NewDownlinkMessageDeliveryData(string(123)) // DownlinkMessageDeliveryData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageDeliveriesCollectionCollectionApi.CreateDownlinkMessageDelivery(context.Background(), subscriptionId).DownlinkMessageDeliveryData(downlinkMessageDeliveryData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageDeliveriesCollectionCollectionApi.CreateDownlinkMessageDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDownlinkMessageDelivery`: DownlinkMessageDeliveryData
    fmt.Fprintf(os.Stdout, "Response from `MessageDeliveriesCollectionCollectionApi.CreateDownlinkMessageDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Individual Message Delivery Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownlinkMessageDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downlinkMessageDeliveryData** | [**DownlinkMessageDeliveryData**](DownlinkMessageDeliveryData.md) |  | 

### Return type

[**DownlinkMessageDeliveryData**](DownlinkMessageDeliveryData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

