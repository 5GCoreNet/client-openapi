# \UEContextInPGWDataRetrievalApi

All URIs are relative to *https://example.com/nhss-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUeCtxInPgwData**](UEContextInPGWDataRetrievalApi.md#GetUeCtxInPgwData) | **Get** /{ueId}/ue-context-in-pgw-data | Retrieve the UE Context In PGW



## GetUeCtxInPgwData

> UeContextInPgwData GetUeCtxInPgwData(ctx, ueId).Execute()

Retrieve the UE Context In PGW

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_SDM"
)

func main() {
    ueId := "ueId_example" // string | Identifier of the UE

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UEContextInPGWDataRetrievalApi.GetUeCtxInPgwData(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UEContextInPGWDataRetrievalApi.GetUeCtxInPgwData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUeCtxInPgwData`: UeContextInPgwData
    fmt.Fprintf(os.Stdout, "Response from `UEContextInPGWDataRetrievalApi.GetUeCtxInPgwData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUeCtxInPgwDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UeContextInPgwData**](UeContextInPgwData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

