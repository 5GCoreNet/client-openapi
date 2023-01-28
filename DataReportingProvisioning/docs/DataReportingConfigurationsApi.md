# \DataReportingConfigurationsApi

All URIs are relative to *https://example.com/3gpp-data-reporting-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataRepConfig**](DataReportingConfigurationsApi.md#CreateDataRepConfig) | **Post** /sessions/{sessionId}/configurations | Create a new Data Reporting Configuration resource.



## CreateDataRepConfig

> DataReportingConfiguration CreateDataRepConfig(ctx, sessionId).DataReportingConfiguration(dataReportingConfiguration).Execute()

Create a new Data Reporting Configuration resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/DataReportingProvisioning"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Provisioning Session.
    dataReportingConfiguration := *openapiclient.NewDataReportingConfiguration("DataReportingConfigurationId_example", *openapiclient.NewDataCollectionClientType(), []openapiclient.DataAccessProfile{*openapiclient.NewDataAccessProfile("DataAccessProfileId_example", []openapiclient.EventConsumerType{*openapiclient.NewEventConsumerType()}, []string{"Parameters_example"})}) // DataReportingConfiguration | Representation of the Data Reporting Configuration to be created in the NEF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataReportingConfigurationsApi.CreateDataRepConfig(context.Background(), sessionId).DataReportingConfiguration(dataReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataReportingConfigurationsApi.CreateDataRepConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataRepConfig`: DataReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DataReportingConfigurationsApi.CreateDataRepConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataRepConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataReportingConfiguration** | [**DataReportingConfiguration**](DataReportingConfiguration.md) | Representation of the Data Reporting Configuration to be created in the NEF.  | 

### Return type

[**DataReportingConfiguration**](DataReportingConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

