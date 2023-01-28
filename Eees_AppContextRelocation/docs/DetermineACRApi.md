# \DetermineACRApi

All URIs are relative to *https://example.com/eees-appctxtreloc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Determine**](DetermineACRApi.md#Determine) | **Post** /determine | Request ACR determination.



## Determine

> Determine(ctx).AcrDetermReq(acrDetermReq).Execute()

Request ACR determination.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_AppContextRelocation"
)

func main() {
    acrDetermReq := *openapiclient.NewAcrDetermReq("RequestorId_example", openapiclient.EndPoint{Interface{}: new(interface{})}) // AcrDetermReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DetermineACRApi.Determine(context.Background()).AcrDetermReq(acrDetermReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetermineACRApi.Determine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDetermineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acrDetermReq** | [**AcrDetermReq**](AcrDetermReq.md) |  | 

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

