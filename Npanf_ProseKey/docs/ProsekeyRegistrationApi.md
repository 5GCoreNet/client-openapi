# \ProsekeyRegistrationApi

All URIs are relative to *https://example.com/npanf-prosekey/&lt;apiVersion&gt;*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProseKeyRegistration**](ProsekeyRegistrationApi.md#ProseKeyRegistration) | **Post** /prose-keys/register | Register the Prose Key



## ProseKeyRegistration

> ProseKeyRegistration(ctx).ProseContextInfo(proseContextInfo).Execute()

Register the Prose Key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npanf_ProseKey"
)

func main() {
    proseContextInfo := *openapiclient.NewProseContextInfo("Supi_example", "Var5gPruk_example", "Var5gPrukId_example", int32(123)) // ProseContextInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProsekeyRegistrationApi.ProseKeyRegistration(context.Background()).ProseContextInfo(proseContextInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProsekeyRegistrationApi.ProseKeyRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProseKeyRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proseContextInfo** | [**ProseContextInfo**](ProseContextInfo.md) |  | 

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

