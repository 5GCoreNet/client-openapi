# \IndividualDataReportingConfigurationApi

All URIs are relative to *https://example.com/3gpp-data-reporting-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndDataRepConfig**](IndividualDataReportingConfigurationApi.md#DeleteIndDataRepConfig) | **Delete** /sessions/{sessionId}/configurations/{configurationId} | Deletes an already existing Data Reporting Configuration resource.
[**GetIndDataRepConfig**](IndividualDataReportingConfigurationApi.md#GetIndDataRepConfig) | **Get** /sessions/{sessionId}/configurations/{configurationId} | Request the retrieval of an existing Individual Data Reporting Configuration resource.
[**ModifyIndDataRepConfig**](IndividualDataReportingConfigurationApi.md#ModifyIndDataRepConfig) | **Patch** /sessions/{sessionId}/configurations/{configurationId} | Request to modify an existing Individual Data Reporting Configuration resource.
[**UpdateIndDataRepConfig**](IndividualDataReportingConfigurationApi.md#UpdateIndDataRepConfig) | **Put** /sessions/{sessionId}/configurations/{configurationId} | Request to update an existing Individual Data Reporting Configuration resource.



## DeleteIndDataRepConfig

> DeleteIndDataRepConfig(ctx, sessionId, configurationId).Execute()

Deletes an already existing Data Reporting Configuration resource.

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
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Provisioning Session.
    configurationId := "configurationId_example" // string | Identifier of the Data Reporting Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingConfigurationApi.DeleteIndDataRepConfig(context.Background(), sessionId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingConfigurationApi.DeleteIndDataRepConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Provisioning Session. | 
**configurationId** | **string** | Identifier of the Data Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndDataRepConfigRequest struct via the builder pattern


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


## GetIndDataRepConfig

> DataReportingConfiguration GetIndDataRepConfig(ctx, sessionId, configurationId).Execute()

Request the retrieval of an existing Individual Data Reporting Configuration resource.

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
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Provisioning Session.
    configurationId := "configurationId_example" // string | Identifier of the Data Reporting Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingConfigurationApi.GetIndDataRepConfig(context.Background(), sessionId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingConfigurationApi.GetIndDataRepConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndDataRepConfig`: DataReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualDataReportingConfigurationApi.GetIndDataRepConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Provisioning Session. | 
**configurationId** | **string** | Identifier of the Data Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndDataRepConfigRequest struct via the builder pattern


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


## ModifyIndDataRepConfig

> DataReportingConfiguration ModifyIndDataRepConfig(ctx, sessionId, configurationId).DataReportingConfigurationPatch(dataReportingConfigurationPatch).Execute()

Request to modify an existing Individual Data Reporting Configuration resource.

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
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Provisioning Session.
    configurationId := "configurationId_example" // string | Identifier of the Data Reporting Configuration.
    dataReportingConfigurationPatch := *openapiclient.NewDataReportingConfigurationPatch() // DataReportingConfigurationPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingConfigurationApi.ModifyIndDataRepConfig(context.Background(), sessionId, configurationId).DataReportingConfigurationPatch(dataReportingConfigurationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingConfigurationApi.ModifyIndDataRepConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndDataRepConfig`: DataReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualDataReportingConfigurationApi.ModifyIndDataRepConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Provisioning Session. | 
**configurationId** | **string** | Identifier of the Data Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndDataRepConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataReportingConfigurationPatch** | [**DataReportingConfigurationPatch**](DataReportingConfigurationPatch.md) |  | 

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


## UpdateIndDataRepConfig

> DataReportingConfiguration UpdateIndDataRepConfig(ctx, sessionId, configurationId).DataReportingConfiguration(dataReportingConfiguration).Execute()

Request to update an existing Individual Data Reporting Configuration resource.

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
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Provisioning Session.
    configurationId := "configurationId_example" // string | Identifier of the Data Reporting Configuration.
    dataReportingConfiguration := *openapiclient.NewDataReportingConfiguration("DataReportingConfigurationId_example", *openapiclient.NewDataCollectionClientType(), []openapiclient.DataAccessProfile{*openapiclient.NewDataAccessProfile("DataAccessProfileId_example", []openapiclient.EventConsumerType{*openapiclient.NewEventConsumerType()}, []string{"Parameters_example"})}) // DataReportingConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingConfigurationApi.UpdateIndDataRepConfig(context.Background(), sessionId, configurationId).DataReportingConfiguration(dataReportingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingConfigurationApi.UpdateIndDataRepConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndDataRepConfig`: DataReportingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualDataReportingConfigurationApi.UpdateIndDataRepConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Provisioning Session. | 
**configurationId** | **string** | Identifier of the Data Reporting Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndDataRepConfigRequest struct via the builder pattern


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

