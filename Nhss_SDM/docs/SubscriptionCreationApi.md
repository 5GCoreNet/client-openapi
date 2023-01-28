# \SubscriptionCreationApi

All URIs are relative to *https://example.com/nhss-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Subscribe**](SubscriptionCreationApi.md#Subscribe) | **Post** /{ueId}/subscriptions | subscribe to notifications



## Subscribe

> SubscriptionData Subscribe(ctx, ueId).SubscriptionData(subscriptionData).Execute()

subscribe to notifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_SDM"
)

func main() {
    ueId := "ueId_example" // string | IMSI of the user
    subscriptionData := *openapiclient.NewSubscriptionData("NfInstanceId_example", "CallbackReference_example", []string{"MonitoredResourceUris_example"}) // SubscriptionData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionCreationApi.Subscribe(context.Background(), ueId).SubscriptionData(subscriptionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionCreationApi.Subscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Subscribe`: SubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionCreationApi.Subscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | IMSI of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionData** | [**SubscriptionData**](SubscriptionData.md) |  | 

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

