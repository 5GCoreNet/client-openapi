# \DNSContextsCollectionApi

All URIs are relative to *https://example.com/neasdf-dnscontext/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnsContext**](DNSContextsCollectionApi.md#CreateDnsContext) | **Post** /dns-contexts | Create



## CreateDnsContext

> DnsContextCreatedData CreateDnsContext(ctx).DnsContextCreateData(dnsContextCreateData).Execute()

Create

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Neasdf_DNSContext"
)

func main() {
    dnsContextCreateData := *openapiclient.NewDnsContextCreateData("Dnn_example", *openapiclient.NewSnssai(int32(123)), map[string]DnsRule{"key": *openapiclient.NewDnsRule(map[string]Action{"key": *openapiclient.NewAction(*openapiclient.NewApplyAction())})}) // DnsContextCreateData | representation of the DNS context to be created in the EASDF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSContextsCollectionApi.CreateDnsContext(context.Background()).DnsContextCreateData(dnsContextCreateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSContextsCollectionApi.CreateDnsContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDnsContext`: DnsContextCreatedData
    fmt.Fprintf(os.Stdout, "Response from `DNSContextsCollectionApi.CreateDnsContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDnsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsContextCreateData** | [**DnsContextCreateData**](DnsContextCreateData.md) | representation of the DNS context to be created in the EASDF | 

### Return type

[**DnsContextCreatedData**](DnsContextCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

