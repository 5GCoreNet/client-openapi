# \DefaultApi

All URIs are relative to *https://example.com/3gpp-ndcaf_data-reporting-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](DefaultApi.md#CreateConfiguration) | **Post** /sessions/{sessionId}/configurations/{configurationId} | Create a new Data Reporting Configuration subresource within the scope of an existing Data Reporting Provisioning Session
[**CreateSession**](DefaultApi.md#CreateSession) | **Post** /sessions | Create a new Data Reporting Provisioning Session
[**DestroyConfiguration**](DefaultApi.md#DestroyConfiguration) | **Delete** /sessions/{sessionId}/configurations/{configurationId} | Destroy an existing Data Reporting Configuration
[**DestroySession**](DefaultApi.md#DestroySession) | **Delete** /sessions/{sessionId} | Destroy an existing Data Reporting Provisioning Session
[**ModifyConfiguration**](DefaultApi.md#ModifyConfiguration) | **Patch** /sessions/{sessionId}/configurations/{configurationId} | Modify an existing Data Reporting Configuration subresource
[**RetrieveConfiguration**](DefaultApi.md#RetrieveConfiguration) | **Get** /sessions/{sessionId}/configurations/{configurationId} | Retrieve an existing Data Reporting Configuration
[**RetrieveSession**](DefaultApi.md#RetrieveSession) | **Get** /sessions/{sessionId} | Retrieve an existing Data Reporting Provisioning Session
[**UpdateConfiguration**](DefaultApi.md#UpdateConfiguration) | **Put** /sessions/{sessionId}/configurations/{configurationId} | Replace an existing Data Reporting Configuration subresource



## CreateConfiguration

> DataReportingConfiguration CreateConfiguration(ctx, sessionId, configurationId).DataReportingConfiguration(dataReportingConfiguration).Execute()

Create a new Data Reporting Configuration subresource within the scope of an existing Data Reporting Provisioning Session

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
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Provisioning Session.
    configurationId := "configurationId_example" // string | The resource identifier of an existing Data Reporting Configuration.
    dataReportingConfiguration := *openapiclient.NewDataReportingConfiguration("DataReportingConfigurationId_example", *openapiclient.NewDataCollectionClientType(), []openapiclient.DataAccessProfile{*openapiclient.NewDataAccessProfile("DataAccessProfileId_example", []openapiclient.EventConsumerType{*openapiclient.NewEventConsumerType()}, []string{"Parameters_example"})}) // DataReportingConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateConfiguration(context.Background(), sessionId, configurationId).DataReportingConfiguration(dataReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfiguration`: DataReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The resource identifier of an existing Data Reporting Provisioning Session. | 
**configurationId** | **string** | The resource identifier of an existing Data Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataReportingConfiguration** | [**DataReportingConfiguration**](DataReportingConfiguration.md) |  | 

### Return type

[**DataReportingConfiguration**](DataReportingConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> DataReportingProvisioningSession CreateSession(ctx).DataReportingProvisioningSession(dataReportingProvisioningSession).Execute()

Create a new Data Reporting Provisioning Session

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
    dataReportingProvisioningSession := *openapiclient.NewDataReportingProvisioningSession("ProvisioningSessionId_example", "AspId_example", "ExternalApplicationId_example", *openapiclient.NewAfEvent(), []string{"DataReportingConfigurationIds_example"}) // DataReportingProvisioningSession | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSession(context.Background()).DataReportingProvisioningSession(dataReportingProvisioningSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSession`: DataReportingProvisioningSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataReportingProvisioningSession** | [**DataReportingProvisioningSession**](DataReportingProvisioningSession.md) |  | 

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


## DestroyConfiguration

> DestroyConfiguration(ctx, sessionId, configurationId).Execute()

Destroy an existing Data Reporting Configuration

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
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Provisioning Session.
    configurationId := "configurationId_example" // string | The resource identifier of an existing Data Reporting Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyConfiguration(context.Background(), sessionId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The resource identifier of an existing Data Reporting Provisioning Session. | 
**configurationId** | **string** | The resource identifier of an existing Data Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyConfigurationRequest struct via the builder pattern


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


## DestroySession

> DestroySession(ctx, sessionId).Execute()

Destroy an existing Data Reporting Provisioning Session

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
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Provisioning Session.

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
**sessionId** | **string** | The resource identifier of an existing Data Reporting Provisioning Session. | 

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


## ModifyConfiguration

> DataReportingConfiguration ModifyConfiguration(ctx, sessionId, configurationId).DataReportingConfigurationPatch(dataReportingConfigurationPatch).Execute()

Modify an existing Data Reporting Configuration subresource

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
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Provisioning Session.
    configurationId := "configurationId_example" // string | The resource identifier of an existing Data Reporting Configuration.
    dataReportingConfigurationPatch := *openapiclient.NewDataReportingConfigurationPatch() // DataReportingConfigurationPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ModifyConfiguration(context.Background(), sessionId, configurationId).DataReportingConfigurationPatch(dataReportingConfigurationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyConfiguration`: DataReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The resource identifier of an existing Data Reporting Provisioning Session. | 
**configurationId** | **string** | The resource identifier of an existing Data Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataReportingConfigurationPatch** | [**DataReportingConfigurationPatch**](DataReportingConfigurationPatch.md) |  | 

### Return type

[**DataReportingConfiguration**](DataReportingConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveConfiguration

> DataReportingConfiguration RetrieveConfiguration(ctx, sessionId, configurationId).Execute()

Retrieve an existing Data Reporting Configuration

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
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Provisioning Session.
    configurationId := "configurationId_example" // string | The resource identifier of an existing Data Reporting Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveConfiguration(context.Background(), sessionId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveConfiguration`: DataReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The resource identifier of an existing Data Reporting Provisioning Session. | 
**configurationId** | **string** | The resource identifier of an existing Data Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataReportingConfiguration**](DataReportingConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSession

> DataReportingProvisioningSession RetrieveSession(ctx, sessionId).Execute()

Retrieve an existing Data Reporting Provisioning Session

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
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Provisioning Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSession`: DataReportingProvisioningSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The resource identifier of an existing Data Reporting Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataReportingProvisioningSession**](DataReportingProvisioningSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> DataReportingConfiguration UpdateConfiguration(ctx, sessionId, configurationId).DataReportingConfiguration(dataReportingConfiguration).Execute()

Replace an existing Data Reporting Configuration subresource

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
    sessionId := "sessionId_example" // string | The resource identifier of an existing Data Reporting Provisioning Session.
    configurationId := "configurationId_example" // string | The resource identifier of an existing Data Reporting Configuration.
    dataReportingConfiguration := *openapiclient.NewDataReportingConfiguration("DataReportingConfigurationId_example", *openapiclient.NewDataCollectionClientType(), []openapiclient.DataAccessProfile{*openapiclient.NewDataAccessProfile("DataAccessProfileId_example", []openapiclient.EventConsumerType{*openapiclient.NewEventConsumerType()}, []string{"Parameters_example"})}) // DataReportingConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateConfiguration(context.Background(), sessionId, configurationId).DataReportingConfiguration(dataReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: DataReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The resource identifier of an existing Data Reporting Provisioning Session. | 
**configurationId** | **string** | The resource identifier of an existing Data Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataReportingConfiguration** | [**DataReportingConfiguration**](DataReportingConfiguration.md) |  | 

### Return type

[**DataReportingConfiguration**](DataReportingConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

