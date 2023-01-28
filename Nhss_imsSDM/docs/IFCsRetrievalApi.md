# \IFCsRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIfcs**](IFCsRetrievalApi.md#GetIfcs) | **Get** /{imsUeId}/ims-data/profile-data/ifcs | Retrieve the Initial Filter Criteria for the associated IMS subscription



## GetIfcs

> Ifcs GetIfcs(ctx, imsUeId).ApplicationServerName(applicationServerName).SupportedFeatures(supportedFeatures).Execute()

Retrieve the Initial Filter Criteria for the associated IMS subscription

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
    applicationServerName := "applicationServerName_example" // string | SIP URI of the Application Server Name (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IFCsRetrievalApi.GetIfcs(context.Background(), imsUeId).ApplicationServerName(applicationServerName).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IFCsRetrievalApi.GetIfcs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIfcs`: Ifcs
    fmt.Fprintf(os.Stdout, "Response from `IFCsRetrievalApi.GetIfcs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIfcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationServerName** | **string** | SIP URI of the Application Server Name | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**Ifcs**](Ifcs.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

