# \GBASDMSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nhss-gba-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GbaSdmSubscribe**](GBASDMSubscriptionsCollectionApi.md#GbaSdmSubscribe) | **Post** /{ueId}/subscriptions | Subscribe to notifications



## GbaSdmSubscribe

> GbaSdmSubscription GbaSdmSubscribe(ctx, ueId).GbaSdmSubscription(gbaSdmSubscription).Execute()

Subscribe to notifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_gbaSDM"
)

func main() {
    ueId := *openapiclient.NewUeId() // UeId | UE Identity
    gbaSdmSubscription := *openapiclient.NewGbaSdmSubscription("NfInstanceId_example", "CallbackReference_example", []string{"MonitoredResourceUris_example"}) // GbaSdmSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GBASDMSubscriptionsCollectionApi.GbaSdmSubscribe(context.Background(), ueId).GbaSdmSubscription(gbaSdmSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GBASDMSubscriptionsCollectionApi.GbaSdmSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GbaSdmSubscribe`: GbaSdmSubscription
    fmt.Fprintf(os.Stdout, "Response from `GBASDMSubscriptionsCollectionApi.GbaSdmSubscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**UeId**](.md) | UE Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGbaSdmSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gbaSdmSubscription** | [**GbaSdmSubscription**](GbaSdmSubscription.md) |  | 

### Return type

[**GbaSdmSubscription**](GbaSdmSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

