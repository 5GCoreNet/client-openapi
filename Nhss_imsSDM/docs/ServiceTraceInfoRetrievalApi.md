# \ServiceTraceInfoRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceTraceInfo**](ServiceTraceInfoRetrievalApi.md#GetServiceTraceInfo) | **Get** /{imsUeId}/ims-data/profile-data/service-level-trace-information | Retrieve the IMS service level trace information for the associated user



## GetServiceTraceInfo

> ServiceLevelTraceInformation GetServiceTraceInfo(ctx, imsUeId).SupportedFeatures(supportedFeatures).Execute()

Retrieve the IMS service level trace information for the associated user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsSDM"
)

func main() {
    imsUeId := "imsUeId_example" // string | IMS Identity
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceTraceInfoRetrievalApi.GetServiceTraceInfo(context.Background(), imsUeId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceTraceInfoRetrievalApi.GetServiceTraceInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceTraceInfo`: ServiceLevelTraceInformation
    fmt.Fprintf(os.Stdout, "Response from `ServiceTraceInfoRetrievalApi.GetServiceTraceInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTraceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**ServiceLevelTraceInformation**](ServiceLevelTraceInformation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

