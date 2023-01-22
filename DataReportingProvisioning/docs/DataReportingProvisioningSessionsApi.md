# \DataReportingProvisioningSessionsApi

All URIs are relative to *https://example.com/3gpp-data-reporting-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataRepProvSession**](DataReportingProvisioningSessionsApi.md#CreateDataRepProvSession) | **Post** /sessions | Create a new Data Reporting Provisioning Session.



## CreateDataRepProvSession

> DataReportingProvisioningSession CreateDataRepProvSession(ctx).DataReportingProvisioningSession(dataReportingProvisioningSession).Execute()

Create a new Data Reporting Provisioning Session.

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
    dataReportingProvisioningSession := *openapiclient.NewDataReportingProvisioningSession("ProvisioningSessionId_example", "AspId_example", "ExternalApplicationId_example", *openapiclient.NewAfEvent(), []string{"DataReportingConfigurationIds_example"}) // DataReportingProvisioningSession | Representation of the Data Reporting Provisioning Session to be created in the NEF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataReportingProvisioningSessionsApi.CreateDataRepProvSession(context.Background()).DataReportingProvisioningSession(dataReportingProvisioningSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataReportingProvisioningSessionsApi.CreateDataRepProvSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataRepProvSession`: DataReportingProvisioningSession
    fmt.Fprintf(os.Stdout, "Response from `DataReportingProvisioningSessionsApi.CreateDataRepProvSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataRepProvSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataReportingProvisioningSession** | [**DataReportingProvisioningSession**](DataReportingProvisioningSession.md) | Representation of the Data Reporting Provisioning Session to be created in the NEF.  | 

### Return type

[**DataReportingProvisioningSession**](DataReportingProvisioningSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

