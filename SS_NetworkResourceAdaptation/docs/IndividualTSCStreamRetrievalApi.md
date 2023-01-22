# \IndividualTSCStreamRetrievalApi

All URIs are relative to *https://example.com/ss-nra/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTscStreamData**](IndividualTSCStreamRetrievalApi.md#GetTscStreamData) | **Get** /tsc-streams/{valStreamId} | Reads an existing Individual TSC stream data information



## GetTscStreamData

> TscStreamData GetTscStreamData(ctx, valStreamId).Execute()

Reads an existing Individual TSC stream data information

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
    valStreamId := "valStreamId_example" // string | The VAL Stream ID identifies the TSC stream.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTSCStreamRetrievalApi.GetTscStreamData(context.Background(), valStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTSCStreamRetrievalApi.GetTscStreamData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTscStreamData`: TscStreamData
    fmt.Fprintf(os.Stdout, "Response from `IndividualTSCStreamRetrievalApi.GetTscStreamData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valStreamId** | **string** | The VAL Stream ID identifies the TSC stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTscStreamDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TscStreamData**](TscStreamData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

