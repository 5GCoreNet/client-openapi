# \IndividualDNSContextApi

All URIs are relative to *https://example.com/neasdf-dnscontext/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDnsContext**](IndividualDNSContextApi.md#DeleteDnsContext) | **Delete** /dns-contexts/{dnsContextId} | Delete the DNS Context
[**ReplaceDnsContext**](IndividualDNSContextApi.md#ReplaceDnsContext) | **Put** /dns-contexts/{dnsContextId} | Updates the DNS context (complete replacement)
[**UpdateDnsContext**](IndividualDNSContextApi.md#UpdateDnsContext) | **Patch** /dns-contexts/{dnsContextId} | Updates the DNS context



## DeleteDnsContext

> DeleteDnsContext(ctx, dnsContextId).Execute()

Delete the DNS Context

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
    dnsContextId := "dnsContextId_example" // string | DNS context Identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDNSContextApi.DeleteDnsContext(context.Background(), dnsContextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDNSContextApi.DeleteDnsContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnsContextId** | **string** | DNS context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDnsContextRequest struct via the builder pattern


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


## ReplaceDnsContext

> ReplaceDnsContext(ctx, dnsContextId).DnsContextCreateData(dnsContextCreateData).ContentEncoding(contentEncoding).Execute()

Updates the DNS context (complete replacement)

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
    dnsContextId := "dnsContextId_example" // string | DNS context Identifier
    dnsContextCreateData := *openapiclient.NewDnsContextCreateData("Dnn_example", *openapiclient.NewSnssai(int32(123)), map[string]DnsRule{"key": *openapiclient.NewDnsRule(map[string]Action{"key": *openapiclient.NewAction(*openapiclient.NewApplyAction())})}) // DnsContextCreateData | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDNSContextApi.ReplaceDnsContext(context.Background(), dnsContextId).DnsContextCreateData(dnsContextCreateData).ContentEncoding(contentEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDNSContextApi.ReplaceDnsContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnsContextId** | **string** | DNS context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDnsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsContextCreateData** | [**DnsContextCreateData**](DnsContextCreateData.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDnsContext

> PatchResult UpdateDnsContext(ctx, dnsContextId).PatchItem(patchItem).ContentEncoding(contentEncoding).Execute()

Updates the DNS context

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
    dnsContextId := "dnsContextId_example" // string | DNS context Identifier
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDNSContextApi.UpdateDnsContext(context.Background(), dnsContextId).PatchItem(patchItem).ContentEncoding(contentEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDNSContextApi.UpdateDnsContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDnsContext`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `IndividualDNSContextApi.UpdateDnsContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnsContextId** | **string** | DNS context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDnsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

