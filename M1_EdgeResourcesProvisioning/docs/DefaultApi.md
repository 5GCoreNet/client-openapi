# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m1/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeResourcesConfiguration**](DefaultApi.md#CreateEdgeResourcesConfiguration) | **Post** /provisioning-sessions/{provisioningSessionId}/edge-resources-configurations | Create an Edge Resources Configuration within the scope of the specified Provisioning Session
[**DestroyEdgeResourcesConfiguration**](DefaultApi.md#DestroyEdgeResourcesConfiguration) | **Delete** /provisioning-sessions/{provisioningSessionId}/edge-resources-configurations/{edgeResourcesConfigurationId} | 
[**PatchEdgeResourcesConfiguration**](DefaultApi.md#PatchEdgeResourcesConfiguration) | **Patch** /provisioning-sessions/{provisioningSessionId}/edge-resources-configurations/{edgeResourcesConfigurationId} | Patch the Edge Resources Configuration for the specified Provisioning Session
[**RetrieveEdgeResourcesConfiguration**](DefaultApi.md#RetrieveEdgeResourcesConfiguration) | **Get** /provisioning-sessions/{provisioningSessionId}/edge-resources-configurations/{edgeResourcesConfigurationId} | Retrieve the Edge Resources Configuration of the specified Provisioning Session
[**UpdateEdgeResourcesConfiguration**](DefaultApi.md#UpdateEdgeResourcesConfiguration) | **Put** /provisioning-sessions/{provisioningSessionId}/edge-resources-configurations/{edgeResourcesConfigurationId} | Update an Edge Resources Configuration for the specified Provisioning Session



## CreateEdgeResourcesConfiguration

> CreateEdgeResourcesConfiguration(ctx, provisioningSessionId).EdgeResourcesConfiguration(edgeResourcesConfiguration).Execute()

Create an Edge Resources Configuration within the scope of the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EdgeResourcesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    edgeResourcesConfiguration := *openapiclient.NewEdgeResourcesConfiguration("EdgeResourcesConfigurationId_example", *openapiclient.NewEdgeManagementMode(), *openapiclient.NewEASRequirements([]string{"EasProviderIds_example"}, "EasType_example", []string{"EasFeatures_example"}, []openapiclient.ScheduledCommunicationTime{*openapiclient.NewScheduledCommunicationTime()}, []openapiclient.ACRScenario{*openapiclient.NewACRScenario()})) // EdgeResourcesConfiguration | A JSON representation of an Edge Resources Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateEdgeResourcesConfiguration(context.Background(), provisioningSessionId).EdgeResourcesConfiguration(edgeResourcesConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEdgeResourcesConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateEdgeResourcesConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeResourcesConfiguration** | [**EdgeResourcesConfiguration**](EdgeResourcesConfiguration.md) | A JSON representation of an Edge Resources Configuration | 

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


## DestroyEdgeResourcesConfiguration

> DestroyEdgeResourcesConfiguration(ctx, provisioningSessionId, edgeResourcesConfigurationId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EdgeResourcesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    edgeResourcesConfigurationId := "edgeResourcesConfigurationId_example" // string | The resource identifier of an existing Edge Resources Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyEdgeResourcesConfiguration(context.Background(), provisioningSessionId, edgeResourcesConfigurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyEdgeResourcesConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**edgeResourcesConfigurationId** | **string** | The resource identifier of an existing Edge Resources Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeResourcesConfigurationRequest struct via the builder pattern


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


## PatchEdgeResourcesConfiguration

> EdgeResourcesConfiguration PatchEdgeResourcesConfiguration(ctx, provisioningSessionId, edgeResourcesConfigurationId).EdgeResourcesConfiguration(edgeResourcesConfiguration).Execute()

Patch the Edge Resources Configuration for the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EdgeResourcesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    edgeResourcesConfigurationId := "edgeResourcesConfigurationId_example" // string | The resource identifier of an existing Edge Resources Configuration.
    edgeResourcesConfiguration := *openapiclient.NewEdgeResourcesConfiguration("EdgeResourcesConfigurationId_example", *openapiclient.NewEdgeManagementMode(), *openapiclient.NewEASRequirements([]string{"EasProviderIds_example"}, "EasType_example", []string{"EasFeatures_example"}, []openapiclient.ScheduledCommunicationTime{*openapiclient.NewScheduledCommunicationTime()}, []openapiclient.ACRScenario{*openapiclient.NewACRScenario()})) // EdgeResourcesConfiguration | A JSON representation of a Edge Resources Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchEdgeResourcesConfiguration(context.Background(), provisioningSessionId, edgeResourcesConfigurationId).EdgeResourcesConfiguration(edgeResourcesConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchEdgeResourcesConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchEdgeResourcesConfiguration`: EdgeResourcesConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchEdgeResourcesConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**edgeResourcesConfigurationId** | **string** | The resource identifier of an existing Edge Resources Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEdgeResourcesConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **edgeResourcesConfiguration** | [**EdgeResourcesConfiguration**](EdgeResourcesConfiguration.md) | A JSON representation of a Edge Resources Configuration | 

### Return type

[**EdgeResourcesConfiguration**](EdgeResourcesConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeResourcesConfiguration

> EdgeResourcesConfiguration RetrieveEdgeResourcesConfiguration(ctx, provisioningSessionId, edgeResourcesConfigurationId).Execute()

Retrieve the Edge Resources Configuration of the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EdgeResourcesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    edgeResourcesConfigurationId := "edgeResourcesConfigurationId_example" // string | The resource identifier of an existing Edge Resources Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveEdgeResourcesConfiguration(context.Background(), provisioningSessionId, edgeResourcesConfigurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveEdgeResourcesConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveEdgeResourcesConfiguration`: EdgeResourcesConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveEdgeResourcesConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**edgeResourcesConfigurationId** | **string** | The resource identifier of an existing Edge Resources Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeResourcesConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EdgeResourcesConfiguration**](EdgeResourcesConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeResourcesConfiguration

> UpdateEdgeResourcesConfiguration(ctx, provisioningSessionId, edgeResourcesConfigurationId).EdgeResourcesConfiguration(edgeResourcesConfiguration).Execute()

Update an Edge Resources Configuration for the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_EdgeResourcesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    edgeResourcesConfigurationId := "edgeResourcesConfigurationId_example" // string | The resource identifier of an existing Edge Resources Configuration.
    edgeResourcesConfiguration := *openapiclient.NewEdgeResourcesConfiguration("EdgeResourcesConfigurationId_example", *openapiclient.NewEdgeManagementMode(), *openapiclient.NewEASRequirements([]string{"EasProviderIds_example"}, "EasType_example", []string{"EasFeatures_example"}, []openapiclient.ScheduledCommunicationTime{*openapiclient.NewScheduledCommunicationTime()}, []openapiclient.ACRScenario{*openapiclient.NewACRScenario()})) // EdgeResourcesConfiguration | A JSON representation of an Edge Resources Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateEdgeResourcesConfiguration(context.Background(), provisioningSessionId, edgeResourcesConfigurationId).EdgeResourcesConfiguration(edgeResourcesConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEdgeResourcesConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**edgeResourcesConfigurationId** | **string** | The resource identifier of an existing Edge Resources Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeResourcesConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **edgeResourcesConfiguration** | [**EdgeResourcesConfiguration**](EdgeResourcesConfiguration.md) | A JSON representation of an Edge Resources Configuration | 

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

