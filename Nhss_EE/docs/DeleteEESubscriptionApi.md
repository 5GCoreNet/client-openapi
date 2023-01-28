# \DeleteEESubscriptionApi

All URIs are relative to *https://example.com/nhss-ee/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEeSubscription**](DeleteEESubscriptionApi.md#DeleteEeSubscription) | **Delete** /{ueId}/ee-subscriptions/{subscriptionId} | Unsubscribe



## DeleteEeSubscription

> DeleteEeSubscription(ctx, ueId, subscriptionId).Execute()

Unsubscribe

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
    ueId := "ueId_example" // string | IMSI of the subscriber
    subscriptionId := "subscriptionId_example" // string | Id of the EE Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeleteEESubscriptionApi.DeleteEeSubscription(context.Background(), ueId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeleteEESubscriptionApi.DeleteEeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | IMSI of the subscriber | 
**subscriptionId** | **string** | Id of the EE Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

