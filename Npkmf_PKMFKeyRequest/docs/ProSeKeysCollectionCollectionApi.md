# \ProSeKeysCollectionCollectionApi

All URIs are relative to *https://example.com/npkmf-keyerquest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProseKey**](ProSeKeysCollectionCollectionApi.md#ProseKey) | **Post** /prose-keys/request | Request Keying Materials for 5G ProSe



## ProseKey

> ProseKeyRspData ProseKey(ctx).ProseKeyReqData(proseKeyReqData).Execute()

Request Keying Materials for 5G ProSe

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npkmf_PKMFKeyRequest"
)

func main() {
    proseKeyReqData := *openapiclient.NewProseKeyReqData(int32(123), "KnrpFreshness1_example") // ProseKeyReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProSeKeysCollectionCollectionApi.ProseKey(context.Background()).ProseKeyReqData(proseKeyReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProSeKeysCollectionCollectionApi.ProseKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProseKey`: ProseKeyRspData
    fmt.Fprintf(os.Stdout, "Response from `ProSeKeysCollectionCollectionApi.ProseKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProseKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proseKeyReqData** | [**ProseKeyReqData**](ProseKeyReqData.md) |  | 

### Return type

[**ProseKeyRspData**](ProseKeyRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

