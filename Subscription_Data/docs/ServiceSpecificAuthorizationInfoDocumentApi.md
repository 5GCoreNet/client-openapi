# \ServiceSpecificAuthorizationInfoDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceSpecificAuthorizationInfo**](ServiceSpecificAuthorizationInfoDocumentApi.md#CreateServiceSpecificAuthorizationInfo) | **Put** /subscription-data/{ueId}/context-data/service-specific-authorizations/{serviceType} | Create Service Specific Authorization Info
[**GetServiceSpecificAuthorizationInfo**](ServiceSpecificAuthorizationInfoDocumentApi.md#GetServiceSpecificAuthorizationInfo) | **Get** /subscription-data/{ueId}/context-data/service-specific-authorizations/{serviceType} | Retrieve Service Specific Authorization Info
[**ModifyServiceSpecificAuthorizationInfo**](ServiceSpecificAuthorizationInfoDocumentApi.md#ModifyServiceSpecificAuthorizationInfo) | **Patch** /subscription-data/{ueId}/context-data/service-specific-authorizations/{serviceType} | Modify Service Specific Authorization Info
[**RemoveServiceSpecificAuthorizationInfo**](ServiceSpecificAuthorizationInfoDocumentApi.md#RemoveServiceSpecificAuthorizationInfo) | **Delete** /subscription-data/{ueId}/context-data/service-specific-authorizations/{serviceType} | Delete Service Specific Authorization Info



## CreateServiceSpecificAuthorizationInfo

> ServiceSpecificAuthorizationInfo CreateServiceSpecificAuthorizationInfo(ctx, ueId, serviceType).ServiceSpecificAuthorizationInfo(serviceSpecificAuthorizationInfo).Execute()

Create Service Specific Authorization Info

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
    ueId := "ueId_example" // string | 
    serviceType := *openapiclient.NewServiceType() // ServiceType | Service Type
    serviceSpecificAuthorizationInfo := *openapiclient.NewServiceSpecificAuthorizationInfo([]openapiclient.AuthorizationInfo{*openapiclient.NewAuthorizationInfo(*openapiclient.NewSnssai(int32(123)), "Dnn_example", "MtcProviderInformation_example", "AuthUpdateCallbackUri_example")}) // ServiceSpecificAuthorizationInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceSpecificAuthorizationInfoDocumentApi.CreateServiceSpecificAuthorizationInfo(context.Background(), ueId, serviceType).ServiceSpecificAuthorizationInfo(serviceSpecificAuthorizationInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceSpecificAuthorizationInfoDocumentApi.CreateServiceSpecificAuthorizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceSpecificAuthorizationInfo`: ServiceSpecificAuthorizationInfo
    fmt.Fprintf(os.Stdout, "Response from `ServiceSpecificAuthorizationInfoDocumentApi.CreateServiceSpecificAuthorizationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**serviceType** | [**ServiceType**](.md) | Service Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceSpecificAuthorizationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceSpecificAuthorizationInfo** | [**ServiceSpecificAuthorizationInfo**](ServiceSpecificAuthorizationInfo.md) |  | 

### Return type

[**ServiceSpecificAuthorizationInfo**](ServiceSpecificAuthorizationInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSpecificAuthorizationInfo

> ServiceSpecificAuthorizationInfo GetServiceSpecificAuthorizationInfo(ctx, ueId, serviceType).Execute()

Retrieve Service Specific Authorization Info

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
    ueId := "ueId_example" // string | 
    serviceType := *openapiclient.NewServiceType() // ServiceType | Service Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceSpecificAuthorizationInfoDocumentApi.GetServiceSpecificAuthorizationInfo(context.Background(), ueId, serviceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceSpecificAuthorizationInfoDocumentApi.GetServiceSpecificAuthorizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceSpecificAuthorizationInfo`: ServiceSpecificAuthorizationInfo
    fmt.Fprintf(os.Stdout, "Response from `ServiceSpecificAuthorizationInfoDocumentApi.GetServiceSpecificAuthorizationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**serviceType** | [**ServiceType**](.md) | Service Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSpecificAuthorizationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceSpecificAuthorizationInfo**](ServiceSpecificAuthorizationInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyServiceSpecificAuthorizationInfo

> PatchResult ModifyServiceSpecificAuthorizationInfo(ctx, ueId, serviceType).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify Service Specific Authorization Info

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
    ueId := "ueId_example" // string | 
    serviceType := *openapiclient.NewServiceType() // ServiceType | Service Type
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceSpecificAuthorizationInfoDocumentApi.ModifyServiceSpecificAuthorizationInfo(context.Background(), ueId, serviceType).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceSpecificAuthorizationInfoDocumentApi.ModifyServiceSpecificAuthorizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyServiceSpecificAuthorizationInfo`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `ServiceSpecificAuthorizationInfoDocumentApi.ModifyServiceSpecificAuthorizationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**serviceType** | [**ServiceType**](.md) | Service Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyServiceSpecificAuthorizationInfoRequest struct via the builder pattern


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


## RemoveServiceSpecificAuthorizationInfo

> RemoveServiceSpecificAuthorizationInfo(ctx, ueId, serviceType).Execute()

Delete Service Specific Authorization Info

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
    ueId := "ueId_example" // string | 
    serviceType := *openapiclient.NewServiceType() // ServiceType | Service Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceSpecificAuthorizationInfoDocumentApi.RemoveServiceSpecificAuthorizationInfo(context.Background(), ueId, serviceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceSpecificAuthorizationInfoDocumentApi.RemoveServiceSpecificAuthorizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**serviceType** | [**ServiceType**](.md) | Service Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServiceSpecificAuthorizationInfoRequest struct via the builder pattern


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

