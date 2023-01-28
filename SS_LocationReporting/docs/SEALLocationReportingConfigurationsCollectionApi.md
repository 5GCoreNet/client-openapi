# \SEALLocationReportingConfigurationsCollectionApi

All URIs are relative to *https://example.com/ss-lr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocReportingConfig**](SEALLocationReportingConfigurationsCollectionApi.md#CreateLocReportingConfig) | **Post** /trigger-configurations | 



## CreateLocReportingConfig

> LocationReportConfiguration CreateLocReportingConfig(ctx).LocationReportConfiguration(locationReportConfiguration).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_LocationReporting"
)

func main() {
    locationReportConfiguration := *openapiclient.NewLocationReportConfiguration("ValServerId_example", openapiclient.ValTargetUe{Interface{}: new(interface{})}) // LocationReportConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SEALLocationReportingConfigurationsCollectionApi.CreateLocReportingConfig(context.Background()).LocationReportConfiguration(locationReportConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SEALLocationReportingConfigurationsCollectionApi.CreateLocReportingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocReportingConfig`: LocationReportConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SEALLocationReportingConfigurationsCollectionApi.CreateLocReportingConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocReportingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationReportConfiguration** | [**LocationReportConfiguration**](LocationReportConfiguration.md) |  | 

### Return type

[**LocationReportConfiguration**](LocationReportConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

