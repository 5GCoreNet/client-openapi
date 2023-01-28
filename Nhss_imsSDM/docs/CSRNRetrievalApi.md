# \CSRNRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCsrn**](CSRNRetrievalApi.md#GetCsrn) | **Get** /{imsUeId}/access-data/cs-domain/csrn | Retrieve the routeing number in CS domain



## GetCsrn

> Csrn GetCsrn(ctx, imsUeId).PrePaging(prePaging).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).CallReferenceInfo(callReferenceInfo).Execute()

Retrieve the routeing number in CS domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsSDM"
)

func main() {
    imsUeId := "imsUeId_example" // string | IMS Public Identity
    prePaging := true // bool | Indicates pre-paging support (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)
    callReferenceInfo := map[string][]openapiclient.CallReferenceInfo{ ... } // CallReferenceInfo | Indicates Call-Reference Number and AS-Number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSRNRetrievalApi.GetCsrn(context.Background(), imsUeId).PrePaging(prePaging).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).CallReferenceInfo(callReferenceInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRNRetrievalApi.GetCsrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCsrn`: Csrn
    fmt.Fprintf(os.Stdout, "Response from `CSRNRetrievalApi.GetCsrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Public Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCsrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prePaging** | **bool** | Indicates pre-paging support | 
 **supportedFeatures** | **string** | Supported Features | 
 **privateIdentity** | **string** | IMS Private Identity | 
 **callReferenceInfo** | [**CallReferenceInfo**](CallReferenceInfo.md) | Indicates Call-Reference Number and AS-Number | 

### Return type

[**Csrn**](Csrn.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

