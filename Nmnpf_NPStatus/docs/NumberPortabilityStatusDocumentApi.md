# \NumberPortabilityStatusDocumentApi

All URIs are relative to *https://example.com/nmnpf-npstatus/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNumberPortabilityStatus**](NumberPortabilityStatusDocumentApi.md#GetNumberPortabilityStatus) | **Get** /{gpsi} | Retrieves the Number Portability status of the UE



## GetNumberPortabilityStatus

> NpStatusInfo GetNumberPortabilityStatus(ctx, gpsi).Execute()

Retrieves the Number Portability status of the UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmnpf_NPStatus"
)

func main() {
    gpsi := "gpsi_example" // string | GPSI of the UE

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NumberPortabilityStatusDocumentApi.GetNumberPortabilityStatus(context.Background(), gpsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NumberPortabilityStatusDocumentApi.GetNumberPortabilityStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNumberPortabilityStatus`: NpStatusInfo
    fmt.Fprintf(os.Stdout, "Response from `NumberPortabilityStatusDocumentApi.GetNumberPortabilityStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gpsi** | **string** | GPSI of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNumberPortabilityStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NpStatusInfo**](NpStatusInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

