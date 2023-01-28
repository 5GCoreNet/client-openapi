# \SMFEventSubscriptionInfoDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSMFSubscriptions**](SMFEventSubscriptionInfoDocumentApi.md#CreateSMFSubscriptions) | **Put** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/smf-subscriptions | Create SMF Subscription Info
[**GetSmfGroupSubscriptions**](SMFEventSubscriptionInfoDocumentApi.md#GetSmfGroupSubscriptions) | **Get** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/smf-subscriptions | Retrieve SMF Subscription Info for a group of UEs or any UE
[**GetSmfSubscriptionInfo**](SMFEventSubscriptionInfoDocumentApi.md#GetSmfSubscriptionInfo) | **Get** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/smf-subscriptions | Retrieve SMF Subscription Info
[**ModifySmfGroupSubscriptions**](SMFEventSubscriptionInfoDocumentApi.md#ModifySmfGroupSubscriptions) | **Patch** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/smf-subscriptions | Modify SMF Subscription Info for a group of UEs or any UE
[**ModifySmfSubscriptionInfo**](SMFEventSubscriptionInfoDocumentApi.md#ModifySmfSubscriptionInfo) | **Patch** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/smf-subscriptions | Modify SMF Subscription Info
[**RemoveSmfGroupSubscriptions**](SMFEventSubscriptionInfoDocumentApi.md#RemoveSmfGroupSubscriptions) | **Delete** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/smf-subscriptions | Delete SMF Subscription Info for a group of UEs or any UE
[**RemoveSmfSubscriptionsInfo**](SMFEventSubscriptionInfoDocumentApi.md#RemoveSmfSubscriptionsInfo) | **Delete** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/smf-subscriptions | Delete SMF Subscription Info



## CreateSMFSubscriptions

> SmfSubscriptionInfo CreateSMFSubscriptions(ctx, ueId, subsId).SmfSubscriptionInfo(smfSubscriptionInfo).Execute()

Create SMF Subscription Info

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
    subsId := "subsId_example" // string | 
    smfSubscriptionInfo := *openapiclient.NewSmfSubscriptionInfo([]openapiclient.SmfSubscriptionItem{*openapiclient.NewSmfSubscriptionItem("SmfInstanceId_example", "SubscriptionId_example")}) // SmfSubscriptionInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFEventSubscriptionInfoDocumentApi.CreateSMFSubscriptions(context.Background(), ueId, subsId).SmfSubscriptionInfo(smfSubscriptionInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFEventSubscriptionInfoDocumentApi.CreateSMFSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSMFSubscriptions`: SmfSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `SMFEventSubscriptionInfoDocumentApi.CreateSMFSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSMFSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smfSubscriptionInfo** | [**SmfSubscriptionInfo**](SmfSubscriptionInfo.md) |  | 

### Return type

[**SmfSubscriptionInfo**](SmfSubscriptionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmfGroupSubscriptions

> SmfSubscriptionInfo GetSmfGroupSubscriptions(ctx, ueGroupId, subsId).Execute()

Retrieve SMF Subscription Info for a group of UEs or any UE

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
    ueGroupId := "ueGroupId_example" // string | 
    subsId := "subsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFEventSubscriptionInfoDocumentApi.GetSmfGroupSubscriptions(context.Background(), ueGroupId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFEventSubscriptionInfoDocumentApi.GetSmfGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmfGroupSubscriptions`: SmfSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `SMFEventSubscriptionInfoDocumentApi.GetSmfGroupSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmfGroupSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SmfSubscriptionInfo**](SmfSubscriptionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmfSubscriptionInfo

> SmfSubscriptionInfo GetSmfSubscriptionInfo(ctx, ueId, subsId).Execute()

Retrieve SMF Subscription Info

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
    subsId := "subsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFEventSubscriptionInfoDocumentApi.GetSmfSubscriptionInfo(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFEventSubscriptionInfoDocumentApi.GetSmfSubscriptionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmfSubscriptionInfo`: SmfSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `SMFEventSubscriptionInfoDocumentApi.GetSmfSubscriptionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmfSubscriptionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SmfSubscriptionInfo**](SmfSubscriptionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySmfGroupSubscriptions

> PatchResult ModifySmfGroupSubscriptions(ctx, ueGroupId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify SMF Subscription Info for a group of UEs or any UE

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
    ueGroupId := "ueGroupId_example" // string | 
    subsId := "subsId_example" // string | 
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFEventSubscriptionInfoDocumentApi.ModifySmfGroupSubscriptions(context.Background(), ueGroupId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFEventSubscriptionInfoDocumentApi.ModifySmfGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifySmfGroupSubscriptions`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `SMFEventSubscriptionInfoDocumentApi.ModifySmfGroupSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySmfGroupSubscriptionsRequest struct via the builder pattern


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


## ModifySmfSubscriptionInfo

> PatchResult ModifySmfSubscriptionInfo(ctx, ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify SMF Subscription Info

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
    subsId := "subsId_example" // string | 
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFEventSubscriptionInfoDocumentApi.ModifySmfSubscriptionInfo(context.Background(), ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFEventSubscriptionInfoDocumentApi.ModifySmfSubscriptionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifySmfSubscriptionInfo`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `SMFEventSubscriptionInfoDocumentApi.ModifySmfSubscriptionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySmfSubscriptionInfoRequest struct via the builder pattern


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


## RemoveSmfGroupSubscriptions

> RemoveSmfGroupSubscriptions(ctx, ueGroupId, subsId).Execute()

Delete SMF Subscription Info for a group of UEs or any UE

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
    ueGroupId := "ueGroupId_example" // string | 
    subsId := "subsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFEventSubscriptionInfoDocumentApi.RemoveSmfGroupSubscriptions(context.Background(), ueGroupId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFEventSubscriptionInfoDocumentApi.RemoveSmfGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSmfGroupSubscriptionsRequest struct via the builder pattern


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


## RemoveSmfSubscriptionsInfo

> RemoveSmfSubscriptionsInfo(ctx, ueId, subsId).Execute()

Delete SMF Subscription Info

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
    subsId := "subsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFEventSubscriptionInfoDocumentApi.RemoveSmfSubscriptionsInfo(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFEventSubscriptionInfoDocumentApi.RemoveSmfSubscriptionsInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSmfSubscriptionsInfoRequest struct via the builder pattern


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

