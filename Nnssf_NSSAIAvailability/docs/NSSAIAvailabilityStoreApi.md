# \NSSAIAvailabilityStoreApi

All URIs are relative to *https://example.com/nnssf-nssaiavailability/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NSSAIAvailabilityOptions**](NSSAIAvailabilityStoreApi.md#NSSAIAvailabilityOptions) | **Options** /nssai-availability | Discover communication options supported by NSSF for NSSAI Availability



## NSSAIAvailabilityOptions

> NSSAIAvailabilityOptions(ctx).Execute()

Discover communication options supported by NSSF for NSSAI Availability

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnssf_NSSAIAvailability"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NSSAIAvailabilityStoreApi.NSSAIAvailabilityOptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NSSAIAvailabilityStoreApi.NSSAIAvailabilityOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNSSAIAvailabilityOptionsRequest struct via the builder pattern


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

