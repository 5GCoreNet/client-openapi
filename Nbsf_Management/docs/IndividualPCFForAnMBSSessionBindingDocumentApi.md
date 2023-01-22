# \IndividualPCFForAnMBSSessionBindingDocumentApi

All URIs are relative to *https://example.com/nbsf-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndPCFMbsBinding**](IndividualPCFForAnMBSSessionBindingDocumentApi.md#DeleteIndPCFMbsBinding) | **Delete** /pcf-mbs-bindings/{bindingId} | Request the deletion of an existing Individual PCF for an MBS Session Binding.
[**ModifyIndPCFMbsBinding**](IndividualPCFForAnMBSSessionBindingDocumentApi.md#ModifyIndPCFMbsBinding) | **Patch** /pcf-mbs-bindings/{bindingId} | Request the modification of an existing Individual PCF for an MBS Session Binding resource.



## DeleteIndPCFMbsBinding

> DeleteIndPCFMbsBinding(ctx, bindingId).Execute()

Request the deletion of an existing Individual PCF for an MBS Session Binding.

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
    bindingId := "bindingId_example" // string | Represents the identifier of the Individual PCF for an MBS Session Binding resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPCFForAnMBSSessionBindingDocumentApi.DeleteIndPCFMbsBinding(context.Background(), bindingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPCFForAnMBSSessionBindingDocumentApi.DeleteIndPCFMbsBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bindingId** | **string** | Represents the identifier of the Individual PCF for an MBS Session Binding resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndPCFMbsBindingRequest struct via the builder pattern


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


## ModifyIndPCFMbsBinding

> PcfForUeBinding ModifyIndPCFMbsBinding(ctx, bindingId).PcfMbsBindingPatch(pcfMbsBindingPatch).Execute()

Request the modification of an existing Individual PCF for an MBS Session Binding resource.

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
    bindingId := "bindingId_example" // string | Represents the identifier of the Individual PCF for an MBS Session Binding resource. 
    pcfMbsBindingPatch := *openapiclient.NewPcfMbsBindingPatch() // PcfMbsBindingPatch | Parameters to request the modification of the PCF for an MBS Session Binding.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPCFForAnMBSSessionBindingDocumentApi.ModifyIndPCFMbsBinding(context.Background(), bindingId).PcfMbsBindingPatch(pcfMbsBindingPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPCFForAnMBSSessionBindingDocumentApi.ModifyIndPCFMbsBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndPCFMbsBinding`: PcfForUeBinding
    fmt.Fprintf(os.Stdout, "Response from `IndividualPCFForAnMBSSessionBindingDocumentApi.ModifyIndPCFMbsBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bindingId** | **string** | Represents the identifier of the Individual PCF for an MBS Session Binding resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndPCFMbsBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pcfMbsBindingPatch** | [**PcfMbsBindingPatch**](PcfMbsBindingPatch.md) | Parameters to request the modification of the PCF for an MBS Session Binding. | 

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

