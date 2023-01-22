# \SlicePolicyControlDataDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadSlicePolicyControlData**](SlicePolicyControlDataDocumentApi.md#ReadSlicePolicyControlData) | **Get** /policy-data/slice-control-data/{snssai} | Retrieves a network Slice specific policy control data resource
[**UpdateSlicePolicyControlData**](SlicePolicyControlDataDocumentApi.md#UpdateSlicePolicyControlData) | **Patch** /policy-data/slice-control-data/{snssai} | Modify a network Slice specific policy control data resource



## ReadSlicePolicyControlData

> SlicePolicyData ReadSlicePolicyControlData(ctx, snssai).SuppFeat(suppFeat).Execute()

Retrieves a network Slice specific policy control data resource

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
    snssai := map[string][]openapiclient.Snssai{ ... } // Snssai | 
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlicePolicyControlDataDocumentApi.ReadSlicePolicyControlData(context.Background(), snssai).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlicePolicyControlDataDocumentApi.ReadSlicePolicyControlData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSlicePolicyControlData`: SlicePolicyData
    fmt.Fprintf(os.Stdout, "Response from `SlicePolicyControlDataDocumentApi.ReadSlicePolicyControlData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snssai** | [**Snssai**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSlicePolicyControlDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppFeat** | **string** | Supported Features | 

### Return type

[**SlicePolicyData**](SlicePolicyData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlicePolicyControlData

> SlicePolicyData UpdateSlicePolicyControlData(ctx, snssai).SlicePolicyDataPatch(slicePolicyDataPatch).Execute()

Modify a network Slice specific policy control data resource

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
    snssai := map[string][]openapiclient.Snssai{ ... } // Snssai | 
    slicePolicyDataPatch := openapiclient.SlicePolicyDataPatch{Interface{}: new(interface{})} // SlicePolicyDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlicePolicyControlDataDocumentApi.UpdateSlicePolicyControlData(context.Background(), snssai).SlicePolicyDataPatch(slicePolicyDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlicePolicyControlDataDocumentApi.UpdateSlicePolicyControlData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSlicePolicyControlData`: SlicePolicyData
    fmt.Fprintf(os.Stdout, "Response from `SlicePolicyControlDataDocumentApi.UpdateSlicePolicyControlData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snssai** | [**Snssai**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSlicePolicyControlDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slicePolicyDataPatch** | [**SlicePolicyDataPatch**](SlicePolicyDataPatch.md) |  | 

### Return type

[**SlicePolicyData**](SlicePolicyData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

