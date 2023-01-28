# \ParameterExchangeApi

All URIs are relative to *https://example.com/n32c-handshake/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostExchangeParams**](ParameterExchangeApi.md#PostExchangeParams) | **Post** /exchange-params | Parameter Exchange



## PostExchangeParams

> SecParamExchRspData PostExchangeParams(ctx).SecParamExchReqData(secParamExchReqData).Execute()

Parameter Exchange

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
    secParamExchReqData := *openapiclient.NewSecParamExchReqData("N32fContextId_example") // SecParamExchReqData | Custom operation for parameter exchange

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParameterExchangeApi.PostExchangeParams(context.Background()).SecParamExchReqData(secParamExchReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterExchangeApi.PostExchangeParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostExchangeParams`: SecParamExchRspData
    fmt.Fprintf(os.Stdout, "Response from `ParameterExchangeApi.PostExchangeParams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExchangeParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secParamExchReqData** | [**SecParamExchReqData**](SecParamExchReqData.md) | Custom operation for parameter exchange | 

### Return type

[**SecParamExchRspData**](SecParamExchRspData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

