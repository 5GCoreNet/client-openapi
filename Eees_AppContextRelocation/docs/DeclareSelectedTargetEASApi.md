# \DeclareSelectedTargetEASApi

All URIs are relative to *https://example.com/eees-appctxtreloc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Declare**](DeclareSelectedTargetEASApi.md#Declare) | **Post** /declare | Informs about the selected target EAS and provides the associated information.



## Declare

> Declare(ctx).AcrDecReq(acrDecReq).Execute()

Informs about the selected target EAS and provides the associated information.

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
    acrDecReq := *openapiclient.NewAcrDecReq("UeId_example", "TEasId_example", openapiclient.EndPoint{Interface{}: new(interface{})}) // AcrDecReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeclareSelectedTargetEASApi.Declare(context.Background()).AcrDecReq(acrDecReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeclareSelectedTargetEASApi.Declare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeclareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acrDecReq** | [**AcrDecReq**](AcrDecReq.md) |  | 

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

