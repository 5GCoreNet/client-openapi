# \GBASDMSubscriptionDocumentApi

All URIs are relative to *https://example.com/nhss-gba-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GbaSdmSubsModify**](GBASDMSubscriptionDocumentApi.md#GbaSdmSubsModify) | **Patch** /{ueId}/subscriptions/{subscriptionId} | Modify the subscription
[**GbaSdmUnsubscribe**](GBASDMSubscriptionDocumentApi.md#GbaSdmUnsubscribe) | **Delete** /{ueId}/subscriptions/{subscriptionId} | Unsubscribe from notifications



## GbaSdmSubsModify

> PatchResult GbaSdmSubsModify(ctx, ueId, subscriptionId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify the subscription

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
    subscriptionId := "subscriptionId_example" // string | Id of the Subscription
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GBASDMSubscriptionDocumentApi.GbaSdmSubsModify(context.Background(), ueId, subscriptionId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GBASDMSubscriptionDocumentApi.GbaSdmSubsModify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GbaSdmSubsModify`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `GBASDMSubscriptionDocumentApi.GbaSdmSubsModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**UeId**](.md) | UE Identity | 
**subscriptionId** | **string** | Id of the Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGbaSdmSubsModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GbaSdmUnsubscribe

> GbaSdmUnsubscribe(ctx, ueId, subscriptionId).Execute()

Unsubscribe from notifications

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
    subscriptionId := "subscriptionId_example" // string | Id of the Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GBASDMSubscriptionDocumentApi.GbaSdmUnsubscribe(context.Background(), ueId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GBASDMSubscriptionDocumentApi.GbaSdmUnsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**UeId**](.md) | UE Identity | 
**subscriptionId** | **string** | Id of the Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGbaSdmUnsubscribeRequest struct via the builder pattern


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

