# \BroadcastMBSSessionContextsCollectionCollectionApi

All URIs are relative to *https://example.com/namf-mbs-bc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContextCreate**](BroadcastMBSSessionContextsCollectionCollectionApi.md#ContextCreate) | **Post** /mbs-contexts | Namf_MBSBroadcast ContextCreate service Operation



## ContextCreate

> ContextCreateRspData ContextCreate(ctx).JsonData(jsonData).BinaryDataN2Information(binaryDataN2Information).Execute()

Namf_MBSBroadcast ContextCreate service Operation

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
    jsonData := openapiclient.ContextCreateReqData{Interface{}: new(interface{})} // ContextCreateReqData |  (optional)
    binaryDataN2Information := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastMBSSessionContextsCollectionCollectionApi.ContextCreate(context.Background()).JsonData(jsonData).BinaryDataN2Information(binaryDataN2Information).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastMBSSessionContextsCollectionCollectionApi.ContextCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContextCreate`: ContextCreateRspData
    fmt.Fprintf(os.Stdout, "Response from `BroadcastMBSSessionContextsCollectionCollectionApi.ContextCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContextCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonData** | [**ContextCreateReqData**](ContextCreateReqData.md) |  | 
 **binaryDataN2Information** | ***os.File** |  | 

### Return type

[**ContextCreateRspData**](ContextCreateRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

