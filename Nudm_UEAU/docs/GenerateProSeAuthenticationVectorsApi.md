# \GenerateProSeAuthenticationVectorsApi

All URIs are relative to *https://example.com/nudm-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateProseAV**](GenerateProSeAuthenticationVectorsApi.md#GenerateProseAV) | **Post** /{supiOrSuci}/prose-security-information/generate-av | Generate authentication data for ProSe



## GenerateProseAV

> ProSeAuthenticationInfoResult GenerateProseAV(ctx, supiOrSuci).ProSeAuthenticationInfoRequest(proSeAuthenticationInfoRequest).Execute()

Generate authentication data for ProSe

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
    supiOrSuci := "supiOrSuci_example" // string | SUPI or SUCI of the user
    proSeAuthenticationInfoRequest := *openapiclient.NewProSeAuthenticationInfoRequest(int32(123)) // ProSeAuthenticationInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerateProSeAuthenticationVectorsApi.GenerateProseAV(context.Background(), supiOrSuci).ProSeAuthenticationInfoRequest(proSeAuthenticationInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerateProSeAuthenticationVectorsApi.GenerateProseAV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateProseAV`: ProSeAuthenticationInfoResult
    fmt.Fprintf(os.Stdout, "Response from `GenerateProSeAuthenticationVectorsApi.GenerateProseAV`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supiOrSuci** | **string** | SUPI or SUCI of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateProseAVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **proSeAuthenticationInfoRequest** | [**ProSeAuthenticationInfoRequest**](ProSeAuthenticationInfoRequest.md) |  | 

### Return type

[**ProSeAuthenticationInfoResult**](ProSeAuthenticationInfoResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

