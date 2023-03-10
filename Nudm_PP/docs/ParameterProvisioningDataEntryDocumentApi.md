# \ParameterProvisioningDataEntryDocumentApi

All URIs are relative to *https://example.com/nudm-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePPDataEntry**](ParameterProvisioningDataEntryDocumentApi.md#CreatePPDataEntry) | **Put** /{ueId}/pp-data-store/{afInstanceId} | Create a Provisioning Parameter Data Entry
[**DeletePPDataEntry**](ParameterProvisioningDataEntryDocumentApi.md#DeletePPDataEntry) | **Delete** /{ueId}/pp-data-store/{afInstanceId} | Delete a Provisioning Parameter Data Entry
[**GetPPDataEntry**](ParameterProvisioningDataEntryDocumentApi.md#GetPPDataEntry) | **Get** /{ueId}/pp-data-store/{afInstanceId} | get Parameter Provisioning Data Entry



## CreatePPDataEntry

> PpDataEntry CreatePPDataEntry(ctx, ueId, afInstanceId).PpDataEntry(ppDataEntry).Execute()

Create a Provisioning Parameter Data Entry

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudm_PP"
)

func main() {
    ueId := *openapiclient.NewGetPPDataEntryUeIdParameter() // GetPPDataEntryUeIdParameter | Identifier of the UE
    afInstanceId := "afInstanceId_example" // string | Application Function Instance Identifier
    ppDataEntry := *openapiclient.NewPpDataEntry() // PpDataEntry | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParameterProvisioningDataEntryDocumentApi.CreatePPDataEntry(context.Background(), ueId, afInstanceId).PpDataEntry(ppDataEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterProvisioningDataEntryDocumentApi.CreatePPDataEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePPDataEntry`: PpDataEntry
    fmt.Fprintf(os.Stdout, "Response from `ParameterProvisioningDataEntryDocumentApi.CreatePPDataEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**GetPPDataEntryUeIdParameter**](.md) | Identifier of the UE | 
**afInstanceId** | **string** | Application Function Instance Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePPDataEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ppDataEntry** | [**PpDataEntry**](PpDataEntry.md) |  | 

### Return type

[**PpDataEntry**](PpDataEntry.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePPDataEntry

> DeletePPDataEntry(ctx, ueId, afInstanceId).Execute()

Delete a Provisioning Parameter Data Entry

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudm_PP"
)

func main() {
    ueId := *openapiclient.NewGetPPDataEntryUeIdParameter() // GetPPDataEntryUeIdParameter | Identifier of the UE
    afInstanceId := "afInstanceId_example" // string | Application Function Instance Identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParameterProvisioningDataEntryDocumentApi.DeletePPDataEntry(context.Background(), ueId, afInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterProvisioningDataEntryDocumentApi.DeletePPDataEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**GetPPDataEntryUeIdParameter**](.md) | Identifier of the UE | 
**afInstanceId** | **string** | Application Function Instance Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePPDataEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPPDataEntry

> PpDataEntry GetPPDataEntry(ctx, ueId, afInstanceId).SupportedFeatures(supportedFeatures).Execute()

get Parameter Provisioning Data Entry

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudm_PP"
)

func main() {
    ueId := *openapiclient.NewGetPPDataEntryUeIdParameter() // GetPPDataEntryUeIdParameter | Identifier of the UE
    afInstanceId := "afInstanceId_example" // string | Application Function Instance Identifier
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParameterProvisioningDataEntryDocumentApi.GetPPDataEntry(context.Background(), ueId, afInstanceId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterProvisioningDataEntryDocumentApi.GetPPDataEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPPDataEntry`: PpDataEntry
    fmt.Fprintf(os.Stdout, "Response from `ParameterProvisioningDataEntryDocumentApi.GetPPDataEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**GetPPDataEntryUeIdParameter**](.md) | Identifier of the UE | 
**afInstanceId** | **string** | Application Function Instance Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPPDataEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PpDataEntry**](PpDataEntry.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

