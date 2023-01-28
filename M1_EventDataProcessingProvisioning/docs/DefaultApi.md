# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m1/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEventDataProcessingConfiguration**](DefaultApi.md#CreateEventDataProcessingConfiguration) | **Post** /provisioning-sessions/{provisioningSessionId}/event-data-processing-configurations | Supply an Event Data Processing Configuration for the specified Provisioning Session
[**DestroyEventDataProcessingConfiguration**](DefaultApi.md#DestroyEventDataProcessingConfiguration) | **Delete** /provisioning-sessions/{provisioningSessionId}/event-data-processing-configurations/{eventDataProcessingConfigurationId} | Destroy the specified Event Data Processing Configuration of the specified Provisioning Session
[**PatchEventDataProcessingConfiguration**](DefaultApi.md#PatchEventDataProcessingConfiguration) | **Patch** /provisioning-sessions/{provisioningSessionId}/event-data-processing-configurations/{eventDataProcessingConfigurationId} | Patch the specified Event Data Processing Configuration for the specified Provisioning Session
[**RetrieveEventDataProcessingConfiguration**](DefaultApi.md#RetrieveEventDataProcessingConfiguration) | **Get** /provisioning-sessions/{provisioningSessionId}/event-data-processing-configurations/{eventDataProcessingConfigurationId} | Retrieve the specified Event Data Processing Configuration of the specified Provisioning Session
[**UpdateEventDataProcessingConfiguration**](DefaultApi.md#UpdateEventDataProcessingConfiguration) | **Put** /provisioning-sessions/{provisioningSessionId}/event-data-processing-configurations/{eventDataProcessingConfigurationId} | Update the specified Event Data Processing Configuration for the specified Provisioning Session



## CreateEventDataProcessingConfiguration

> CreateEventDataProcessingConfiguration(ctx, provisioningSessionId).EventDataProcessingConfiguration(eventDataProcessingConfiguration).Execute()

Supply an Event Data Processing Configuration for the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EventDataProcessingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    eventDataProcessingConfiguration := *openapiclient.NewEventDataProcessingConfiguration("EventDataProcessingConfigurationId_example", *openapiclient.NewAfEvent(), []openapiclient.DataAccessProfile{*openapiclient.NewDataAccessProfile("DataAccessProfileId_example", []openapiclient.EventConsumerType{*openapiclient.NewEventConsumerType()}, []string{"Parameters_example"})}) // EventDataProcessingConfiguration | A JSON representation of a Event Data Processing Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateEventDataProcessingConfiguration(context.Background(), provisioningSessionId).EventDataProcessingConfiguration(eventDataProcessingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEventDataProcessingConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateEventDataProcessingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventDataProcessingConfiguration** | [**EventDataProcessingConfiguration**](EventDataProcessingConfiguration.md) | A JSON representation of a Event Data Processing Configuration | 

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


## DestroyEventDataProcessingConfiguration

> DestroyEventDataProcessingConfiguration(ctx, provisioningSessionId, eventDataProcessingConfigurationId).Execute()

Destroy the specified Event Data Processing Configuration of the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EventDataProcessingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    eventDataProcessingConfigurationId := "eventDataProcessingConfigurationId_example" // string | The resource identifier of an Event Data Processing Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyEventDataProcessingConfiguration(context.Background(), provisioningSessionId, eventDataProcessingConfigurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyEventDataProcessingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**eventDataProcessingConfigurationId** | **string** | The resource identifier of an Event Data Processing Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEventDataProcessingConfigurationRequest struct via the builder pattern


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


## PatchEventDataProcessingConfiguration

> EventDataProcessingConfiguration PatchEventDataProcessingConfiguration(ctx, provisioningSessionId, eventDataProcessingConfigurationId).EventDataProcessingConfiguration(eventDataProcessingConfiguration).Execute()

Patch the specified Event Data Processing Configuration for the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EventDataProcessingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    eventDataProcessingConfigurationId := "eventDataProcessingConfigurationId_example" // string | The resource identifier of an Event Data Processing Configuration.
    eventDataProcessingConfiguration := *openapiclient.NewEventDataProcessingConfiguration("EventDataProcessingConfigurationId_example", *openapiclient.NewAfEvent(), []openapiclient.DataAccessProfile{*openapiclient.NewDataAccessProfile("DataAccessProfileId_example", []openapiclient.EventConsumerType{*openapiclient.NewEventConsumerType()}, []string{"Parameters_example"})}) // EventDataProcessingConfiguration | A JSON representation of a Event Data Processing Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchEventDataProcessingConfiguration(context.Background(), provisioningSessionId, eventDataProcessingConfigurationId).EventDataProcessingConfiguration(eventDataProcessingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchEventDataProcessingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchEventDataProcessingConfiguration`: EventDataProcessingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchEventDataProcessingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**eventDataProcessingConfigurationId** | **string** | The resource identifier of an Event Data Processing Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEventDataProcessingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **eventDataProcessingConfiguration** | [**EventDataProcessingConfiguration**](EventDataProcessingConfiguration.md) | A JSON representation of a Event Data Processing Configuration | 

### Return type

[**EventDataProcessingConfiguration**](EventDataProcessingConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEventDataProcessingConfiguration

> EventDataProcessingConfiguration RetrieveEventDataProcessingConfiguration(ctx, provisioningSessionId, eventDataProcessingConfigurationId).Execute()

Retrieve the specified Event Data Processing Configuration of the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EventDataProcessingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    eventDataProcessingConfigurationId := "eventDataProcessingConfigurationId_example" // string | The resource identifier of an Event Data Processing Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveEventDataProcessingConfiguration(context.Background(), provisioningSessionId, eventDataProcessingConfigurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveEventDataProcessingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveEventDataProcessingConfiguration`: EventDataProcessingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveEventDataProcessingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**eventDataProcessingConfigurationId** | **string** | The resource identifier of an Event Data Processing Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEventDataProcessingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EventDataProcessingConfiguration**](EventDataProcessingConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEventDataProcessingConfiguration

> UpdateEventDataProcessingConfiguration(ctx, provisioningSessionId, eventDataProcessingConfigurationId).EventDataProcessingConfiguration(eventDataProcessingConfiguration).Execute()

Update the specified Event Data Processing Configuration for the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EventDataProcessingProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    eventDataProcessingConfigurationId := "eventDataProcessingConfigurationId_example" // string | The resource identifier of an Event Data Processing Configuration.
    eventDataProcessingConfiguration := *openapiclient.NewEventDataProcessingConfiguration("EventDataProcessingConfigurationId_example", *openapiclient.NewAfEvent(), []openapiclient.DataAccessProfile{*openapiclient.NewDataAccessProfile("DataAccessProfileId_example", []openapiclient.EventConsumerType{*openapiclient.NewEventConsumerType()}, []string{"Parameters_example"})}) // EventDataProcessingConfiguration | A JSON representation of a Event Data Processing Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateEventDataProcessingConfiguration(context.Background(), provisioningSessionId, eventDataProcessingConfigurationId).EventDataProcessingConfiguration(eventDataProcessingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEventDataProcessingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**eventDataProcessingConfigurationId** | **string** | The resource identifier of an Event Data Processing Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEventDataProcessingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **eventDataProcessingConfiguration** | [**EventDataProcessingConfiguration**](EventDataProcessingConfiguration.md) | A JSON representation of a Event Data Processing Configuration | 

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

