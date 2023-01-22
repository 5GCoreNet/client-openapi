# \ProvisionedParameterDataEntryDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePPDataEntry**](ProvisionedParameterDataEntryDocumentApi.md#CreatePPDataEntry) | **Put** /subscription-data/{ueId}/pp-data-store/{afInstanceId} | create a Provisioning Parameter Data Entry
[**DeletePPDataEntry**](ProvisionedParameterDataEntryDocumentApi.md#DeletePPDataEntry) | **Delete** /subscription-data/{ueId}/pp-data-store/{afInstanceId} | Delete a Provisioning Parameter Data Entry
[**GetPPDataEntry**](ProvisionedParameterDataEntryDocumentApi.md#GetPPDataEntry) | **Get** /subscription-data/{ueId}/pp-data-store/{afInstanceId} | get a Parameter Provisioning Data Entry



## CreatePPDataEntry

> PpDataEntry CreatePPDataEntry(ctx, ueId, afInstanceId).PpDataEntry(ppDataEntry).Execute()

create a Provisioning Parameter Data Entry

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
    ueId := *openapiclient.NewGetPPDataEntryUeIdParameter() // GetPPDataEntryUeIdParameter | Identifier of the UE
    afInstanceId := "afInstanceId_example" // string | Application Function Instance Identifier
    ppDataEntry := *openapiclient.NewPpDataEntry() // PpDataEntry | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisionedParameterDataEntryDocumentApi.CreatePPDataEntry(context.Background(), ueId, afInstanceId).PpDataEntry(ppDataEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedParameterDataEntryDocumentApi.CreatePPDataEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePPDataEntry`: PpDataEntry
    fmt.Fprintf(os.Stdout, "Response from `ProvisionedParameterDataEntryDocumentApi.CreatePPDataEntry`: %v\n", resp)
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

No authorization required

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
    openapiclient "./openapi"
)

func main() {
    ueId := *openapiclient.NewGetPPDataEntryUeIdParameter() // GetPPDataEntryUeIdParameter | Identifier of the UE
    afInstanceId := "afInstanceId_example" // string | Application Function Instance Identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisionedParameterDataEntryDocumentApi.DeletePPDataEntry(context.Background(), ueId, afInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedParameterDataEntryDocumentApi.DeletePPDataEntry``: %v\n", err)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPPDataEntry

> PpDataEntry GetPPDataEntry(ctx, ueId, afInstanceId).SupportedFeatures(supportedFeatures).Execute()

get a Parameter Provisioning Data Entry

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
    ueId := *openapiclient.NewGetPPDataEntryUeIdParameter() // GetPPDataEntryUeIdParameter | Identifier of the UE
    afInstanceId := "afInstanceId_example" // string | Application Function Instance Identifier
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisionedParameterDataEntryDocumentApi.GetPPDataEntry(context.Background(), ueId, afInstanceId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedParameterDataEntryDocumentApi.GetPPDataEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPPDataEntry`: PpDataEntry
    fmt.Fprintf(os.Stdout, "Response from `ProvisionedParameterDataEntryDocumentApi.GetPPDataEntry`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

