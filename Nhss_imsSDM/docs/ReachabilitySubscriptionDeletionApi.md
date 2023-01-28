# \ReachabilitySubscriptionDeletionApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UeReachUnsubscribe**](ReachabilitySubscriptionDeletionApi.md#UeReachUnsubscribe) | **Delete** /{imsUeId}/access-data/ps-domain/ue-reach-subscriptions/{subscriptionId} | unsubscribe from notifications to UE reachability



## UeReachUnsubscribe

> UeReachUnsubscribe(ctx, imsUeId, subscriptionId).PrivateIdentity(privateIdentity).Execute()

unsubscribe from notifications to UE reachability

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
    imsUeId := "imsUeId_example" // string | IMS Identity
    subscriptionId := "subscriptionId_example" // string | Id of the Subscription
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReachabilitySubscriptionDeletionApi.UeReachUnsubscribe(context.Background(), imsUeId, subscriptionId).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReachabilitySubscriptionDeletionApi.UeReachUnsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 
**subscriptionId** | **string** | Id of the Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUeReachUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateIdentity** | **string** | IMS Private Identity | 

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

