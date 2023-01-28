# \NotificationSubscriptionsCRUDApi

All URIs are relative to *https://example.com/nudsf-dr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationSubscriptions**](NotificationSubscriptionsCRUDApi.md#GetNotificationSubscriptions) | **Get** /{realmId}/{storageId}/subs-to-notify | Notification subscription retrieval



## GetNotificationSubscriptions

> []NotificationSubscription GetNotificationSubscriptions(ctx, realmId, storageId).LimitRange(limitRange).SupportedFeatures(supportedFeatures).Execute()

Notification subscription retrieval



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudsf_DataRepository"
)

func main() {
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    limitRange := int32(56) // int32 | The maximum number of NotificationSubscriptions to fetch (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSubscriptionsCRUDApi.GetNotificationSubscriptions(context.Background(), realmId, storageId).LimitRange(limitRange).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSubscriptionsCRUDApi.GetNotificationSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationSubscriptions`: []NotificationSubscription
    fmt.Fprintf(os.Stdout, "Response from `NotificationSubscriptionsCRUDApi.GetNotificationSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limitRange** | **int32** | The maximum number of NotificationSubscriptions to fetch | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**[]NotificationSubscription**](NotificationSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

