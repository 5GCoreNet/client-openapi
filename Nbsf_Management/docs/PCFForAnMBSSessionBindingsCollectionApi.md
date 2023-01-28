# \PCFForAnMBSSessionBindingsCollectionApi

All URIs are relative to *https://example.com/nbsf-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePCFMbsBinding**](PCFForAnMBSSessionBindingsCollectionApi.md#CreatePCFMbsBinding) | **Post** /pcf-mbs-bindings | Create a new Individual PCF for an MBS Session binding.
[**GetPCFMbsBinding**](PCFForAnMBSSessionBindingsCollectionApi.md#GetPCFMbsBinding) | **Get** /pcf-mbs-bindings | Retrieve an existing PCF for an MBS Session binding.



## CreatePCFMbsBinding

> PcfMbsBinding CreatePCFMbsBinding(ctx).PcfMbsBinding(pcfMbsBinding).Execute()

Create a new Individual PCF for an MBS Session binding.

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
    pcfMbsBinding := *openapiclient.NewPcfMbsBinding(*openapiclient.NewMbsSessionId()) // PcfMbsBinding | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PCFForAnMBSSessionBindingsCollectionApi.CreatePCFMbsBinding(context.Background()).PcfMbsBinding(pcfMbsBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PCFForAnMBSSessionBindingsCollectionApi.CreatePCFMbsBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePCFMbsBinding`: PcfMbsBinding
    fmt.Fprintf(os.Stdout, "Response from `PCFForAnMBSSessionBindingsCollectionApi.CreatePCFMbsBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePCFMbsBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pcfMbsBinding** | [**PcfMbsBinding**](PcfMbsBinding.md) |  | 

### Return type

[**PcfMbsBinding**](PcfMbsBinding.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPCFMbsBinding

> PcfMbsBinding GetPCFMbsBinding(ctx).MbsSessionId(mbsSessionId).SuppFeat(suppFeat).Execute()

Retrieve an existing PCF for an MBS Session binding.

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
    mbsSessionId := map[string][]openapiclient.MbsSessionId{ ... } // MbsSessionId | Contains the identifier of the MBS Session to which the requested MBS Session binding is related. 
    suppFeat := "suppFeat_example" // string | Contains the list of features supported by the NF service consumer and used to filter irrelevant responses related to unsupported features.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PCFForAnMBSSessionBindingsCollectionApi.GetPCFMbsBinding(context.Background()).MbsSessionId(mbsSessionId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PCFForAnMBSSessionBindingsCollectionApi.GetPCFMbsBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPCFMbsBinding`: PcfMbsBinding
    fmt.Fprintf(os.Stdout, "Response from `PCFForAnMBSSessionBindingsCollectionApi.GetPCFMbsBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPCFMbsBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mbsSessionId** | [**MbsSessionId**](MbsSessionId.md) | Contains the identifier of the MBS Session to which the requested MBS Session binding is related.  | 
 **suppFeat** | **string** | Contains the list of features supported by the NF service consumer and used to filter irrelevant responses related to unsupported features.  | 

### Return type

[**PcfMbsBinding**](PcfMbsBinding.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

