# \FetchAnUELocationInformationApi

All URIs are relative to *https://example.com/eees-uelocation/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUELocation**](FetchAnUELocationInformationApi.md#FetchUELocation) | **Post** /fetch | Fetch an UE location information.



## FetchUELocation

> LocationResponse FetchUELocation(ctx).LocationRequest(locationRequest).Execute()

Fetch an UE location information.

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
    locationRequest := *openapiclient.NewLocationRequest("UeId_example") // LocationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FetchAnUELocationInformationApi.FetchUELocation(context.Background()).LocationRequest(locationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FetchAnUELocationInformationApi.FetchUELocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUELocation`: LocationResponse
    fmt.Fprintf(os.Stdout, "Response from `FetchAnUELocationInformationApi.FetchUELocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUELocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationRequest** | [**LocationRequest**](LocationRequest.md) |  | 

### Return type

[**LocationResponse**](LocationResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

