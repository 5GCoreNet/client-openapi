# \ADRFStoredDataApi

All URIs are relative to *https://example.com/nadrf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteADRFData**](ADRFStoredDataApi.md#DeleteADRFData) | **Post** /remove-stored-data-analytics | Remove ADRF data based on data or analytics specification.



## DeleteADRFData

> DeleteADRFData(ctx).NadrfStoredDataSpec(nadrfStoredDataSpec).Execute()

Remove ADRF data based on data or analytics specification.

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
    nadrfStoredDataSpec := openapiclient.NadrfStoredDataSpec{Interface{}: new(interface{})} // NadrfStoredDataSpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ADRFStoredDataApi.DeleteADRFData(context.Background()).NadrfStoredDataSpec(nadrfStoredDataSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ADRFStoredDataApi.DeleteADRFData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteADRFDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nadrfStoredDataSpec** | [**NadrfStoredDataSpec**](NadrfStoredDataSpec.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

