# \N2MessagesHandlerCustomOperationApi

All URIs are relative to *https://example.com/namf-mbs-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**N2MessageTransfer**](N2MessagesHandlerCustomOperationApi.md#N2MessageTransfer) | **Post** /n2-messages/transfer | Namf_MBSCommunication N2 Message Transfer service Operation



## N2MessageTransfer

> MbsN2MessageTransferRspData N2MessageTransfer(ctx).JsonData(jsonData).BinaryDataN2Information(binaryDataN2Information).Execute()

Namf_MBSCommunication N2 Message Transfer service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_MBSCommunication"
)

func main() {
    jsonData := *openapiclient.NewMbsN2MessageTransferReqData(*openapiclient.NewMbsSessionId(), *openapiclient.NewN2MbsSmInfo(*openapiclient.NewMbsNgapIeType(), *openapiclient.NewRefToBinaryData("ContentId_example"))) // MbsN2MessageTransferReqData |  (optional)
    binaryDataN2Information := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N2MessagesHandlerCustomOperationApi.N2MessageTransfer(context.Background()).JsonData(jsonData).BinaryDataN2Information(binaryDataN2Information).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N2MessagesHandlerCustomOperationApi.N2MessageTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `N2MessageTransfer`: MbsN2MessageTransferRspData
    fmt.Fprintf(os.Stdout, "Response from `N2MessagesHandlerCustomOperationApi.N2MessageTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiN2MessageTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonData** | [**MbsN2MessageTransferReqData**](MbsN2MessageTransferReqData.md) |  | 
 **binaryDataN2Information** | ***os.File** |  | 

### Return type

[**MbsN2MessageTransferRspData**](MbsN2MessageTransferRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

