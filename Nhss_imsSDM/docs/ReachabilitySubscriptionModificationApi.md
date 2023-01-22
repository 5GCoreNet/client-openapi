# \ReachabilitySubscriptionModificationApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UeReachSubsModify**](ReachabilitySubscriptionModificationApi.md#UeReachSubsModify) | **Patch** /{imsUeId}/access-data/ps-domain/ue-reach-subscriptions/{subscriptionId} | modify the subscription



## UeReachSubsModify

> PatchResult UeReachSubsModify(ctx, imsUeId, subscriptionId).PatchItem(patchItem).PrivateIdentity(privateIdentity).SupportedFeatures(supportedFeatures).Execute()

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
    imsUeId := "imsUeId_example" // string | IMS Identity
    subscriptionId := "subscriptionId_example" // string | Id of the Subscription
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReachabilitySubscriptionModificationApi.UeReachSubsModify(context.Background(), imsUeId, subscriptionId).PatchItem(patchItem).PrivateIdentity(privateIdentity).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReachabilitySubscriptionModificationApi.UeReachSubsModify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UeReachSubsModify`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `ReachabilitySubscriptionModificationApi.UeReachSubsModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 
**subscriptionId** | **string** | Id of the Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUeReachSubsModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **privateIdentity** | **string** | IMS Private Identity | 
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

