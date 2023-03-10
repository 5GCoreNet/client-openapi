# \EESubscriptionCollectionApi

All URIs are relative to *https://example.com/nhss-ee/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEeSubscription**](EESubscriptionCollectionApi.md#CreateEeSubscription) | **Post** /{ueId}/ee-subscriptions | Subscribe



## CreateEeSubscription

> CreatedEeSubscription CreateEeSubscription(ctx, ueId).EeSubscription(eeSubscription).Execute()

Subscribe

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_EE"
)

func main() {
    ueId := "ueId_example" // string | IMSI of the subscriber or the identity of a group of UEs
    eeSubscription := *openapiclient.NewEeSubscription("CallbackReference_example") // EeSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EESubscriptionCollectionApi.CreateEeSubscription(context.Background(), ueId).EeSubscription(eeSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EESubscriptionCollectionApi.CreateEeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEeSubscription`: CreatedEeSubscription
    fmt.Fprintf(os.Stdout, "Response from `EESubscriptionCollectionApi.CreateEeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | IMSI of the subscriber or the identity of a group of UEs | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eeSubscription** | [**EeSubscription**](EeSubscription.md) |  | 

### Return type

[**CreatedEeSubscription**](CreatedEeSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

