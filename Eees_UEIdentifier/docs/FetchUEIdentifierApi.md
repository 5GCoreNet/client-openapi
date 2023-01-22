# \FetchUEIdentifierApi

All URIs are relative to *https://example.com/eees-ueidentifier/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUEId**](FetchUEIdentifierApi.md#FetchUEId) | **Post** /fetch | Fetch the identifier of an UE.



## FetchUEId

> string FetchUEId(ctx).UserInformation(userInformation).Execute()

Fetch the identifier of an UE.

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
    userInformation := *openapiclient.NewUserInformation("EasId_example", openapiclient.IpAddr{Interface{}: new(interface{})}) // UserInformation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FetchUEIdentifierApi.FetchUEId(context.Background()).UserInformation(userInformation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FetchUEIdentifierApi.FetchUEId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUEId`: string
    fmt.Fprintf(os.Stdout, "Response from `FetchUEIdentifierApi.FetchUEId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUEIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userInformation** | [**UserInformation**](UserInformation.md) |  | 

### Return type

**string**

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

