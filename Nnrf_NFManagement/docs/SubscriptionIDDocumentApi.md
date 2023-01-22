# \SubscriptionIDDocumentApi

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveSubscription**](SubscriptionIDDocumentApi.md#RemoveSubscription) | **Delete** /subscriptions/{subscriptionID} | Deletes a subscription
[**UpdateSubscription**](SubscriptionIDDocumentApi.md#UpdateSubscription) | **Patch** /subscriptions/{subscriptionID} | Updates a subscription



## RemoveSubscription

> RemoveSubscription(ctx, subscriptionID).Execute()

Deletes a subscription

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
    subscriptionID := "subscriptionID_example" // string | Unique ID of the subscription to remove

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionIDDocumentApi.RemoveSubscription(context.Background(), subscriptionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionIDDocumentApi.RemoveSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | **string** | Unique ID of the subscription to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSubscriptionRequest struct via the builder pattern


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


## UpdateSubscription

> SubscriptionData UpdateSubscription(ctx, subscriptionID).PatchItem(patchItem).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()

Updates a subscription

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
    subscriptionID := "subscriptionID_example" // string | Unique ID of the subscription to update
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionIDDocumentApi.UpdateSubscription(context.Background(), subscriptionID).PatchItem(patchItem).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionIDDocumentApi.UpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscription`: SubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionIDDocumentApi.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | **string** | Unique ID of the subscription to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 

### Return type

[**SubscriptionData**](SubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

