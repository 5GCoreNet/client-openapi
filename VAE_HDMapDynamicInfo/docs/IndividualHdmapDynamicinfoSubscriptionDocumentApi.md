# \IndividualHdmapDynamicinfoSubscriptionDocumentApi

All URIs are relative to *https://example.com/vae-hdmap-dynamic-info/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHdMapDynamicInfoSubscription**](IndividualHdmapDynamicinfoSubscriptionDocumentApi.md#DeleteHdMapDynamicInfoSubscription) | **Delete** /subscriptions/{subscriptionId} | VAE HdMap DynamicInfo Subscription resource delete service Operation



## DeleteHdMapDynamicInfoSubscription

> DeleteHdMapDynamicInfoSubscription(ctx, subscriptionId).Execute()

VAE HdMap DynamicInfo Subscription resource delete service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_HDMapDynamicInfo"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Unique ID of the hdmap dynamicinfo subscription to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualHdmapDynamicinfoSubscriptionDocumentApi.DeleteHdMapDynamicInfoSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualHdmapDynamicinfoSubscriptionDocumentApi.DeleteHdMapDynamicInfoSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the hdmap dynamicinfo subscription to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHdMapDynamicInfoSubscriptionRequest struct via the builder pattern


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

