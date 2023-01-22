# \N32FErrorReportApi

All URIs are relative to *https://example.com/n32c-handshake/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostN32fError**](N32FErrorReportApi.md#PostN32fError) | **Post** /n32f-error | N32-f Error Reporting Procedure



## PostN32fError

> PostN32fError(ctx).N32fErrorInfo(n32fErrorInfo).Execute()

N32-f Error Reporting Procedure

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
    n32fErrorInfo := *openapiclient.NewN32fErrorInfo("N32fMessageId_example", *openapiclient.NewN32fErrorType()) // N32fErrorInfo | Custom operation for n32-f error reporting procedure

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N32FErrorReportApi.PostN32fError(context.Background()).N32fErrorInfo(n32fErrorInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N32FErrorReportApi.PostN32fError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostN32fErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n32fErrorInfo** | [**N32fErrorInfo**](N32fErrorInfo.md) | Custom operation for n32-f error reporting procedure | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

