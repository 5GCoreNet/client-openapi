# \TelescopicMappingDocumentApi

All URIs are relative to *https://example.com/nsepp-telescopic/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTelescopicMapping**](TelescopicMappingDocumentApi.md#GetTelescopicMapping) | **Get** /mapping | Maps an FQDN to/from a telescopic FQDN



## GetTelescopicMapping

> TelescopicMapping GetTelescopicMapping(ctx).ForeignFqdn(foreignFqdn).TelescopicLabel(telescopicLabel).Execute()

Maps an FQDN to/from a telescopic FQDN

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
    foreignFqdn := "foreignFqdn_example" // string | FQDN of the NF in the foreign PLMN (optional)
    telescopicLabel := "telescopicLabel_example" // string | Telescopic Label (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TelescopicMappingDocumentApi.GetTelescopicMapping(context.Background()).ForeignFqdn(foreignFqdn).TelescopicLabel(telescopicLabel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelescopicMappingDocumentApi.GetTelescopicMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTelescopicMapping`: TelescopicMapping
    fmt.Fprintf(os.Stdout, "Response from `TelescopicMappingDocumentApi.GetTelescopicMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelescopicMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **foreignFqdn** | **string** | FQDN of the NF in the foreign PLMN | 
 **telescopicLabel** | **string** | Telescopic Label | 

### Return type

[**TelescopicMapping**](TelescopicMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

