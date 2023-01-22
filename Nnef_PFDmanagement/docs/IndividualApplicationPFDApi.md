# \IndividualApplicationPFDApi

All URIs are relative to *https://example.com/nnef-pfdmanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NnefPFDmanagementIndAppFetch**](IndividualApplicationPFDApi.md#NnefPFDmanagementIndAppFetch) | **Get** /applications/{appId} | Retrieve the PFD for an application.



## NnefPFDmanagementIndAppFetch

> PfdDataForApp NnefPFDmanagementIndAppFetch(ctx, appId).SupportedFeatures(supportedFeatures).Execute()

Retrieve the PFD for an application.

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
    appId := "appId_example" // string | The required application identifier(s) for the returned PFDs.
    supportedFeatures := "supportedFeatures_example" // string | To filter irrelevant responses related to unsupported features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationPFDApi.NnefPFDmanagementIndAppFetch(context.Background(), appId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationPFDApi.NnefPFDmanagementIndAppFetch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NnefPFDmanagementIndAppFetch`: PfdDataForApp
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationPFDApi.NnefPFDmanagementIndAppFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The required application identifier(s) for the returned PFDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNnefPFDmanagementIndAppFetchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | To filter irrelevant responses related to unsupported features | 

### Return type

[**PfdDataForApp**](PfdDataForApp.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

