# \SMSFRegistrationFor3GPPAccessApi

All URIs are relative to *https://example.com/nudm-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Call3GppSmsfRegistration**](SMSFRegistrationFor3GPPAccessApi.md#Call3GppSmsfRegistration) | **Put** /{ueId}/registrations/smsf-3gpp-access | register as SMSF for 3GPP access



## Call3GppSmsfRegistration

> SmsfRegistration Call3GppSmsfRegistration(ctx, ueId).SmsfRegistration(smsfRegistration).Execute()

register as SMSF for 3GPP access

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
    ueId := "ueId_example" // string | Identifier of the UE
    smsfRegistration := *openapiclient.NewSmsfRegistration("SmsfInstanceId_example", *openapiclient.NewPlmnId("Mcc_example", "Mnc_example")) // SmsfRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSFRegistrationFor3GPPAccessApi.Call3GppSmsfRegistration(context.Background(), ueId).SmsfRegistration(smsfRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSFRegistrationFor3GPPAccessApi.Call3GppSmsfRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call3GppSmsfRegistration`: SmsfRegistration
    fmt.Fprintf(os.Stdout, "Response from `SMSFRegistrationFor3GPPAccessApi.Call3GppSmsfRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiCall3GppSmsfRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsfRegistration** | [**SmsfRegistration**](SmsfRegistration.md) |  | 

### Return type

[**SmsfRegistration**](SmsfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
