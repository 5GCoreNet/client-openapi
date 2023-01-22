# \DataReportingSessionsApi

All URIs are relative to *https://example.com/3gpp-data-reporting/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataRepSession**](DataReportingSessionsApi.md#CreateDataRepSession) | **Post** /sessions | Create a new Data Reporting Session.



## CreateDataRepSession

> DataReportingSession CreateDataRepSession(ctx).DataReportingSession(dataReportingSession).Execute()

Create a new Data Reporting Session.

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
    dataReportingSession := *openapiclient.NewDataReportingSession("ExternalApplicationId_example", []openapiclient.DataDomain{*openapiclient.NewDataDomain()}) // DataReportingSession | Representation of the Data Reporting Session to be created in the NEF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataReportingSessionsApi.CreateDataRepSession(context.Background()).DataReportingSession(dataReportingSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataReportingSessionsApi.CreateDataRepSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataRepSession`: DataReportingSession
    fmt.Fprintf(os.Stdout, "Response from `DataReportingSessionsApi.CreateDataRepSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataRepSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataReportingSession** | [**DataReportingSession**](DataReportingSession.md) | Representation of the Data Reporting Session to be created in the NEF.  | 

### Return type

[**DataReportingSession**](DataReportingSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

