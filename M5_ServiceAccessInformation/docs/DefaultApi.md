# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m5/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveServiceAccessInformation**](DefaultApi.md#RetrieveServiceAccessInformation) | **Get** /service-access-information/{provisioningSessionId} | Retrieve the Service Access Information resource



## RetrieveServiceAccessInformation

> ServiceAccessInformationResource RetrieveServiceAccessInformation(ctx, provisioningSessionId).Execute()

Retrieve the Service Access Information resource

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveServiceAccessInformation(context.Background(), provisioningSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveServiceAccessInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveServiceAccessInformation`: ServiceAccessInformationResource
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveServiceAccessInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveServiceAccessInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccessInformationResource**](ServiceAccessInformationResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

