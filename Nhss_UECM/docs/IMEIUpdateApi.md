# \IMEIUpdateApi

All URIs are relative to *https://example.com/nhss-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IMEIUpdate**](IMEIUpdateApi.md#IMEIUpdate) | **Post** /imei-update | IMEI Update



## IMEIUpdate

> ImeiUpdateResponse IMEIUpdate(ctx).ImeiUpdateInfo(imeiUpdateInfo).Execute()

IMEI Update

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_UECM"
)

func main() {
    imeiUpdateInfo := openapiclient.ImeiUpdateInfo{Interface{}: new(interface{})} // ImeiUpdateInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IMEIUpdateApi.IMEIUpdate(context.Background()).ImeiUpdateInfo(imeiUpdateInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IMEIUpdateApi.IMEIUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IMEIUpdate`: ImeiUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `IMEIUpdateApi.IMEIUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIMEIUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imeiUpdateInfo** | [**ImeiUpdateInfo**](ImeiUpdateInfo.md) |  | 

### Return type

[**ImeiUpdateResponse**](ImeiUpdateResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

