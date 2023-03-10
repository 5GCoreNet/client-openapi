# \GenerateGBAAuthenticationVectorsApi

All URIs are relative to *https://example.com/nudm-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateGbaAv**](GenerateGBAAuthenticationVectorsApi.md#GenerateGbaAv) | **Post** /{supi}/gba-security-information/generate-av | Generate authentication data for the UE in GBA domain



## GenerateGbaAv

> GbaAuthenticationInfoResult GenerateGbaAv(ctx, supi).GbaAuthenticationInfoRequest(gbaAuthenticationInfoRequest).Execute()

Generate authentication data for the UE in GBA domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudm_UEAU"
)

func main() {
    supi := "supi_example" // string | SUPI of the user
    gbaAuthenticationInfoRequest := *openapiclient.NewGbaAuthenticationInfoRequest(*openapiclient.NewGbaAuthType()) // GbaAuthenticationInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerateGBAAuthenticationVectorsApi.GenerateGbaAv(context.Background(), supi).GbaAuthenticationInfoRequest(gbaAuthenticationInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerateGBAAuthenticationVectorsApi.GenerateGbaAv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateGbaAv`: GbaAuthenticationInfoResult
    fmt.Fprintf(os.Stdout, "Response from `GenerateGBAAuthenticationVectorsApi.GenerateGbaAv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | SUPI of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateGbaAvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gbaAuthenticationInfoRequest** | [**GbaAuthenticationInfoRequest**](GbaAuthenticationInfoRequest.md) |  | 

### Return type

[**GbaAuthenticationInfoResult**](GbaAuthenticationInfoResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

