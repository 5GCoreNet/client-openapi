# \IndividualPCFForAPDUSessionBindingDocumentApi

All URIs are relative to *https://example.com/nbsf-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateIndPCFBinding**](IndividualPCFForAPDUSessionBindingDocumentApi.md#UpdateIndPCFBinding) | **Patch** /pcfBindings/{bindingId} | Update an existing Individual PCF for a PDU Session Binding information



## UpdateIndPCFBinding

> PcfBinding UpdateIndPCFBinding(ctx, bindingId).PcfBindingPatch(pcfBindingPatch).Execute()

Update an existing Individual PCF for a PDU Session Binding information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nbsf_Management"
)

func main() {
    bindingId := "bindingId_example" // string | Represents the individual PCF for a PDU Session Binding.
    pcfBindingPatch := *openapiclient.NewPcfBindingPatch() // PcfBindingPatch | Parameters to update the existing PCF for a PDU Session binding.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPCFForAPDUSessionBindingDocumentApi.UpdateIndPCFBinding(context.Background(), bindingId).PcfBindingPatch(pcfBindingPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPCFForAPDUSessionBindingDocumentApi.UpdateIndPCFBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndPCFBinding`: PcfBinding
    fmt.Fprintf(os.Stdout, "Response from `IndividualPCFForAPDUSessionBindingDocumentApi.UpdateIndPCFBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bindingId** | **string** | Represents the individual PCF for a PDU Session Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndPCFBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pcfBindingPatch** | [**PcfBindingPatch**](PcfBindingPatch.md) | Parameters to update the existing PCF for a PDU Session binding. | 

### Return type

[**PcfBinding**](PcfBinding.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

