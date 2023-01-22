# \DefaultApi

All URIs are relative to *https://example.com/api-invocation-logs/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AefIdLogsPost**](DefaultApi.md#AefIdLogsPost) | **Post** /{aefId}/logs | 



## AefIdLogsPost

> InvocationLog AefIdLogsPost(ctx, aefId).InvocationLog(invocationLog).Execute()





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
    aefId := "aefId_example" // string | Identifier of the API exposing function
    invocationLog := *openapiclient.NewInvocationLog("AefId_example", "ApiInvokerId_example", []openapiclient.Log{*openapiclient.NewLog("ApiId_example", "ApiName_example", "ApiVersion_example", "ResourceName_example", *openapiclient.NewProtocol(), "Result_example")}) // InvocationLog | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AefIdLogsPost(context.Background(), aefId).InvocationLog(invocationLog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AefIdLogsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AefIdLogsPost`: InvocationLog
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AefIdLogsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aefId** | **string** | Identifier of the API exposing function | 

### Other Parameters

Other parameters are passed through a pointer to a apiAefIdLogsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invocationLog** | [**InvocationLog**](InvocationLog.md) |  | 

### Return type

[**InvocationLog**](InvocationLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

