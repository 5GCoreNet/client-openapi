# \UpdateEESubscriptionApi

All URIs are relative to *https://example.com/nhss-ee/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateEeSubscription**](UpdateEESubscriptionApi.md#UpdateEeSubscription) | **Patch** /{ueId}/ee-subscriptions/{subscriptionId} | Patch



## UpdateEeSubscription

> PatchResult UpdateEeSubscription(ctx, ueId, subscriptionId).PatchItem(patchItem).Execute()

Patch

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
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateEESubscriptionApi.UpdateEeSubscription(context.Background(), ueId, subscriptionId).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateEESubscriptionApi.UpdateEeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEeSubscription`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `UpdateEESubscriptionApi.UpdateEeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | IMSI of the subscriber | 
**subscriptionId** | **string** | Id of the EE Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 

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

