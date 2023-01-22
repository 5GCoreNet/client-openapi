# \TSCStreamRetrievalApi

All URIs are relative to *https://example.com/ss-nra/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTscStream**](TSCStreamRetrievalApi.md#GetTscStream) | **Get** /tsc-streams | Retrieval of TSC stream data.



## GetTscStream

> []TscStreamData GetTscStream(ctx).ValStreamIds(valStreamIds).Execute()

Retrieval of TSC stream data.

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
    valStreamIds := []string{"Inner_example"} // []string | Retrieval of TSC Stream data, identified by the VAL Stream ID(s). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TSCStreamRetrievalApi.GetTscStream(context.Background()).ValStreamIds(valStreamIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TSCStreamRetrievalApi.GetTscStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTscStream`: []TscStreamData
    fmt.Fprintf(os.Stdout, "Response from `TSCStreamRetrievalApi.GetTscStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTscStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **valStreamIds** | **[]string** | Retrieval of TSC Stream data, identified by the VAL Stream ID(s). | 

### Return type

[**[]TscStreamData**](TscStreamData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

