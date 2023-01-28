# \ConfirmSliceAuthenticationApi

All URIs are relative to *https://example.com/nnssaaf-nssaa/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmSliceAuthentication**](ConfirmSliceAuthenticationApi.md#ConfirmSliceAuthentication) | **Put** /slice-authentications/{authCtxId} | Confirm the slice authentication result



## ConfirmSliceAuthentication

> SliceAuthConfirmationResponse ConfirmSliceAuthentication(ctx, authCtxId).SliceAuthConfirmationData(sliceAuthConfirmationData).Execute()

Confirm the slice authentication result

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnssaaf_NSSAA"
)

func main() {
    authCtxId := "authCtxId_example" // string | 
    sliceAuthConfirmationData := *openapiclient.NewSliceAuthConfirmationData("Gpsi_example", *openapiclient.NewSnssai(int32(123)), NullableString(123)) // SliceAuthConfirmationData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfirmSliceAuthenticationApi.ConfirmSliceAuthentication(context.Background(), authCtxId).SliceAuthConfirmationData(sliceAuthConfirmationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfirmSliceAuthenticationApi.ConfirmSliceAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmSliceAuthentication`: SliceAuthConfirmationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfirmSliceAuthenticationApi.ConfirmSliceAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCtxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmSliceAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sliceAuthConfirmationData** | [**SliceAuthConfirmationData**](SliceAuthConfirmationData.md) |  | 

### Return type

[**SliceAuthConfirmationResponse**](SliceAuthConfirmationResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

