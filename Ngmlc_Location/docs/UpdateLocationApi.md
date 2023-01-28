# \UpdateLocationApi

All URIs are relative to *https://example.com/ngmlc-loc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateLocation**](UpdateLocationApi.md#UpdateLocation) | **Post** /location-update | update UE location information



## UpdateLocation

> UpdateLocation(ctx).LocUpdateData(locUpdateData).Execute()

update UE location information

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
    locUpdateData := *openapiclient.NewLocUpdateData(*openapiclient.NewLocationRequestType(), *openapiclient.NewGeographicArea(*openapiclient.NewSupportedGADShapes(), *openapiclient.NewGeographicalCoordinates(float64(123), float64(123)), float32(123), *openapiclient.NewUncertaintyEllipse(float32(123), float32(123), int32(123)), int32(123), []openapiclient.GeographicalCoordinates{*openapiclient.NewGeographicalCoordinates(float64(123), float64(123))}, float64(123), float32(123), int32(123), float32(123), int32(123), int32(123)), int32(123), *openapiclient.NewAccuracyFulfilmentIndicator(), *openapiclient.NewLcsQosClass()) // LocUpdateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateLocationApi.UpdateLocation(context.Background()).LocUpdateData(locUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateLocationApi.UpdateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locUpdateData** | [**LocUpdateData**](LocUpdateData.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

