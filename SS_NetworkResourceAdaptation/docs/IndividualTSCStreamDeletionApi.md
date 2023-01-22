# \IndividualTSCStreamDeletionApi

All URIs are relative to *https://example.com/ss-nra/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTscStream**](IndividualTSCStreamDeletionApi.md#DeleteTscStream) | **Delete** /tsc-streams/{valStreamId} | Delete an existing Individual TSC stream



## DeleteTscStream

> DeleteTscStream(ctx, valStreamId).Execute()

Delete an existing Individual TSC stream

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
    resp, r, err := apiClient.IndividualTSCStreamDeletionApi.DeleteTscStream(context.Background(), valStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTSCStreamDeletionApi.DeleteTscStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valStreamId** | **string** | The VAL Stream ID identifies the TSC stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTscStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

