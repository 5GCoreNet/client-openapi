# \RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryPeiInformation**](RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApi.md#QueryPeiInformation) | **Get** /subscription-data/{ueId}/context-data/pei-info | Retrieves the PEI Information of the 5GC/EPC domains



## QueryPeiInformation

> PeiUpdateInfo QueryPeiInformation(ctx, ueId).Execute()

Retrieves the PEI Information of the 5GC/EPC domains

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApi.QueryPeiInformation(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApi.QueryPeiInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryPeiInformation`: PeiUpdateInfo
    fmt.Fprintf(os.Stdout, "Response from `RetrieveTheRoamingInformationOfThe5GCEPCDomainsDocumentApi.QueryPeiInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryPeiInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PeiUpdateInfo**](PeiUpdateInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

