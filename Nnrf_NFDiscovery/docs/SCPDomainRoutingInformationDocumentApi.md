# \SCPDomainRoutingInformationDocumentApi

All URIs are relative to *https://example.com/nnrf-disc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SCPDomainRoutingInfoGet**](SCPDomainRoutingInformationDocumentApi.md#SCPDomainRoutingInfoGet) | **Get** /scp-domain-routing-info | 



## SCPDomainRoutingInfoGet

> ScpDomainRoutingInformation SCPDomainRoutingInfoGet(ctx).Local(local).AcceptEncoding(acceptEncoding).Execute()



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
    local := true // bool | Indication of local SCP Domain Routing Information (optional) (default to false)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SCPDomainRoutingInformationDocumentApi.SCPDomainRoutingInfoGet(context.Background()).Local(local).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCPDomainRoutingInformationDocumentApi.SCPDomainRoutingInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SCPDomainRoutingInfoGet`: ScpDomainRoutingInformation
    fmt.Fprintf(os.Stdout, "Response from `SCPDomainRoutingInformationDocumentApi.SCPDomainRoutingInfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSCPDomainRoutingInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **local** | **bool** | Indication of local SCP Domain Routing Information | [default to false]
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 

### Return type

[**ScpDomainRoutingInformation**](ScpDomainRoutingInformation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

