# \RequestACRUpdateApi

All URIs are relative to *https://example.com/eees-acrstatus-update/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestACRUpdate**](RequestACRUpdateApi.md#RequestACRUpdate) | **Post** /request-acrupdate | Request to update the information related to ACR (e.g. indicate the status of ACT, update the notification target address).



## RequestACRUpdate

> ACRDataStatus RequestACRUpdate(ctx).ACRUpdateData(aCRUpdateData).Execute()

Request to update the information related to ACR (e.g. indicate the status of ACT, update the notification target address).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_ACRStatusUpdate"
)

func main() {
    aCRUpdateData := *openapiclient.NewACRUpdateData("EasId_example") // ACRUpdateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestACRUpdateApi.RequestACRUpdate(context.Background()).ACRUpdateData(aCRUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestACRUpdateApi.RequestACRUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestACRUpdate`: ACRDataStatus
    fmt.Fprintf(os.Stdout, "Response from `RequestACRUpdateApi.RequestACRUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestACRUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aCRUpdateData** | [**ACRUpdateData**](ACRUpdateData.md) |  | 

### Return type

[**ACRDataStatus**](ACRDataStatus.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

