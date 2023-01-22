# \CompleteStoredSearchDocumentApi

All URIs are relative to *https://example.com/nnrf-disc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveCompleteSearch**](CompleteStoredSearchDocumentApi.md#RetrieveCompleteSearch) | **Get** /searches/{searchId}/complete | 



## RetrieveCompleteSearch

> StoredSearchResult RetrieveCompleteSearch(ctx, searchId).AcceptEncoding(acceptEncoding).Execute()



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
    searchId := "searchId_example" // string | Id of a stored search
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompleteStoredSearchDocumentApi.RetrieveCompleteSearch(context.Background(), searchId).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompleteStoredSearchDocumentApi.RetrieveCompleteSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCompleteSearch`: StoredSearchResult
    fmt.Fprintf(os.Stdout, "Response from `CompleteStoredSearchDocumentApi.RetrieveCompleteSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchId** | **string** | Id of a stored search | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCompleteSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 

### Return type

[**StoredSearchResult**](StoredSearchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

