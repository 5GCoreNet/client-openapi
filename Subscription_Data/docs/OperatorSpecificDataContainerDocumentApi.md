# \OperatorSpecificDataContainerDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOperSpecData**](OperatorSpecificDataContainerDocumentApi.md#CreateOperSpecData) | **Put** /subscription-data/{ueId}/operator-specific-data | To create an operator-specific data resource of a UE
[**DeleteOperSpecData**](OperatorSpecificDataContainerDocumentApi.md#DeleteOperSpecData) | **Delete** /subscription-data/{ueId}/operator-specific-data | To remove an operator-specific data resource of a UE
[**ModifyOperSpecData**](OperatorSpecificDataContainerDocumentApi.md#ModifyOperSpecData) | **Patch** /subscription-data/{ueId}/operator-specific-data | To modify operator specific data of a UE
[**QueryOperSpecData**](OperatorSpecificDataContainerDocumentApi.md#QueryOperSpecData) | **Get** /subscription-data/{ueId}/operator-specific-data | Retrieves the operator specific data of a UE



## CreateOperSpecData

> map[string]OperatorSpecificDataContainer CreateOperSpecData(ctx, ueId).RequestBody(requestBody).SupportedFeatures(supportedFeatures).Execute()

To create an operator-specific data resource of a UE

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
    requestBody := map[string]OperatorSpecificDataContainer{"key": *openapiclient.NewOperatorSpecificDataContainer("DataType_example", openapiclient.OperatorSpecificDataContainer_value{Array: new(Array)})} // map[string]OperatorSpecificDataContainer | 
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorSpecificDataContainerDocumentApi.CreateOperSpecData(context.Background(), ueId).RequestBody(requestBody).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorSpecificDataContainerDocumentApi.CreateOperSpecData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOperSpecData`: map[string]OperatorSpecificDataContainer
    fmt.Fprintf(os.Stdout, "Response from `OperatorSpecificDataContainerDocumentApi.CreateOperSpecData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperSpecDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**map[string]OperatorSpecificDataContainer**](OperatorSpecificDataContainer.md) |  | 
 **supportedFeatures** | **string** | Supported Features | 

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


## DeleteOperSpecData

> DeleteOperSpecData(ctx, ueId).Execute()

To remove an operator-specific data resource of a UE

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorSpecificDataContainerDocumentApi.DeleteOperSpecData(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorSpecificDataContainerDocumentApi.DeleteOperSpecData``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOperSpecDataRequest struct via the builder pattern


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


## ModifyOperSpecData

> PatchResult ModifyOperSpecData(ctx, ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

To modify operator specific data of a UE

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
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorSpecificDataContainerDocumentApi.ModifyOperSpecData(context.Background(), ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorSpecificDataContainerDocumentApi.ModifyOperSpecData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyOperSpecData`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `OperatorSpecificDataContainerDocumentApi.ModifyOperSpecData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOperSpecDataRequest struct via the builder pattern


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


## QueryOperSpecData

> map[string]OperatorSpecificDataContainer QueryOperSpecData(ctx, ueId).Fields(fields).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Retrieves the operator specific data of a UE

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
    fields := []string{"Inner_example"} // []string | attributes to be retrieved (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorSpecificDataContainerDocumentApi.QueryOperSpecData(context.Background(), ueId).Fields(fields).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorSpecificDataContainerDocumentApi.QueryOperSpecData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryOperSpecData`: map[string]OperatorSpecificDataContainer
    fmt.Fprintf(os.Stdout, "Response from `OperatorSpecificDataContainerDocumentApi.QueryOperSpecData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryOperSpecDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | attributes to be retrieved | 
 **supportedFeatures** | **string** | Supported Features | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**map[string]OperatorSpecificDataContainer**](OperatorSpecificDataContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

