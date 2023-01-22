# \UpdateThePEIInformationOfThe5GCEPCDomainsDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdatePeiInformation**](UpdateThePEIInformationOfThe5GCEPCDomainsDocumentApi.md#CreateOrUpdatePeiInformation) | **Put** /subscription-data/{ueId}/context-data/pei-info | Update the PEI Information of the 5GC/EPC domains



## CreateOrUpdatePeiInformation

> PeiUpdateInfo CreateOrUpdatePeiInformation(ctx, ueId).PeiUpdateInfo(peiUpdateInfo).Execute()

Update the PEI Information of the 5GC/EPC domains

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
    ueId := "ueId_example" // string | UE id
    peiUpdateInfo := *openapiclient.NewPeiUpdateInfo("Pei_example") // PeiUpdateInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateThePEIInformationOfThe5GCEPCDomainsDocumentApi.CreateOrUpdatePeiInformation(context.Background(), ueId).PeiUpdateInfo(peiUpdateInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateThePEIInformationOfThe5GCEPCDomainsDocumentApi.CreateOrUpdatePeiInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdatePeiInformation`: PeiUpdateInfo
    fmt.Fprintf(os.Stdout, "Response from `UpdateThePEIInformationOfThe5GCEPCDomainsDocumentApi.CreateOrUpdatePeiInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdatePeiInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **peiUpdateInfo** | [**PeiUpdateInfo**](PeiUpdateInfo.md) |  | 

### Return type

[**PeiUpdateInfo**](PeiUpdateInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

