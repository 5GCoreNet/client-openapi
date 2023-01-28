# \DefaultApi

All URIs are relative to *https://example.com/3gpp-ueid/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveUEId**](DefaultApi.md#RetrieveUEId) | **Post** /retrieve | Retrieve AF specific UE ID.



## RetrieveUEId

> UeIdInfo RetrieveUEId(ctx).UeIdReq(ueIdReq).Execute()

Retrieve AF specific UE ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/UEId"
)

func main() {
    ueIdReq := openapiclient.UeIdReq{Interface{}: new(interface{})} // UeIdReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveUEId(context.Background()).UeIdReq(ueIdReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveUEId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUEId`: UeIdInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveUEId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUEIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ueIdReq** | [**UeIdReq**](UeIdReq.md) |  | 

### Return type

[**UeIdInfo**](UeIdInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

