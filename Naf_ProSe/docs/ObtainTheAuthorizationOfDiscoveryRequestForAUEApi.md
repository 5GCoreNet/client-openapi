# \ObtainTheAuthorizationOfDiscoveryRequestForAUEApi

All URIs are relative to *https://example.com/naf-prose/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObtainDiscAuth**](ObtainTheAuthorizationOfDiscoveryRequestForAUEApi.md#ObtainDiscAuth) | **Post** /authorize-discovery | Obtain the authorization of Discovery Request from 5G DDNMF for a UE



## ObtainDiscAuth

> AuthDisResData ObtainDiscAuth(ctx).AuthDisReqData(authDisReqData).Execute()

Obtain the authorization of Discovery Request from 5G DDNMF for a UE

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
    authDisReqData := *openapiclient.NewAuthDisReqData(*openapiclient.NewAuthRequestType()) // AuthDisReqData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObtainTheAuthorizationOfDiscoveryRequestForAUEApi.ObtainDiscAuth(context.Background()).AuthDisReqData(authDisReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObtainTheAuthorizationOfDiscoveryRequestForAUEApi.ObtainDiscAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ObtainDiscAuth`: AuthDisResData
    fmt.Fprintf(os.Stdout, "Response from `ObtainTheAuthorizationOfDiscoveryRequestForAUEApi.ObtainDiscAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObtainDiscAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authDisReqData** | [**AuthDisReqData**](AuthDisReqData.md) |  | 

### Return type

[**AuthDisResData**](AuthDisResData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

