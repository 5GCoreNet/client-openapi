# \AuthenticationsCollectionsApi

All URIs are relative to *https://example.com/nnssaaf-aiw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationContext**](AuthenticationsCollectionsApi.md#CreateAuthenticationContext) | **Post** /authentications | Create authentication context



## CreateAuthenticationContext

> AuthContext CreateAuthenticationContext(ctx).AuthInfo(authInfo).Execute()

Create authentication context

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
    authInfo := *openapiclient.NewAuthInfo("Supi_example") // AuthInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationsCollectionsApi.CreateAuthenticationContext(context.Background()).AuthInfo(authInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationsCollectionsApi.CreateAuthenticationContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthenticationContext`: AuthContext
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationsCollectionsApi.CreateAuthenticationContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authInfo** | [**AuthInfo**](AuthInfo.md) |  | 

### Return type

[**AuthContext**](AuthContext.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

