# \IndividualSubscriptionDocumentApi

All URIs are relative to *https://example.com/nnsacf-slice-ee/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteModifySubscription**](IndividualSubscriptionDocumentApi.md#CompleteModifySubscription) | **Put** /subscriptions/{subscriptionId} | Nnsacf_SliceEventExposure Subscribe complete modify service Operation
[**DeleteSubscription**](IndividualSubscriptionDocumentApi.md#DeleteSubscription) | **Delete** /subscriptions/{subscriptionId} | Nnsacf_SliceEventExposure Unsubscribe service Operation
[**PartialModifySubscription**](IndividualSubscriptionDocumentApi.md#PartialModifySubscription) | **Patch** /subscriptions/{subscriptionId} | Nnsacf_SliceEventExposure Subscribe partial modify service Operation



## CompleteModifySubscription

> CreatedSACEventSubscription CompleteModifySubscription(ctx, subscriptionId).SACEventSubscription(sACEventSubscription).Execute()

Nnsacf_SliceEventExposure Subscribe complete modify service Operation

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the subscription to be modified
    sACEventSubscription := *openapiclient.NewSACEventSubscription(*openapiclient.NewSACEvent(*openapiclient.NewSACEventType(), []openapiclient.Snssai{*openapiclient.NewSnssai(int32(123))}), "EventNotifyUri_example", "NfId_example") // SACEventSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.CompleteModifySubscription(context.Background(), subscriptionId).SACEventSubscription(sACEventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.CompleteModifySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteModifySubscription`: CreatedSACEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionDocumentApi.CompleteModifySubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the subscription to be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteModifySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sACEventSubscription** | [**SACEventSubscription**](SACEventSubscription.md) |  | 

### Return type

[**CreatedSACEventSubscription**](CreatedSACEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, subscriptionId).Execute()

Nnsacf_SliceEventExposure Unsubscribe service Operation

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the subscription to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.DeleteSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the subscription to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


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


## PartialModifySubscription

> CreatedSACEventSubscription PartialModifySubscription(ctx, subscriptionId).PatchItem(patchItem).Execute()

Nnsacf_SliceEventExposure Subscribe partial modify service Operation

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the subscription to be modified
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.PartialModifySubscription(context.Background(), subscriptionId).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.PartialModifySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialModifySubscription`: CreatedSACEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionDocumentApi.PartialModifySubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the subscription to be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialModifySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 

### Return type

[**CreatedSACEventSubscription**](CreatedSACEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

