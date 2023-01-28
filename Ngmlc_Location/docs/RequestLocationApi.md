# \RequestLocationApi

All URIs are relative to *https://example.com/ngmlc-loc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestLocation**](RequestLocationApi.md#RequestLocation) | **Post** /provide-location | Request Location of an UE



## RequestLocation

> LocationData RequestLocation(ctx).InputData(inputData).Execute()

Request Location of an UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ngmlc_Location"
)

func main() {
    inputData := *openapiclient.NewInputData(*openapiclient.NewExternalClientType()) // InputData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestLocationApi.RequestLocation(context.Background()).InputData(inputData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestLocationApi.RequestLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestLocation`: LocationData
    fmt.Fprintf(os.Stdout, "Response from `RequestLocationApi.RequestLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputData** | [**InputData**](InputData.md) |  | 

### Return type

[**LocationData**](LocationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

