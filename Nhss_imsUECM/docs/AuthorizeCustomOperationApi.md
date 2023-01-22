# \AuthorizeCustomOperationApi

All URIs are relative to *https://example.com/nhss-ims-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](AuthorizeCustomOperationApi.md#Authorize) | **Post** /{impu}/authorize | Authorize IMS Identities to register in the network or establish multimedia sessions and return CSCF location if it is stored 



## Authorize

> AuthorizationResponse Authorize(ctx, impu).AuthorizationRequest(authorizationRequest).Execute()

Authorize IMS Identities to register in the network or establish multimedia sessions and return CSCF location if it is stored 

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
    impu := "impu_example" // string | Public identity of the user.
    authorizationRequest := *openapiclient.NewAuthorizationRequest(*openapiclient.NewAuthorizationType()) // AuthorizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeCustomOperationApi.Authorize(context.Background(), impu).AuthorizationRequest(authorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeCustomOperationApi.Authorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Authorize`: AuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeCustomOperationApi.Authorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**impu** | **string** | Public identity of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationRequest** | [**AuthorizationRequest**](AuthorizationRequest.md) |  | 

### Return type

[**AuthorizationResponse**](AuthorizationResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

