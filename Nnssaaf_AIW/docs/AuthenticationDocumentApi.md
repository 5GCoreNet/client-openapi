# \AuthenticationDocumentApi

All URIs are relative to *https://example.com/nnssaaf-aiw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmAuthentication**](AuthenticationDocumentApi.md#ConfirmAuthentication) | **Put** /authentications/{authCtxId} | Confirm the authentication result



## ConfirmAuthentication

> AuthConfirmationResponse ConfirmAuthentication(ctx, authCtxId).AuthConfirmationData(authConfirmationData).Execute()

Confirm the authentication result

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnssaaf_AIW"
)

func main() {
    authCtxId := "authCtxId_example" // string | 
    authConfirmationData := *openapiclient.NewAuthConfirmationData("Supi_example", NullableString(123)) // AuthConfirmationData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationDocumentApi.ConfirmAuthentication(context.Background(), authCtxId).AuthConfirmationData(authConfirmationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationDocumentApi.ConfirmAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmAuthentication`: AuthConfirmationResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationDocumentApi.ConfirmAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCtxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authConfirmationData** | [**AuthConfirmationData**](AuthConfirmationData.md) |  | 

### Return type

[**AuthConfirmationResponse**](AuthConfirmationResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

