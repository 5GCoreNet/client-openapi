# \N32FContextTerminateApi

All URIs are relative to *https://example.com/n32c-handshake/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostN32fTerminate**](N32FContextTerminateApi.md#PostN32fTerminate) | **Post** /n32f-terminate | N32-f Context Terminate



## PostN32fTerminate

> N32fContextInfo PostN32fTerminate(ctx).N32fContextInfo(n32fContextInfo).Execute()

N32-f Context Terminate

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/N32_Handshake"
)

func main() {
    n32fContextInfo := *openapiclient.NewN32fContextInfo("N32fContextId_example") // N32fContextInfo | Custom operation for n32-f context termination

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N32FContextTerminateApi.PostN32fTerminate(context.Background()).N32fContextInfo(n32fContextInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N32FContextTerminateApi.PostN32fTerminate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostN32fTerminate`: N32fContextInfo
    fmt.Fprintf(os.Stdout, "Response from `N32FContextTerminateApi.PostN32fTerminate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostN32fTerminateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n32fContextInfo** | [**N32fContextInfo**](N32fContextInfo.md) | Custom operation for n32-f context termination | 

### Return type

[**N32fContextInfo**](N32fContextInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

