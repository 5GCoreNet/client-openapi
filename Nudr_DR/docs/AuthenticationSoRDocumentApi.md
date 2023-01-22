# \AuthenticationSoRDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationSoR**](AuthenticationSoRDocumentApi.md#CreateAuthenticationSoR) | **Put** /subscription-data/{ueId}/ue-update-confirmation-data/sor-data | To store the SoR acknowledgement information of a UE and ME support of SOR CMCI
[**QueryAuthSoR**](AuthenticationSoRDocumentApi.md#QueryAuthSoR) | **Get** /subscription-data/{ueId}/ue-update-confirmation-data/sor-data | Retrieves the SoR acknowledgement information of a UE and ME support of SOR CMCI
[**UpdateAuthenticationSoR**](AuthenticationSoRDocumentApi.md#UpdateAuthenticationSoR) | **Patch** /subscription-data/{ueId}/ue-update-confirmation-data/sor-data | Updates the ME support of SOR CMCI information of a UE



## CreateAuthenticationSoR

> CreateAuthenticationSoR(ctx, ueId).SorData(sorData).SupportedFeatures(supportedFeatures).Execute()

To store the SoR acknowledgement information of a UE and ME support of SOR CMCI

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    ueId := "ueId_example" // string | UE id
    sorData := *openapiclient.NewSorData(time.Now(), openapiclient.UeUpdateStatus("NOT_SENT")) // SorData | 
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationSoRDocumentApi.CreateAuthenticationSoR(context.Background(), ueId).SorData(sorData).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSoRDocumentApi.CreateAuthenticationSoR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationSoRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sorData** | [**SorData**](SorData.md) |  | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryAuthSoR

> SorData QueryAuthSoR(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

Retrieves the SoR acknowledgement information of a UE and ME support of SOR CMCI

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
    ueId := "ueId_example" // string | UE id
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationSoRDocumentApi.QueryAuthSoR(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSoRDocumentApi.QueryAuthSoR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAuthSoR`: SorData
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationSoRDocumentApi.QueryAuthSoR`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryAuthSoRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**SorData**](SorData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticationSoR

> PatchResult UpdateAuthenticationSoR(ctx, ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Updates the ME support of SOR CMCI information of a UE

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
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationSoRDocumentApi.UpdateAuthenticationSoR(context.Background(), ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSoRDocumentApi.UpdateAuthenticationSoR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthenticationSoR`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationSoRDocumentApi.UpdateAuthenticationSoR`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationSoRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

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

