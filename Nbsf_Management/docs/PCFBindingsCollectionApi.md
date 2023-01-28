# \PCFBindingsCollectionApi

All URIs are relative to *https://example.com/nbsf-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePCFBinding**](PCFBindingsCollectionApi.md#CreatePCFBinding) | **Post** /pcfBindings | Create a new Individual PCF for a PDU Session binding information
[**GetPCFBindings**](PCFBindingsCollectionApi.md#GetPCFBindings) | **Get** /pcfBindings | Read PCF for a PDU Session Bindings information



## CreatePCFBinding

> PcfBinding CreatePCFBinding(ctx).PcfBinding(pcfBinding).Execute()

Create a new Individual PCF for a PDU Session binding information

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
    pcfBinding := *openapiclient.NewPcfBinding("Dnn_example", *openapiclient.NewSnssai(int32(123))) // PcfBinding | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PCFBindingsCollectionApi.CreatePCFBinding(context.Background()).PcfBinding(pcfBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PCFBindingsCollectionApi.CreatePCFBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePCFBinding`: PcfBinding
    fmt.Fprintf(os.Stdout, "Response from `PCFBindingsCollectionApi.CreatePCFBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePCFBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pcfBinding** | [**PcfBinding**](PcfBinding.md) |  | 

### Return type

[**PcfBinding**](PcfBinding.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPCFBindings

> PcfBinding GetPCFBindings(ctx).Ipv4Addr(ipv4Addr).Ipv6Prefix(ipv6Prefix).MacAddr48(macAddr48).Dnn(dnn).Supi(supi).Gpsi(gpsi).Snssai(snssai).IpDomain(ipDomain).SuppFeat(suppFeat).Execute()

Read PCF for a PDU Session Bindings information

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
    ipv4Addr := "ipv4Addr_example" // string | The IPv4 Address of the served UE. (optional)
    ipv6Prefix := "ipv6Prefix_example" // Ipv6Prefix | The IPv6 Address of the served UE. The NF service consumer shall append '/128' to the  IPv6 address in the attribute value. E.g. '2001:db8:85a3::8a2e:370:7334/128'.  (optional)
    macAddr48 := "macAddr48_example" // string | The MAC Address of the served UE. (optional)
    dnn := "dnn_example" // string | DNN. (optional)
    supi := "supi_example" // string | Subscription Permanent Identifier. (optional)
    gpsi := "gpsi_example" // string | Generic Public Subscription Identifier (optional)
    snssai := map[string][]openapiclient.Snssai{ ... } // Snssai | The identification of slice. (optional)
    ipDomain := "ipDomain_example" // string | The IPv4 address domain identifier. (optional)
    suppFeat := "suppFeat_example" // string | To filter irrelevant responses related to unsupported features. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PCFBindingsCollectionApi.GetPCFBindings(context.Background()).Ipv4Addr(ipv4Addr).Ipv6Prefix(ipv6Prefix).MacAddr48(macAddr48).Dnn(dnn).Supi(supi).Gpsi(gpsi).Snssai(snssai).IpDomain(ipDomain).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PCFBindingsCollectionApi.GetPCFBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPCFBindings`: PcfBinding
    fmt.Fprintf(os.Stdout, "Response from `PCFBindingsCollectionApi.GetPCFBindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPCFBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipv4Addr** | **string** | The IPv4 Address of the served UE. | 
 **ipv6Prefix** | **Ipv6Prefix** | The IPv6 Address of the served UE. The NF service consumer shall append &#39;/128&#39; to the  IPv6 address in the attribute value. E.g. &#39;2001:db8:85a3::8a2e:370:7334/128&#39;.  | 
 **macAddr48** | **string** | The MAC Address of the served UE. | 
 **dnn** | **string** | DNN. | 
 **supi** | **string** | Subscription Permanent Identifier. | 
 **gpsi** | **string** | Generic Public Subscription Identifier | 
 **snssai** | [**Snssai**](Snssai.md) | The identification of slice. | 
 **ipDomain** | **string** | The IPv4 address domain identifier. | 
 **suppFeat** | **string** | To filter irrelevant responses related to unsupported features. | 

### Return type

[**PcfBinding**](PcfBinding.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

