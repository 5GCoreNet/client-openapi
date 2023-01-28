# \IndividualSubscriptionForAnMBSContextApi

All URIs are relative to *https://example.com/nmbsmf-mbssession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContextStatusSubscribeMod**](IndividualSubscriptionForAnMBSContextApi.md#ContextStatusSubscribeMod) | **Patch** /mbs-sessions/contexts/subscriptions/{subscriptionId} | ContextStatusSubscribe modifying an individual subscription
[**ContextStatusUnSubscribe**](IndividualSubscriptionForAnMBSContextApi.md#ContextStatusUnSubscribe) | **Delete** /mbs-sessions/contexts/subscriptions/{subscriptionId} | ContextStatusUnSubscribe



## ContextStatusSubscribeMod

> ContextStatusSubscription ContextStatusSubscribeMod(ctx, subscriptionId).PatchItem(patchItem).Execute()

ContextStatusSubscribe modifying an individual subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsmf_MBSSession"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Unique ID of the subscription to be modified
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | Data within the ContextStatusSubscribe Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionForAnMBSContextApi.ContextStatusSubscribeMod(context.Background(), subscriptionId).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionForAnMBSContextApi.ContextStatusSubscribeMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContextStatusSubscribeMod`: ContextStatusSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionForAnMBSContextApi.ContextStatusSubscribeMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the subscription to be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiContextStatusSubscribeModRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) | Data within the ContextStatusSubscribe Request | 

### Return type

[**ContextStatusSubscription**](ContextStatusSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContextStatusUnSubscribe

> ContextStatusUnSubscribe(ctx, subscriptionId).Execute()

ContextStatusUnSubscribe

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsmf_MBSSession"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Unique ID of the subscription to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionForAnMBSContextApi.ContextStatusUnSubscribe(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionForAnMBSContextApi.ContextStatusUnSubscribe``: %v\n", err)
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

Other parameters are passed through a pointer to a apiContextStatusUnSubscribeRequest struct via the builder pattern


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

