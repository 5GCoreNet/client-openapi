# \IndividualGeographicalAreaDocumentApi

All URIs are relative to *https://example.com/vae-service-continuity/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryServiceContinuity**](IndividualGeographicalAreaDocumentApi.md#QueryServiceContinuity) | **Get** /geo-areas/{geoId} | VAE service continuity query service operation



## QueryServiceContinuity

> V2xServiceInfo QueryServiceContinuity(ctx, geoId).ServiceId(serviceId).SuppFeat(suppFeat).Execute()

VAE service continuity query service operation

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
    geoId := "geoId_example" // string | Identifier of a geographical area
    serviceId := "serviceId_example" // string | Identifier of a V2X service
    suppFeat := "suppFeat_example" // string | To filter irrelevant responses related to unsupported features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualGeographicalAreaDocumentApi.QueryServiceContinuity(context.Background(), geoId).ServiceId(serviceId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualGeographicalAreaDocumentApi.QueryServiceContinuity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryServiceContinuity`: V2xServiceInfo
    fmt.Fprintf(os.Stdout, "Response from `IndividualGeographicalAreaDocumentApi.QueryServiceContinuity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geoId** | **string** | Identifier of a geographical area | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryServiceContinuityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceId** | **string** | Identifier of a V2X service | 
 **suppFeat** | **string** | To filter irrelevant responses related to unsupported features | 

### Return type

[**V2xServiceInfo**](V2xServiceInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

