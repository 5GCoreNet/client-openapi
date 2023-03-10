# \OperatorSpecificDataDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOperatorSpecificData**](OperatorSpecificDataDocumentApi.md#DeleteOperatorSpecificData) | **Delete** /policy-data/ues/{ueId}/operator-specific-data | When the feature OSDResource_Create_Delete is supported, delete OperatorSpecificData resource
[**ReadOperatorSpecificData**](OperatorSpecificDataDocumentApi.md#ReadOperatorSpecificData) | **Get** /policy-data/ues/{ueId}/operator-specific-data | Retrieve the operator specific policy data of an UE
[**ReplaceOperatorSpecificData**](OperatorSpecificDataDocumentApi.md#ReplaceOperatorSpecificData) | **Put** /policy-data/ues/{ueId}/operator-specific-data | Create or modify the operator specific policy data of a UE
[**UpdateOperatorSpecificData**](OperatorSpecificDataDocumentApi.md#UpdateOperatorSpecificData) | **Patch** /policy-data/ues/{ueId}/operator-specific-data | Modify the operator specific policy data of a UE



## DeleteOperatorSpecificData

> DeleteOperatorSpecificData(ctx, ueId).Execute()

When the feature OSDResource_Create_Delete is supported, delete OperatorSpecificData resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Policy_Data"
)

func main() {
    ueId := "ueId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorSpecificDataDocumentApi.DeleteOperatorSpecificData(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorSpecificDataDocumentApi.DeleteOperatorSpecificData``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOperatorSpecificDataRequest struct via the builder pattern


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


## ReadOperatorSpecificData

> map[string]OperatorSpecificDataContainer ReadOperatorSpecificData(ctx, ueId).Fields(fields).SuppFeat(suppFeat).Execute()

Retrieve the operator specific policy data of an UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Policy_Data"
)

func main() {
    ueId := "ueId_example" // string | UE Id
    fields := []string{"Inner_example"} // []string | attributes to be retrieved (optional)
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorSpecificDataDocumentApi.ReadOperatorSpecificData(context.Background(), ueId).Fields(fields).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorSpecificDataDocumentApi.ReadOperatorSpecificData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOperatorSpecificData`: map[string]OperatorSpecificDataContainer
    fmt.Fprintf(os.Stdout, "Response from `OperatorSpecificDataDocumentApi.ReadOperatorSpecificData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOperatorSpecificDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | attributes to be retrieved | 
 **suppFeat** | **string** | Supported Features | 

### Return type

[**map[string]OperatorSpecificDataContainer**](OperatorSpecificDataContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOperatorSpecificData

> map[string]OperatorSpecificDataContainer ReplaceOperatorSpecificData(ctx, ueId).RequestBody(requestBody).Execute()

Create or modify the operator specific policy data of a UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Policy_Data"
)

func main() {
    ueId := "ueId_example" // string | UE Id
    requestBody := map[string]OperatorSpecificDataContainer{"key": *openapiclient.NewOperatorSpecificDataContainer("DataType_example", openapiclient.OperatorSpecificDataContainer_value{Array: new(Array)})} // map[string]OperatorSpecificDataContainer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorSpecificDataDocumentApi.ReplaceOperatorSpecificData(context.Background(), ueId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorSpecificDataDocumentApi.ReplaceOperatorSpecificData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceOperatorSpecificData`: map[string]OperatorSpecificDataContainer
    fmt.Fprintf(os.Stdout, "Response from `OperatorSpecificDataDocumentApi.ReplaceOperatorSpecificData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOperatorSpecificDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**map[string]OperatorSpecificDataContainer**](OperatorSpecificDataContainer.md) |  | 

### Return type

[**map[string]OperatorSpecificDataContainer**](OperatorSpecificDataContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOperatorSpecificData

> PatchResult UpdateOperatorSpecificData(ctx, ueId).PatchItem(patchItem).Execute()

Modify the operator specific policy data of a UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Policy_Data"
)

func main() {
    ueId := "ueId_example" // string | UE Id
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorSpecificDataDocumentApi.UpdateOperatorSpecificData(context.Background(), ueId).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorSpecificDataDocumentApi.UpdateOperatorSpecificData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOperatorSpecificData`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `OperatorSpecificDataDocumentApi.UpdateOperatorSpecificData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOperatorSpecificDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 

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

