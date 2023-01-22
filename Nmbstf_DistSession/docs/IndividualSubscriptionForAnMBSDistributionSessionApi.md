# \IndividualSubscriptionForAnMBSDistributionSessionApi

All URIs are relative to *https://example.com/nmbstf-distsession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusSubscribeMod**](IndividualSubscriptionForAnMBSDistributionSessionApi.md#StatusSubscribeMod) | **Patch** /dist-sessions/{distSessionRef}/subscriptions/{subscriptionId} | StatusSubscribe to modify (update or renew) an individual subscription



## StatusSubscribeMod

> DistSessionSubscription StatusSubscribeMod(ctx, subscriptionId, distSessionRef).PatchItem(patchItem).Execute()

StatusSubscribe to modify (update or renew) an individual subscription

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the individual subscription to be modified
    distSessionRef := "distSessionRef_example" // string | Unique ID of the MBS distribution session
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | Data to be modified in the DistSessionSubscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionForAnMBSDistributionSessionApi.StatusSubscribeMod(context.Background(), subscriptionId, distSessionRef).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionForAnMBSDistributionSessionApi.StatusSubscribeMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusSubscribeMod`: DistSessionSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionForAnMBSDistributionSessionApi.StatusSubscribeMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the individual subscription to be modified | 
**distSessionRef** | **string** | Unique ID of the MBS distribution session | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusSubscribeModRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchItem** | [**[]PatchItem**](PatchItem.md) | Data to be modified in the DistSessionSubscription | 

### Return type

[**DistSessionSubscription**](DistSessionSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

