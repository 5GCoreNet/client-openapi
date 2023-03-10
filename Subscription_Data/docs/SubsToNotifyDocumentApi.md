# \SubsToNotifyDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModifysubscriptionDataSubscription**](SubsToNotifyDocumentApi.md#ModifysubscriptionDataSubscription) | **Patch** /subscription-data/subs-to-notify/{subsId} | Modify an individual subscriptionDataSubscription
[**QuerySubscriptionDataSubscriptions**](SubsToNotifyDocumentApi.md#QuerySubscriptionDataSubscriptions) | **Get** /subscription-data/subs-to-notify/{subsId} | Retrieves a individual subscriptionDataSubscription identified by subsId
[**RemovesubscriptionDataSubscriptions**](SubsToNotifyDocumentApi.md#RemovesubscriptionDataSubscriptions) | **Delete** /subscription-data/subs-to-notify/{subsId} | Deletes a subscriptionDataSubscriptions



## ModifysubscriptionDataSubscription

> ModifysubscriptionDataSubscription200Response ModifysubscriptionDataSubscription(ctx, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify an individual subscriptionDataSubscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    subsId := "subsId_example" // string | 
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubsToNotifyDocumentApi.ModifysubscriptionDataSubscription(context.Background(), subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubsToNotifyDocumentApi.ModifysubscriptionDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifysubscriptionDataSubscription`: ModifysubscriptionDataSubscription200Response
    fmt.Fprintf(os.Stdout, "Response from `SubsToNotifyDocumentApi.ModifysubscriptionDataSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifysubscriptionDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**ModifysubscriptionDataSubscription200Response**](ModifysubscriptionDataSubscription200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySubscriptionDataSubscriptions

> []SubscriptionDataSubscriptions QuerySubscriptionDataSubscriptions(ctx, subsId).Execute()

Retrieves a individual subscriptionDataSubscription identified by subsId

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    subsId := "subsId_example" // string | Unique ID of the subscription to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubsToNotifyDocumentApi.QuerySubscriptionDataSubscriptions(context.Background(), subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubsToNotifyDocumentApi.QuerySubscriptionDataSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySubscriptionDataSubscriptions`: []SubscriptionDataSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `SubsToNotifyDocumentApi.QuerySubscriptionDataSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subsId** | **string** | Unique ID of the subscription to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySubscriptionDataSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SubscriptionDataSubscriptions**](SubscriptionDataSubscriptions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovesubscriptionDataSubscriptions

> RemovesubscriptionDataSubscriptions(ctx, subsId).Execute()

Deletes a subscriptionDataSubscriptions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    subsId := "subsId_example" // string | Unique ID of the subscription to remove

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubsToNotifyDocumentApi.RemovesubscriptionDataSubscriptions(context.Background(), subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubsToNotifyDocumentApi.RemovesubscriptionDataSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subsId** | **string** | Unique ID of the subscription to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovesubscriptionDataSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

