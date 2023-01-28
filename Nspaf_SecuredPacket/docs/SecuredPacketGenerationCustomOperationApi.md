# \SecuredPacketGenerationCustomOperationApi

All URIs are relative to *https://example.com/nspaf-secured-packet/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvideSecuredPacket**](SecuredPacketGenerationCustomOperationApi.md#ProvideSecuredPacket) | **Post** /{supi}/provide-secured-packet | request generation of a secured packet



## ProvideSecuredPacket

> string ProvideSecuredPacket(ctx, supi).UiccConfigurationParameter(uiccConfigurationParameter).Execute()

request generation of a secured packet

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nspaf_SecuredPacket"
)

func main() {
    supi := "supi_example" // string | SUPI of the user
    uiccConfigurationParameter := openapiclient.UiccConfigurationParameter{Interface{}: new(interface{})} // UiccConfigurationParameter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecuredPacketGenerationCustomOperationApi.ProvideSecuredPacket(context.Background(), supi).UiccConfigurationParameter(uiccConfigurationParameter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuredPacketGenerationCustomOperationApi.ProvideSecuredPacket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvideSecuredPacket`: string
    fmt.Fprintf(os.Stdout, "Response from `SecuredPacketGenerationCustomOperationApi.ProvideSecuredPacket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | SUPI of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvideSecuredPacketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uiccConfigurationParameter** | [**UiccConfigurationParameter**](UiccConfigurationParameter.md) |  | 

### Return type

**string**

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

