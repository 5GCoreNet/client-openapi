# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m1/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateConsumptionReporting**](DefaultApi.md#ActivateConsumptionReporting) | **Post** /provisioning-sessions/{provisioningSessionId}/consumption-reporting-configuration | Activate the consumption reporting procedure for the specified Provisioning Session by providing the Consumption Reporting Configuration
[**DestroyConsumptionReportingConfiguration**](DefaultApi.md#DestroyConsumptionReportingConfiguration) | **Delete** /provisioning-sessions/{provisioningSessionId}/consumption-reporting-configuration | Destroy the current Consumption Reporting Configuration of the specified Provisioning Session
[**PatchConsumptionReportingConfiguration**](DefaultApi.md#PatchConsumptionReportingConfiguration) | **Patch** /provisioning-sessions/{provisioningSessionId}/consumption-reporting-configuration | Patch the Consumption Reporting Configuration for the specified Provisioning Session
[**RetrieveConsumptionReportingConfiguration**](DefaultApi.md#RetrieveConsumptionReportingConfiguration) | **Get** /provisioning-sessions/{provisioningSessionId}/consumption-reporting-configuration | Retrieve the Consumption Reporting Configuration of the specified Provisioning Session
[**UpdateConsumptionReportingConfiguration**](DefaultApi.md#UpdateConsumptionReportingConfiguration) | **Put** /provisioning-sessions/{provisioningSessionId}/consumption-reporting-configuration | Update the Consumption Reporting Configuration for the specified Provisioning Session



## ActivateConsumptionReporting

> ActivateConsumptionReporting(ctx, provisioningSessionId).ConsumptionReportingConfiguration(consumptionReportingConfiguration).Execute()

Activate the consumption reporting procedure for the specified Provisioning Session by providing the Consumption Reporting Configuration

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
    consumptionReportingConfiguration := *openapiclient.NewConsumptionReportingConfiguration() // ConsumptionReportingConfiguration | A JSON representation of a Consumption Reporting Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ActivateConsumptionReporting(context.Background(), provisioningSessionId).ConsumptionReportingConfiguration(consumptionReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ActivateConsumptionReporting``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActivateConsumptionReportingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumptionReportingConfiguration** | [**ConsumptionReportingConfiguration**](ConsumptionReportingConfiguration.md) | A JSON representation of a Consumption Reporting Configuration | 

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


## DestroyConsumptionReportingConfiguration

> DestroyConsumptionReportingConfiguration(ctx, provisioningSessionId).Execute()

Destroy the current Consumption Reporting Configuration of the specified Provisioning Session

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyConsumptionReportingConfiguration(context.Background(), provisioningSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyConsumptionReportingConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDestroyConsumptionReportingConfigurationRequest struct via the builder pattern


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


## PatchConsumptionReportingConfiguration

> ConsumptionReportingConfiguration PatchConsumptionReportingConfiguration(ctx, provisioningSessionId).ConsumptionReportingConfiguration(consumptionReportingConfiguration).Execute()

Patch the Consumption Reporting Configuration for the specified Provisioning Session

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
    consumptionReportingConfiguration := *openapiclient.NewConsumptionReportingConfiguration() // ConsumptionReportingConfiguration | A JSON representation of a Consumption Reporting Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchConsumptionReportingConfiguration(context.Background(), provisioningSessionId).ConsumptionReportingConfiguration(consumptionReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchConsumptionReportingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchConsumptionReportingConfiguration`: ConsumptionReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchConsumptionReportingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConsumptionReportingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumptionReportingConfiguration** | [**ConsumptionReportingConfiguration**](ConsumptionReportingConfiguration.md) | A JSON representation of a Consumption Reporting Configuration | 

### Return type

[**ConsumptionReportingConfiguration**](ConsumptionReportingConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveConsumptionReportingConfiguration

> ConsumptionReportingConfiguration RetrieveConsumptionReportingConfiguration(ctx, provisioningSessionId).Execute()

Retrieve the Consumption Reporting Configuration of the specified Provisioning Session

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveConsumptionReportingConfiguration(context.Background(), provisioningSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveConsumptionReportingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveConsumptionReportingConfiguration`: ConsumptionReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveConsumptionReportingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveConsumptionReportingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsumptionReportingConfiguration**](ConsumptionReportingConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConsumptionReportingConfiguration

> UpdateConsumptionReportingConfiguration(ctx, provisioningSessionId).ConsumptionReportingConfiguration(consumptionReportingConfiguration).Execute()

Update the Consumption Reporting Configuration for the specified Provisioning Session

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
    consumptionReportingConfiguration := *openapiclient.NewConsumptionReportingConfiguration() // ConsumptionReportingConfiguration | A JSON representation of a Consumption Reporting Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateConsumptionReportingConfiguration(context.Background(), provisioningSessionId).ConsumptionReportingConfiguration(consumptionReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConsumptionReportingConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateConsumptionReportingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumptionReportingConfiguration** | [**ConsumptionReportingConfiguration**](ConsumptionReportingConfiguration.md) | A JSON representation of a Consumption Reporting Configuration | 

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

