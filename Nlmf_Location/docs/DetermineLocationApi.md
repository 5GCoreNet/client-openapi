# \DetermineLocationApi

All URIs are relative to *https://example.com/nlmf-loc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DetermineLocation**](DetermineLocationApi.md#DetermineLocation) | **Post** /determine-location | Determine Location of an UE



## DetermineLocation

> LocationData DetermineLocation(ctx).InputData(inputData).Execute()

Determine Location of an UE

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
    inputData := *openapiclient.NewInputData() // InputData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DetermineLocationApi.DetermineLocation(context.Background()).InputData(inputData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetermineLocationApi.DetermineLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetermineLocation`: LocationData
    fmt.Fprintf(os.Stdout, "Response from `DetermineLocationApi.DetermineLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDetermineLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputData** | [**InputData**](InputData.md) |  | 

### Return type

[**LocationData**](LocationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

