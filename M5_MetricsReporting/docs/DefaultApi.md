# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m5/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitMetricsReport**](DefaultApi.md#SubmitMetricsReport) | **Post** /metrics-reporting/{provisioningSessionId}/{metricsReportingConfigurationId} | Submit a Metrics Report



## SubmitMetricsReport

> SubmitMetricsReport(ctx, provisioningSessionId, metricsReportingConfigurationId).Body(body).Execute()

Submit a Metrics Report

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    metricsReportingConfigurationId := "metricsReportingConfigurationId_example" // string | The resource identifier of a Metrics Configuration in the specified Provisioning Session.
    body := "body_example" // string | A Metrics Report

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubmitMetricsReport(context.Background(), provisioningSessionId, metricsReportingConfigurationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubmitMetricsReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**metricsReportingConfigurationId** | **string** | The resource identifier of a Metrics Configuration in the specified Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitMetricsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | A Metrics Report | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/3gpdash-qoe-report+xml, application/*
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

