# \PFDOfApplicationsApi

All URIs are relative to *https://example.com/nnef-pfdmanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NnefPFDmanagementAllFetch**](PFDOfApplicationsApi.md#NnefPFDmanagementAllFetch) | **Get** /applications | Retrieve PFDs for all applications or for one or multiple applications with query parameter.



## NnefPFDmanagementAllFetch

> []PfdDataForApp NnefPFDmanagementAllFetch(ctx).ApplicationIds(applicationIds).SupportedFeatures(supportedFeatures).Execute()

Retrieve PFDs for all applications or for one or multiple applications with query parameter.

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
    applicationIds := []string{"Inner_example"} // []string | The required application identifier(s) for the returned PFDs.
    supportedFeatures := "supportedFeatures_example" // string | To filter irrelevant responses related to unsupported features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PFDOfApplicationsApi.NnefPFDmanagementAllFetch(context.Background()).ApplicationIds(applicationIds).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PFDOfApplicationsApi.NnefPFDmanagementAllFetch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NnefPFDmanagementAllFetch`: []PfdDataForApp
    fmt.Fprintf(os.Stdout, "Response from `PFDOfApplicationsApi.NnefPFDmanagementAllFetch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNnefPFDmanagementAllFetchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **[]string** | The required application identifier(s) for the returned PFDs. | 
 **supportedFeatures** | **string** | To filter irrelevant responses related to unsupported features | 

### Return type

[**[]PfdDataForApp**](PfdDataForApp.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

