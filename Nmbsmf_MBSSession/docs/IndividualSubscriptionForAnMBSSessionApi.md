# \IndividualSubscriptionForAnMBSSessionApi

All URIs are relative to *https://example.com/nmbsmf-mbssession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusSubscribeMod**](IndividualSubscriptionForAnMBSSessionApi.md#StatusSubscribeMod) | **Patch** /mbs-sessions/subscriptions/{subscriptionId} | StatusSubscribe to modify (update or renew) an individual subscription
[**StatusUnSubscribe**](IndividualSubscriptionForAnMBSSessionApi.md#StatusUnSubscribe) | **Delete** /mbs-sessions/subscriptions/{subscriptionId} | StatusUnSubscribe to unsubscribe from the Status Subscription



## StatusSubscribeMod

> MbsSessionSubscription StatusSubscribeMod(ctx, subscriptionId).PatchItem(patchItem).Execute()

StatusSubscribe to modify (update or renew) an individual subscription

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the individual subscription to be modified
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | Data to be modified in the MBSSessionSubscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionForAnMBSSessionApi.StatusSubscribeMod(context.Background(), subscriptionId).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionForAnMBSSessionApi.StatusSubscribeMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusSubscribeMod`: MbsSessionSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionForAnMBSSessionApi.StatusSubscribeMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the individual subscription to be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusSubscribeModRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) | Data to be modified in the MBSSessionSubscription | 

### Return type

[**MbsSessionSubscription**](MbsSessionSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusUnSubscribe

> StatusUnSubscribe(ctx, subscriptionId).Execute()

StatusUnSubscribe to unsubscribe from the Status Subscription

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
    resp, r, err := apiClient.IndividualSubscriptionForAnMBSSessionApi.StatusUnSubscribe(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionForAnMBSSessionApi.StatusUnSubscribe``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStatusUnSubscribeRequest struct via the builder pattern


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

