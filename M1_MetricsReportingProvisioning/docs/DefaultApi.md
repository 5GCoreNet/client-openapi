# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m1/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateMetricsReporting**](DefaultApi.md#ActivateMetricsReporting) | **Post** /provisioning-sessions/{provisioningSessionId}/metrics-reporting-configurations | Activate the Metrics reporting procedure for the specified Provisioning Session by providing the Metrics Reporting Configuration
[**DestroyMetricsReportingConfiguration**](DefaultApi.md#DestroyMetricsReportingConfiguration) | **Delete** /provisioning-sessions/{provisioningSessionId}/metrics-reporting-configurations/{metricsReportingConfigurationId} | Destroy the specified Metrics Reporting Configuration of the specified Provisioning Session
[**PatchMetricsReportingConfiguration**](DefaultApi.md#PatchMetricsReportingConfiguration) | **Patch** /provisioning-sessions/{provisioningSessionId}/metrics-reporting-configurations/{metricsReportingConfigurationId} | Patch the specified Metrics Reporting Configuration for the specified Provisioning Session
[**RetrieveMetricsReportingConfiguration**](DefaultApi.md#RetrieveMetricsReportingConfiguration) | **Get** /provisioning-sessions/{provisioningSessionId}/metrics-reporting-configurations/{metricsReportingConfigurationId} | Retrieve the specified Metrics Reporting Configuration of the specified Provisioning Session
[**UpdateMetricsReportingConfiguration**](DefaultApi.md#UpdateMetricsReportingConfiguration) | **Put** /provisioning-sessions/{provisioningSessionId}/metrics-reporting-configurations/{metricsReportingConfigurationId} | Update the specified Metrics Reporting Configuration for the specified Provisioning Session



## ActivateMetricsReporting

> ActivateMetricsReporting(ctx, provisioningSessionId).MetricsReportingConfiguration(metricsReportingConfiguration).Execute()

Activate the Metrics reporting procedure for the specified Provisioning Session by providing the Metrics Reporting Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_MetricsReportingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    metricsReportingConfiguration := *openapiclient.NewMetricsReportingConfiguration("MetricsReportingConfigurationId_example", "Scheme_example") // MetricsReportingConfiguration | A JSON representation of a Metrics Reporting Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ActivateMetricsReporting(context.Background(), provisioningSessionId).MetricsReportingConfiguration(metricsReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ActivateMetricsReporting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateMetricsReportingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsReportingConfiguration** | [**MetricsReportingConfiguration**](MetricsReportingConfiguration.md) | A JSON representation of a Metrics Reporting Configuration | 

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


## DestroyMetricsReportingConfiguration

> DestroyMetricsReportingConfiguration(ctx, provisioningSessionId, metricsReportingConfigurationId).Execute()

Destroy the specified Metrics Reporting Configuration of the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_MetricsReportingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    metricsReportingConfigurationId := "metricsReportingConfigurationId_example" // string | The resource identifier of a Metrics Reporting Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyMetricsReportingConfiguration(context.Background(), provisioningSessionId, metricsReportingConfigurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyMetricsReportingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**metricsReportingConfigurationId** | **string** | The resource identifier of a Metrics Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyMetricsReportingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMetricsReportingConfiguration

> MetricsReportingConfiguration PatchMetricsReportingConfiguration(ctx, provisioningSessionId, metricsReportingConfigurationId).MetricsReportingConfiguration(metricsReportingConfiguration).Execute()

Patch the specified Metrics Reporting Configuration for the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_MetricsReportingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    metricsReportingConfigurationId := "metricsReportingConfigurationId_example" // string | The resource identifier of a Metrics Reporting Configuration.
    metricsReportingConfiguration := *openapiclient.NewMetricsReportingConfiguration("MetricsReportingConfigurationId_example", "Scheme_example") // MetricsReportingConfiguration | A JSON representation of a Metrics Reporting Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchMetricsReportingConfiguration(context.Background(), provisioningSessionId, metricsReportingConfigurationId).MetricsReportingConfiguration(metricsReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchMetricsReportingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMetricsReportingConfiguration`: MetricsReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchMetricsReportingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**metricsReportingConfigurationId** | **string** | The resource identifier of a Metrics Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMetricsReportingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **metricsReportingConfiguration** | [**MetricsReportingConfiguration**](MetricsReportingConfiguration.md) | A JSON representation of a Metrics Reporting Configuration | 

### Return type

[**MetricsReportingConfiguration**](MetricsReportingConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMetricsReportingConfiguration

> MetricsReportingConfiguration RetrieveMetricsReportingConfiguration(ctx, provisioningSessionId, metricsReportingConfigurationId).Execute()

Retrieve the specified Metrics Reporting Configuration of the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_MetricsReportingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    metricsReportingConfigurationId := "metricsReportingConfigurationId_example" // string | The resource identifier of a Metrics Reporting Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveMetricsReportingConfiguration(context.Background(), provisioningSessionId, metricsReportingConfigurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveMetricsReportingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMetricsReportingConfiguration`: MetricsReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveMetricsReportingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**metricsReportingConfigurationId** | **string** | The resource identifier of a Metrics Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMetricsReportingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MetricsReportingConfiguration**](MetricsReportingConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricsReportingConfiguration

> UpdateMetricsReportingConfiguration(ctx, provisioningSessionId, metricsReportingConfigurationId).MetricsReportingConfiguration(metricsReportingConfiguration).Execute()

Update the specified Metrics Reporting Configuration for the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_MetricsReportingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    metricsReportingConfigurationId := "metricsReportingConfigurationId_example" // string | The resource identifier of a Metrics Reporting Configuration.
    metricsReportingConfiguration := *openapiclient.NewMetricsReportingConfiguration("MetricsReportingConfigurationId_example", "Scheme_example") // MetricsReportingConfiguration | A JSON representation of a Metrics Reporting Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateMetricsReportingConfiguration(context.Background(), provisioningSessionId, metricsReportingConfigurationId).MetricsReportingConfiguration(metricsReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMetricsReportingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**metricsReportingConfigurationId** | **string** | The resource identifier of a Metrics Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricsReportingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **metricsReportingConfiguration** | [**MetricsReportingConfiguration**](MetricsReportingConfiguration.md) | A JSON representation of a Metrics Reporting Configuration | 

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

