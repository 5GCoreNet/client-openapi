# \NIDDAuthorizationInfoDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNIDDAuthorizationInfo**](NIDDAuthorizationInfoDocumentApi.md#CreateNIDDAuthorizationInfo) | **Put** /subscription-data/{ueId}/context-data/nidd-authorizations | Create NIDD Authorization Info
[**GetNiddAuthorizationInfo**](NIDDAuthorizationInfoDocumentApi.md#GetNiddAuthorizationInfo) | **Get** /subscription-data/{ueId}/context-data/nidd-authorizations | Retrieve NIDD Authorization Info
[**ModifyNiddAuthorizationInfo**](NIDDAuthorizationInfoDocumentApi.md#ModifyNiddAuthorizationInfo) | **Patch** /subscription-data/{ueId}/context-data/nidd-authorizations | Modify NIDD Authorization Info
[**RemoveNiddAuthorizationInfo**](NIDDAuthorizationInfoDocumentApi.md#RemoveNiddAuthorizationInfo) | **Delete** /subscription-data/{ueId}/context-data/nidd-authorizations | Delete NIDD Authorization Info



## CreateNIDDAuthorizationInfo

> NiddAuthorizationInfo CreateNIDDAuthorizationInfo(ctx, ueId).NiddAuthorizationInfo(niddAuthorizationInfo).Execute()

Create NIDD Authorization Info

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
    ueId := "ueId_example" // string | 
    niddAuthorizationInfo := *openapiclient.NewNiddAuthorizationInfo([]openapiclient.AuthorizationInfo{*openapiclient.NewAuthorizationInfo(*openapiclient.NewSnssai(int32(123)), "Dnn_example", "MtcProviderInformation_example", "AuthUpdateCallbackUri_example")}) // NiddAuthorizationInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NIDDAuthorizationInfoDocumentApi.CreateNIDDAuthorizationInfo(context.Background(), ueId).NiddAuthorizationInfo(niddAuthorizationInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NIDDAuthorizationInfoDocumentApi.CreateNIDDAuthorizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNIDDAuthorizationInfo`: NiddAuthorizationInfo
    fmt.Fprintf(os.Stdout, "Response from `NIDDAuthorizationInfoDocumentApi.CreateNIDDAuthorizationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNIDDAuthorizationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **niddAuthorizationInfo** | [**NiddAuthorizationInfo**](NiddAuthorizationInfo.md) |  | 

### Return type

[**NiddAuthorizationInfo**](NiddAuthorizationInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNiddAuthorizationInfo

> NiddAuthorizationInfo GetNiddAuthorizationInfo(ctx, ueId).Execute()

Retrieve NIDD Authorization Info

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
    ueId := "ueId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NIDDAuthorizationInfoDocumentApi.GetNiddAuthorizationInfo(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NIDDAuthorizationInfoDocumentApi.GetNiddAuthorizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNiddAuthorizationInfo`: NiddAuthorizationInfo
    fmt.Fprintf(os.Stdout, "Response from `NIDDAuthorizationInfoDocumentApi.GetNiddAuthorizationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNiddAuthorizationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NiddAuthorizationInfo**](NiddAuthorizationInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyNiddAuthorizationInfo

> PatchResult ModifyNiddAuthorizationInfo(ctx, ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify NIDD Authorization Info

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
    ueId := "ueId_example" // string | 
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NIDDAuthorizationInfoDocumentApi.ModifyNiddAuthorizationInfo(context.Background(), ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NIDDAuthorizationInfoDocumentApi.ModifyNiddAuthorizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyNiddAuthorizationInfo`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `NIDDAuthorizationInfoDocumentApi.ModifyNiddAuthorizationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyNiddAuthorizationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNiddAuthorizationInfo

> RemoveNiddAuthorizationInfo(ctx, ueId).Execute()

Delete NIDD Authorization Info

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
    ueId := "ueId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NIDDAuthorizationInfoDocumentApi.RemoveNiddAuthorizationInfo(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NIDDAuthorizationInfoDocumentApi.RemoveNiddAuthorizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNiddAuthorizationInfoRequest struct via the builder pattern


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

