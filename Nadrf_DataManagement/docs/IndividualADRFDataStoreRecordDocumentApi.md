# \IndividualADRFDataStoreRecordDocumentApi

All URIs are relative to *https://example.com/nadrf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteADRFDataStoreRecord**](IndividualADRFDataStoreRecordDocumentApi.md#DeleteADRFDataStoreRecord) | **Delete** /data-store-records/{storeTransId} | Delete an existing Individual ADRF Data Store Record.



## DeleteADRFDataStoreRecord

> DeleteADRFDataStoreRecord(ctx, storeTransId).Execute()

Delete an existing Individual ADRF Data Store Record.

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
    storeTransId := "storeTransId_example" // string | String identifying a Data Store Record in ADRF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualADRFDataStoreRecordDocumentApi.DeleteADRFDataStoreRecord(context.Background(), storeTransId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualADRFDataStoreRecordDocumentApi.DeleteADRFDataStoreRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeTransId** | **string** | String identifying a Data Store Record in ADRF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteADRFDataStoreRecordRequest struct via the builder pattern


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

