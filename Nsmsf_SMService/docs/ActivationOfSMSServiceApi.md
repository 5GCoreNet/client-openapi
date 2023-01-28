# \ActivationOfSMSServiceApi

All URIs are relative to *https://example.com/nsmsf-sms/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SMServiceActivation**](ActivationOfSMSServiceApi.md#SMServiceActivation) | **Put** /ue-contexts/{supi} | Activate SMS Service for a given UE



## SMServiceActivation

> UeSmsContextData SMServiceActivation(ctx, supi).UeSmsContextData(ueSmsContextData).Execute()

Activate SMS Service for a given UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsmsf_SMService"
)

func main() {
    supi := "supi_example" // string | Subscriber Permanent Identifier (SUPI)
    ueSmsContextData := *openapiclient.NewUeSmsContextData("Supi_example", "AmfId_example", openapiclient.AccessType("3GPP_ACCESS")) // UeSmsContextData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivationOfSMSServiceApi.SMServiceActivation(context.Background(), supi).UeSmsContextData(ueSmsContextData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivationOfSMSServiceApi.SMServiceActivation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SMServiceActivation`: UeSmsContextData
    fmt.Fprintf(os.Stdout, "Response from `ActivationOfSMSServiceApi.SMServiceActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Subscriber Permanent Identifier (SUPI) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSMServiceActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ueSmsContextData** | [**UeSmsContextData**](UeSmsContextData.md) |  | 

### Return type

[**UeSmsContextData**](UeSmsContextData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

