# \CancelLocationApi

All URIs are relative to *https://example.com/nlmf-loc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelLocation**](CancelLocationApi.md#CancelLocation) | **Post** /cancel-location | request cancellation of periodic or triggered location



## CancelLocation

> CancelLocation(ctx).CancelLocData(cancelLocData).Execute()

request cancellation of periodic or triggered location

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nlmf_Location"
)

func main() {
    cancelLocData := *openapiclient.NewCancelLocData("HgmlcCallBackURI_example", "LdrReference_example") // CancelLocData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CancelLocationApi.CancelLocation(context.Background()).CancelLocData(cancelLocData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CancelLocationApi.CancelLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelLocData** | [**CancelLocData**](CancelLocData.md) |  | 

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

