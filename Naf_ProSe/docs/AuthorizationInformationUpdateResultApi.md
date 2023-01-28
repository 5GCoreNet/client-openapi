# \AuthorizationInformationUpdateResultApi

All URIs are relative to *https://example.com/naf-prose/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizationUpdateResult**](AuthorizationInformationUpdateResultApi.md#AuthorizationUpdateResult) | **Post** /authorization-update-result | report the result of update of authorization information to revoke discovery  permissions relating to some other users in the NF consumer for Restricted ProSe Direct Discovery 



## AuthorizationUpdateResult

> AuthorizationUpdateResult(ctx).AuthUpdateData(authUpdateData).Execute()

report the result of update of authorization information to revoke discovery  permissions relating to some other users in the NF consumer for Restricted ProSe Direct Discovery 

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Naf_ProSe"
)

func main() {
    authUpdateData := *openapiclient.NewAuthUpdateData("TargetRpauid_example", []openapiclient.BannedAuthData{*openapiclient.NewBannedAuthData("BannedRpauid_example", "BannedPduid_example")}) // AuthUpdateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationInformationUpdateResultApi.AuthorizationUpdateResult(context.Background()).AuthUpdateData(authUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationInformationUpdateResultApi.AuthorizationUpdateResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationUpdateResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authUpdateData** | [**AuthUpdateData**](AuthUpdateData.md) |  | 

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

