# \PCFForAUEBindingsCollectionApi

All URIs are relative to *https://example.com/nbsf-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePCFforUEBinding**](PCFForAUEBindingsCollectionApi.md#CreatePCFforUEBinding) | **Post** /pcf-ue-bindings | Create a new Individual PCF for a UE binding information
[**GetPCFForUeBindings**](PCFForAUEBindingsCollectionApi.md#GetPCFForUeBindings) | **Get** /pcf-ue-bindings | Read PCF for a UE Bindings information



## CreatePCFforUEBinding

> PcfForUeBinding CreatePCFforUEBinding(ctx).PcfForUeBinding(pcfForUeBinding).Execute()

Create a new Individual PCF for a UE binding information

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
    pcfForUeBinding := *openapiclient.NewPcfForUeBinding("Supi_example") // PcfForUeBinding | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PCFForAUEBindingsCollectionApi.CreatePCFforUEBinding(context.Background()).PcfForUeBinding(pcfForUeBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PCFForAUEBindingsCollectionApi.CreatePCFforUEBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePCFforUEBinding`: PcfForUeBinding
    fmt.Fprintf(os.Stdout, "Response from `PCFForAUEBindingsCollectionApi.CreatePCFforUEBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePCFforUEBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pcfForUeBinding** | [**PcfForUeBinding**](PcfForUeBinding.md) |  | 

### Return type

[**PcfForUeBinding**](PcfForUeBinding.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPCFForUeBindings

> []PcfForUeBinding GetPCFForUeBindings(ctx).Supi(supi).Gpsi(gpsi).SuppFeat(suppFeat).Execute()

Read PCF for a UE Bindings information

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
    supi := "supi_example" // string | Subscription Permanent Identifier. (optional)
    gpsi := "gpsi_example" // string | Generic Public Subscription Identifier (optional)
    suppFeat := "suppFeat_example" // string | To filter irrelevant responses related to unsupported features. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PCFForAUEBindingsCollectionApi.GetPCFForUeBindings(context.Background()).Supi(supi).Gpsi(gpsi).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PCFForAUEBindingsCollectionApi.GetPCFForUeBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPCFForUeBindings`: []PcfForUeBinding
    fmt.Fprintf(os.Stdout, "Response from `PCFForAUEBindingsCollectionApi.GetPCFForUeBindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPCFForUeBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supi** | **string** | Subscription Permanent Identifier. | 
 **gpsi** | **string** | Generic Public Subscription Identifier | 
 **suppFeat** | **string** | To filter irrelevant responses related to unsupported features. | 

### Return type

[**[]PcfForUeBinding**](PcfForUeBinding.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

