# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m5/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitConsumptionReport**](DefaultApi.md#SubmitConsumptionReport) | **Post** /consumption-reporting/{aspId} | Submit a Consumption Report



## SubmitConsumptionReport

> SubmitConsumptionReport(ctx, aspId).ConsumptionReport(consumptionReport).Execute()

Submit a Consumption Report

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    aspId := "aspId_example" // string | See 3GPP TS 26.512 clause 11.3.2.
    consumptionReport := *openapiclient.NewConsumptionReport("MediaPlayerEntry_example", "ReportingClientId_example", []openapiclient.ConsumptionReportingUnit{*openapiclient.NewConsumptionReportingUnit("MediaConsumed_example", time.Now(), int32(123))}) // ConsumptionReport | A Consumption Report

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubmitConsumptionReport(context.Background(), aspId).ConsumptionReport(consumptionReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubmitConsumptionReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aspId** | **string** | See 3GPP TS 26.512 clause 11.3.2. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitConsumptionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumptionReport** | [**ConsumptionReport**](ConsumptionReport.md) | A Consumption Report | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

