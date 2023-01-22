# \RequestCipheringKeyDataApi

All URIs are relative to *https://example.com/nlmf-broadcast/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CipheringKeyData**](RequestCipheringKeyDataApi.md#CipheringKeyData) | **Post** /cipher-key-data | Request ciphering key data



## CipheringKeyData

> CipherResponseData CipheringKeyData(ctx).CipherRequestData(cipherRequestData).Execute()

Request ciphering key data

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
    cipherRequestData := *openapiclient.NewCipherRequestData("AmfCallBackURI_example") // CipherRequestData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestCipheringKeyDataApi.CipheringKeyData(context.Background()).CipherRequestData(cipherRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestCipheringKeyDataApi.CipheringKeyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CipheringKeyData`: CipherResponseData
    fmt.Fprintf(os.Stdout, "Response from `RequestCipheringKeyDataApi.CipheringKeyData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCipheringKeyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cipherRequestData** | [**CipherRequestData**](CipherRequestData.md) |  | 

### Return type

[**CipherResponseData**](CipherResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

