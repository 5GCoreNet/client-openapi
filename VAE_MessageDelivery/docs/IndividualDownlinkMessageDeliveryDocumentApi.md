# \IndividualDownlinkMessageDeliveryDocumentApi

All URIs are relative to *https://example.com/vae-message-delivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadIndividualDownlinkMessageDelivery**](IndividualDownlinkMessageDeliveryDocumentApi.md#ReadIndividualDownlinkMessageDelivery) | **Get** /subscriptions/{subscriptionId}/message-deliveries/{dlDeliveryId} | VAE Message delivery resource Read service Operation



## ReadIndividualDownlinkMessageDelivery

> DownlinkMessageDeliveryData ReadIndividualDownlinkMessageDelivery(ctx, subscriptionId, dlDeliveryId).Execute()

VAE Message delivery resource Read service Operation

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
    dlDeliveryId := "dlDeliveryId_example" // string | Identifier of a downlink messge delivery resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDownlinkMessageDeliveryDocumentApi.ReadIndividualDownlinkMessageDelivery(context.Background(), subscriptionId, dlDeliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDownlinkMessageDeliveryDocumentApi.ReadIndividualDownlinkMessageDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualDownlinkMessageDelivery`: DownlinkMessageDeliveryData
    fmt.Fprintf(os.Stdout, "Response from `IndividualDownlinkMessageDeliveryDocumentApi.ReadIndividualDownlinkMessageDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Individual Message Delivery Subscription | 
**dlDeliveryId** | **string** | Identifier of a downlink messge delivery resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualDownlinkMessageDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DownlinkMessageDeliveryData**](DownlinkMessageDeliveryData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

