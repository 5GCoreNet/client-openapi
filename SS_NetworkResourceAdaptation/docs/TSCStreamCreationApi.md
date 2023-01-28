# \TSCStreamCreationApi

All URIs are relative to *https://example.com/ss-nra/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutTscStream**](TSCStreamCreationApi.md#PutTscStream) | **Put** /tsc-streams/{valStreamId} | Create a TSC stream identified by a VAL stream identifier.



## PutTscStream

> TscStreamData PutTscStream(ctx, valStreamId).TscStreamData(tscStreamData).Execute()

Create a TSC stream identified by a VAL stream identifier.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_NetworkResourceAdaptation"
)

func main() {
    valStreamId := "valStreamId_example" // string | VAL stream identifier
    tscStreamData := *openapiclient.NewTscStreamData(*openapiclient.NewStreamSpecification("SrcMacAddr_example", "DstMacAddr_example"), *openapiclient.NewTrafficSpecInformation(int32(123), int32(123), int32(123), int32(123), int32(123))) // TscStreamData | TSC stream creation request data from the VAL server to the NRM server.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TSCStreamCreationApi.PutTscStream(context.Background(), valStreamId).TscStreamData(tscStreamData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TSCStreamCreationApi.PutTscStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutTscStream`: TscStreamData
    fmt.Fprintf(os.Stdout, "Response from `TSCStreamCreationApi.PutTscStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valStreamId** | **string** | VAL stream identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTscStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tscStreamData** | [**TscStreamData**](TscStreamData.md) | TSC stream creation request data from the VAL server to the NRM server. | 

### Return type

[**TscStreamData**](TscStreamData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

