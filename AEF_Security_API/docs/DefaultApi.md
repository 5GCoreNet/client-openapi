# \DefaultApi

All URIs are relative to *https://example.com/aef-security/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAuthenticationPost**](DefaultApi.md#CheckAuthenticationPost) | **Post** /check-authentication | Check authentication.
[**RevokeAuthorizationPost**](DefaultApi.md#RevokeAuthorizationPost) | **Post** /revoke-authorization | Revoke authorization.



## CheckAuthenticationPost

> CheckAuthenticationRsp CheckAuthenticationPost(ctx).CheckAuthenticationReq(checkAuthenticationReq).Execute()

Check authentication.

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
    checkAuthenticationReq := *openapiclient.NewCheckAuthenticationReq("ApiInvokerId_example", "SupportedFeatures_example") // CheckAuthenticationReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CheckAuthenticationPost(context.Background()).CheckAuthenticationReq(checkAuthenticationReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CheckAuthenticationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAuthenticationPost`: CheckAuthenticationRsp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CheckAuthenticationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAuthenticationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAuthenticationReq** | [**CheckAuthenticationReq**](CheckAuthenticationReq.md) |  | 

### Return type

[**CheckAuthenticationRsp**](CheckAuthenticationRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAuthorizationPost

> RevokeAuthorizationRsp RevokeAuthorizationPost(ctx).RevokeAuthorizationReq(revokeAuthorizationReq).Execute()

Revoke authorization.

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
    revokeAuthorizationReq := *openapiclient.NewRevokeAuthorizationReq(*openapiclient.NewSecurityNotification("ApiInvokerId_example", []string{"ApiIds_example"}, *openapiclient.NewCause()), "SupportedFeatures_example") // RevokeAuthorizationReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RevokeAuthorizationPost(context.Background()).RevokeAuthorizationReq(revokeAuthorizationReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RevokeAuthorizationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeAuthorizationPost`: RevokeAuthorizationRsp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RevokeAuthorizationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAuthorizationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revokeAuthorizationReq** | [**RevokeAuthorizationReq**](RevokeAuthorizationReq.md) |  | 

### Return type

[**RevokeAuthorizationRsp**](RevokeAuthorizationRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

