# \IndividualPCFForAUEBindingDocumentApi

All URIs are relative to *https://example.com/nbsf-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndPCFforUEBinding**](IndividualPCFForAUEBindingDocumentApi.md#DeleteIndPCFforUEBinding) | **Delete** /pcf-ue-bindings/{bindingId} | Delete an existing Individual PCF for a UE Binding information
[**UpdateIndPCFforUEBinding**](IndividualPCFForAUEBindingDocumentApi.md#UpdateIndPCFforUEBinding) | **Patch** /pcf-ue-bindings/{bindingId} | Update an existing Individual PCF for a UE Binding information



## DeleteIndPCFforUEBinding

> DeleteIndPCFforUEBinding(ctx, bindingId).Execute()

Delete an existing Individual PCF for a UE Binding information

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
    bindingId := "bindingId_example" // string | Represents the individual PCF for a UE Binding.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPCFForAUEBindingDocumentApi.DeleteIndPCFforUEBinding(context.Background(), bindingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPCFForAUEBindingDocumentApi.DeleteIndPCFforUEBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bindingId** | **string** | Represents the individual PCF for a UE Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndPCFforUEBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndPCFforUEBinding

> PcfForUeBinding UpdateIndPCFforUEBinding(ctx, bindingId).PcfForUeBindingPatch(pcfForUeBindingPatch).Execute()

Update an existing Individual PCF for a UE Binding information

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
    bindingId := "bindingId_example" // string | Represents the individual PCF for a UE Binding.
    pcfForUeBindingPatch := *openapiclient.NewPcfForUeBindingPatch() // PcfForUeBindingPatch | Parameters to update the existing PCF for a UE binding.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPCFForAUEBindingDocumentApi.UpdateIndPCFforUEBinding(context.Background(), bindingId).PcfForUeBindingPatch(pcfForUeBindingPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPCFForAUEBindingDocumentApi.UpdateIndPCFforUEBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndPCFforUEBinding`: PcfForUeBinding
    fmt.Fprintf(os.Stdout, "Response from `IndividualPCFForAUEBindingDocumentApi.UpdateIndPCFforUEBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bindingId** | **string** | Represents the individual PCF for a UE Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndPCFforUEBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pcfForUeBindingPatch** | [**PcfForUeBindingPatch**](PcfForUeBindingPatch.md) | Parameters to update the existing PCF for a UE binding. | 

### Return type

[**PcfForUeBinding**](PcfForUeBinding.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

