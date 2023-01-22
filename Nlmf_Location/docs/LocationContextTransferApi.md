# \LocationContextTransferApi

All URIs are relative to *https://example.com/nlmf-loc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationContextTransfer**](LocationContextTransferApi.md#LocationContextTransfer) | **Post** /location-context-transfer | transfer context information for periodic or triggered location



## LocationContextTransfer

> LocationContextTransfer(ctx).LocContextData(locContextData).Execute()

transfer context information for periodic or triggered location

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
    locContextData := *openapiclient.NewLocContextData("AmfId_example", *openapiclient.NewLdrType(), "HgmlcCallBackURI_example", "LdrReference_example", *openapiclient.NewEventReportMessage(*openapiclient.NewEventClass(), *openapiclient.NewRefToBinaryData("ContentId_example"))) // LocContextData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationContextTransferApi.LocationContextTransfer(context.Background()).LocContextData(locContextData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationContextTransferApi.LocationContextTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationContextTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locContextData** | [**LocContextData**](LocContextData.md) |  | 

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

