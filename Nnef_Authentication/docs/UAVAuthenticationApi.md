# \UAVAuthenticationApi

All URIs are relative to *https://example.com/nnef-authentication/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UavAuthenticationsPost**](UAVAuthenticationApi.md#UavAuthenticationsPost) | **Post** /uav-authentications | UAV authentication



## UavAuthenticationsPost

> UAVAuthResponse UavAuthenticationsPost(ctx).UAVAuthInfo(uAVAuthInfo).Execute()

UAV authentication

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnef_Authentication"
)

func main() {
    uAVAuthInfo := *openapiclient.NewUAVAuthInfo("Gpsi_example", "ServiceLevelId_example", *openapiclient.NewNFType()) // UAVAuthInfo | UAV authentication

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UAVAuthenticationApi.UavAuthenticationsPost(context.Background()).UAVAuthInfo(uAVAuthInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UAVAuthenticationApi.UavAuthenticationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UavAuthenticationsPost`: UAVAuthResponse
    fmt.Fprintf(os.Stdout, "Response from `UAVAuthenticationApi.UavAuthenticationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUavAuthenticationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uAVAuthInfo** | [**UAVAuthInfo**](UAVAuthInfo.md) | UAV authentication | 

### Return type

[**UAVAuthResponse**](UAVAuthResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

