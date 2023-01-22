# \SubscriptionIDDocumentApi

All URIs are relative to *https://example.com/nnssf-nssaiavailability/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NSSAIAvailabilitySubModifyPatch**](SubscriptionIDDocumentApi.md#NSSAIAvailabilitySubModifyPatch) | **Patch** /nssai-availability/subscriptions/{subscriptionId} | updates an already existing NSSAI availability notification subscription
[**NSSAIAvailabilityUnsubscribe**](SubscriptionIDDocumentApi.md#NSSAIAvailabilityUnsubscribe) | **Delete** /nssai-availability/subscriptions/{subscriptionId} | Deletes an already existing NSSAI availability notification subscription



## NSSAIAvailabilitySubModifyPatch

> NssfEventSubscriptionCreatedData NSSAIAvailabilitySubModifyPatch(ctx, subscriptionId).PatchItem(patchItem).ContentEncoding(contentEncoding).Execute()

updates an already existing NSSAI availability notification subscription

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription for notification
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | JSON Patch instructions to update at the NSSF, the NSSAI availability notification subscription
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionIDDocumentApi.NSSAIAvailabilitySubModifyPatch(context.Background(), subscriptionId).PatchItem(patchItem).ContentEncoding(contentEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionIDDocumentApi.NSSAIAvailabilitySubModifyPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NSSAIAvailabilitySubModifyPatch`: NssfEventSubscriptionCreatedData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionIDDocumentApi.NSSAIAvailabilitySubModifyPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the subscription for notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiNSSAIAvailabilitySubModifyPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) | JSON Patch instructions to update at the NSSF, the NSSAI availability notification subscription | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 

### Return type

[**NssfEventSubscriptionCreatedData**](NssfEventSubscriptionCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json:
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NSSAIAvailabilityUnsubscribe

> NSSAIAvailabilityUnsubscribe(ctx, subscriptionId).Execute()

Deletes an already existing NSSAI availability notification subscription

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription for notification

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionIDDocumentApi.NSSAIAvailabilityUnsubscribe(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionIDDocumentApi.NSSAIAvailabilityUnsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the subscription for notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiNSSAIAvailabilityUnsubscribeRequest struct via the builder pattern


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

