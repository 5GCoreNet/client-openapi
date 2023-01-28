# \IndividualBroadcastMBSSessionContextDocumentApi

All URIs are relative to *https://example.com/namf-mbs-bc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContextDelete**](IndividualBroadcastMBSSessionContextDocumentApi.md#ContextDelete) | **Delete** /mbs-contexts/{mbsContextRef} | Namf_MBSBroadcast ContextDelete service Operation
[**ContextUpdate**](IndividualBroadcastMBSSessionContextDocumentApi.md#ContextUpdate) | **Post** /mbs-contexts/{mbsContextRef}/update | Namf_MBSBroadcast ContextUpdate service Operation



## ContextDelete

> ContextDelete(ctx, mbsContextRef).Execute()

Namf_MBSBroadcast ContextDelete service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_MBSBroadcast"
)

func main() {
    mbsContextRef := "mbsContextRef_example" // string | Unique ID of the broadcast MSB session context to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBroadcastMBSSessionContextDocumentApi.ContextDelete(context.Background(), mbsContextRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBroadcastMBSSessionContextDocumentApi.ContextDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsContextRef** | **string** | Unique ID of the broadcast MSB session context to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiContextDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContextUpdate

> ContextUpdateRspData ContextUpdate(ctx, mbsContextRef).JsonData(jsonData).BinaryDataN2Information(binaryDataN2Information).Execute()

Namf_MBSBroadcast ContextUpdate service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_MBSBroadcast"
)

func main() {
    mbsContextRef := "mbsContextRef_example" // string | Unique ID of the broadcast MSB session context to be updated
    jsonData := *openapiclient.NewContextUpdateReqData() // ContextUpdateReqData |  (optional)
    binaryDataN2Information := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBroadcastMBSSessionContextDocumentApi.ContextUpdate(context.Background(), mbsContextRef).JsonData(jsonData).BinaryDataN2Information(binaryDataN2Information).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBroadcastMBSSessionContextDocumentApi.ContextUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContextUpdate`: ContextUpdateRspData
    fmt.Fprintf(os.Stdout, "Response from `IndividualBroadcastMBSSessionContextDocumentApi.ContextUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsContextRef** | **string** | Unique ID of the broadcast MSB session context to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiContextUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**ContextUpdateReqData**](ContextUpdateReqData.md) |  | 
 **binaryDataN2Information** | ***os.File** |  | 

### Return type

[**ContextUpdateRspData**](ContextUpdateRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

