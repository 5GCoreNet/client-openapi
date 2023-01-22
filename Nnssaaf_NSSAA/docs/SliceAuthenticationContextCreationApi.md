# \SliceAuthenticationContextCreationApi

All URIs are relative to *https://example.com/nnssaaf-nssaa/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSliceAuthenticationContext**](SliceAuthenticationContextCreationApi.md#CreateSliceAuthenticationContext) | **Post** /slice-authentications | Create slice authentication context



## CreateSliceAuthenticationContext

> SliceAuthContext CreateSliceAuthenticationContext(ctx).SliceAuthInfo(sliceAuthInfo).Execute()

Create slice authentication context

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
    sliceAuthInfo := *openapiclient.NewSliceAuthInfo("Gpsi_example", *openapiclient.NewSnssai(int32(123)), NullableString(123)) // SliceAuthInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SliceAuthenticationContextCreationApi.CreateSliceAuthenticationContext(context.Background()).SliceAuthInfo(sliceAuthInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SliceAuthenticationContextCreationApi.CreateSliceAuthenticationContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSliceAuthenticationContext`: SliceAuthContext
    fmt.Fprintf(os.Stdout, "Response from `SliceAuthenticationContextCreationApi.CreateSliceAuthenticationContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSliceAuthenticationContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sliceAuthInfo** | [**SliceAuthInfo**](SliceAuthInfo.md) |  | 

### Return type

[**SliceAuthContext**](SliceAuthContext.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

