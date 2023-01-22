# \DeactivationOfSMSServiceApi

All URIs are relative to *https://example.com/nsmsf-sms/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SMServiceDeactivation**](DeactivationOfSMSServiceApi.md#SMServiceDeactivation) | **Delete** /ue-contexts/{supi} | Deactivate SMS Service for a given UE



## SMServiceDeactivation

> SMServiceDeactivation(ctx, supi).IfMatch(ifMatch).Execute()

Deactivate SMS Service for a given UE

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
    supi := "supi_example" // string | Subscriber Permanent Identifier (SUPI)
    ifMatch := "ifMatch_example" // string | Validator for conditional requests, as described in IETF RFC 7232, 3.1 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeactivationOfSMSServiceApi.SMServiceDeactivation(context.Background(), supi).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeactivationOfSMSServiceApi.SMServiceDeactivation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Subscriber Permanent Identifier (SUPI) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSMServiceDeactivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Validator for conditional requests, as described in IETF RFC 7232, 3.1 | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

