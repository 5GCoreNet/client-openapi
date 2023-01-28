# \SubscriptionCreationForSharedDataApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribeToSharedData**](SubscriptionCreationForSharedDataApi.md#SubscribeToSharedData) | **Post** /shared-data-subscriptions | subscribe to notifications for shared data



## SubscribeToSharedData

> ImsSdmSubscription SubscribeToSharedData(ctx).ImsSdmSubscription(imsSdmSubscription).Execute()

subscribe to notifications for shared data

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
    imsSdmSubscription := *openapiclient.NewImsSdmSubscription("NfInstanceId_example", "CallbackReference_example", []string{"MonitoredResourceUris_example"}) // ImsSdmSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionCreationForSharedDataApi.SubscribeToSharedData(context.Background()).ImsSdmSubscription(imsSdmSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionCreationForSharedDataApi.SubscribeToSharedData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscribeToSharedData`: ImsSdmSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionCreationForSharedDataApi.SubscribeToSharedData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeToSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imsSdmSubscription** | [**ImsSdmSubscription**](ImsSdmSubscription.md) |  | 

### Return type

[**ImsSdmSubscription**](ImsSdmSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

