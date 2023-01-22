# \SubscriptionModificationApi

All URIs are relative to *https://example.com/nhss-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Modify**](SubscriptionModificationApi.md#Modify) | **Patch** /{ueId}/subscriptions/{subscriptionId} | modify the subscription



## Modify

> Modify(ctx, ueId, subscriptionId).PatchItem(patchItem).Execute()

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
    ueId := "ueId_example" // string | IMSI of the user
    subscriptionId := "subscriptionId_example" // string | Id of the Subscription
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionModificationApi.Modify(context.Background(), ueId, subscriptionId).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModificationApi.Modify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | IMSI of the user | 
**subscriptionId** | **string** | Id of the Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

