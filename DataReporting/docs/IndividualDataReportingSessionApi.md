# \IndividualDataReportingSessionApi

All URIs are relative to *https://example.com/3gpp-data-reporting/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndDataRepSession**](IndividualDataReportingSessionApi.md#DeleteIndDataRepSession) | **Delete** /sessions/{sessionId} | Deletes an already existing Data Reporting Session resource.
[**GetIndDataRepSession**](IndividualDataReportingSessionApi.md#GetIndDataRepSession) | **Get** /sessions/{sessionId} | Request the retrieval of an existing Individual Data Reporting Session resource.
[**ReportUEData**](IndividualDataReportingSessionApi.md#ReportUEData) | **Post** /sessions/{sessionId}/report | Report collected UE data.
[**UpdateIndDataRepSession**](IndividualDataReportingSessionApi.md#UpdateIndDataRepSession) | **Put** /sessions/{sessionId} | Request the update of an existing Individual Data Reporting Session resource.



## DeleteIndDataRepSession

> DeleteIndDataRepSession(ctx, sessionId).Execute()

Deletes an already existing Data Reporting Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/DataReporting"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingSessionApi.DeleteIndDataRepSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingSessionApi.DeleteIndDataRepSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndDataRepSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndDataRepSession

> DataReportingSession GetIndDataRepSession(ctx, sessionId).Execute()

Request the retrieval of an existing Individual Data Reporting Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/DataReporting"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingSessionApi.GetIndDataRepSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingSessionApi.GetIndDataRepSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndDataRepSession`: DataReportingSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualDataReportingSessionApi.GetIndDataRepSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndDataRepSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataReportingSession**](DataReportingSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportUEData

> DataReportingSession ReportUEData(ctx, sessionId).DataReport(dataReport).Execute()

Report collected UE data.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/DataReporting"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Session.
    dataReport := *openapiclient.NewDataReport("ExternalApplicationId_example") // DataReport | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingSessionApi.ReportUEData(context.Background(), sessionId).DataReport(dataReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingSessionApi.ReportUEData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportUEData`: DataReportingSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualDataReportingSessionApi.ReportUEData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportUEDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataReport** | [**DataReport**](DataReport.md) |  | 

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


## UpdateIndDataRepSession

> DataReportingSession UpdateIndDataRepSession(ctx, sessionId).DataReportingSession(dataReportingSession).Execute()

Request the update of an existing Individual Data Reporting Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/DataReporting"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Session.
    dataReportingSession := *openapiclient.NewDataReportingSession("ExternalApplicationId_example", []openapiclient.DataDomain{*openapiclient.NewDataDomain()}) // DataReportingSession | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingSessionApi.UpdateIndDataRepSession(context.Background(), sessionId).DataReportingSession(dataReportingSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingSessionApi.UpdateIndDataRepSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndDataRepSession`: DataReportingSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualDataReportingSessionApi.UpdateIndDataRepSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndDataRepSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataReportingSession** | [**DataReportingSession**](DataReportingSession.md) |  | 

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

