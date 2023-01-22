# \SubscriptionModificationApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModifySharedDataSubs**](SubscriptionModificationApi.md#ModifySharedDataSubs) | **Patch** /shared-data-subscriptions/{subscriptionId} | modify the subscription



## ModifySharedDataSubs

> PatchResult ModifySharedDataSubs(ctx, subscriptionId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

modify the subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Id of the SDM Subscription
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionModificationApi.ModifySharedDataSubs(context.Background(), subscriptionId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModificationApi.ModifySharedDataSubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifySharedDataSubs`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionModificationApi.ModifySharedDataSubs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Id of the SDM Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySharedDataSubsRequest struct via the builder pattern


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

