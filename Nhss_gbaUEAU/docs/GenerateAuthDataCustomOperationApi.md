# \GenerateAuthDataCustomOperationApi

All URIs are relative to *https://example.com/nhss-gba-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAuthData**](GenerateAuthDataCustomOperationApi.md#GenerateAuthData) | **Post** /{ueId}/security-information/generate-auth-data | Generate GBA authentication data for the UE



## GenerateAuthData

> AuthenticationInfoResult GenerateAuthData(ctx, ueId).AuthenticationInfoRequest(authenticationInfoRequest).Execute()

Generate GBA authentication data for the UE

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
    ueId := *openapiclient.NewUeId() // UeId | UE identity of the user
    authenticationInfoRequest := *openapiclient.NewAuthenticationInfoRequest() // AuthenticationInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerateAuthDataCustomOperationApi.GenerateAuthData(context.Background(), ueId).AuthenticationInfoRequest(authenticationInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerateAuthDataCustomOperationApi.GenerateAuthData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAuthData`: AuthenticationInfoResult
    fmt.Fprintf(os.Stdout, "Response from `GenerateAuthDataCustomOperationApi.GenerateAuthData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**UeId**](.md) | UE identity of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAuthDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationInfoRequest** | [**AuthenticationInfoRequest**](AuthenticationInfoRequest.md) |  | 

### Return type

[**AuthenticationInfoResult**](AuthenticationInfoResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

