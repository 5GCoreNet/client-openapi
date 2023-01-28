# \SDMSubscriptionDeletionApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImsSdmUnsubscribe**](SDMSubscriptionDeletionApi.md#ImsSdmUnsubscribe) | **Delete** /{imsUeId}/subscriptions/{subscriptionId} | unsubscribe from notifications



## ImsSdmUnsubscribe

> ImsSdmUnsubscribe(ctx, imsUeId, subscriptionId).Execute()

unsubscribe from notifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsSDM"
)

func main() {
    imsUeId := "imsUeId_example" // string | IMS Public Identity
    subscriptionId := "subscriptionId_example" // string | Id of the Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDMSubscriptionDeletionApi.ImsSdmUnsubscribe(context.Background(), imsUeId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDMSubscriptionDeletionApi.ImsSdmUnsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Public Identity | 
**subscriptionId** | **string** | Id of the Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiImsSdmUnsubscribeRequest struct via the builder pattern


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

