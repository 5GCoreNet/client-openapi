# \DefaultApi

All URIs are relative to *https://example.com/3gpp-ndcaf_data-reporting/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSession**](DefaultApi.md#CreateSession) | **Post** /sessions | Create a new Data Reporting Session
[**DestroySession**](DefaultApi.md#DestroySession) | **Delete** /sessions/{sessionId} | Destroy an existing Data Reporting Session
[**Report**](DefaultApi.md#Report) | **Post** /sessions/{sessionId}/report | Report UE data in the context of an existing Data Reporting Session
[**RetrieveSession**](DefaultApi.md#RetrieveSession) | **Get** /sessions/{sessionId} | Retrieve an existing Data Reporting Session



## CreateSession

> DataReportingSession CreateSession(ctx).DataReportingSession(dataReportingSession).Execute()

Create a new Data Reporting Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ndcaf_DataReporting"
)

func main() {
    dataReportingSession := *openapiclient.NewDataReportingSession("ExternalApplicationId_example", []openapiclient.DataDomain{*openapiclient.NewDataDomain()}) // DataReportingSession | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSession(context.Background()).DataReportingSession(dataReportingSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSession`: DataReportingSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


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


## DestroySession

> DestroySession(ctx, sessionId).Execute()

Destroy an existing Data Reporting Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ndcaf_DataReporting"
)

func main() {
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroySession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroySession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The resource identifier of an existing Data Reporting Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroySessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Report

> DataReportingSession Report(ctx, sessionId).DataReport(dataReport).Execute()

Report UE data in the context of an existing Data Reporting Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ndcaf_DataReporting"
)

func main() {
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Session.
    dataReport := *openapiclient.NewDataReport("ExternalApplicationId_example") // DataReport | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Report(context.Background(), sessionId).DataReport(dataReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Report``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Report`: DataReportingSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Report`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The resource identifier of an existing Data Reporting Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportRequest struct via the builder pattern


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


## RetrieveSession

> DataReportingSession RetrieveSession(ctx, sessionId).Execute()

Retrieve an existing Data Reporting Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ndcaf_DataReporting"
)

func main() {
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSession`: DataReportingSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The resource identifier of an existing Data Reporting Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSessionRequest struct via the builder pattern


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

