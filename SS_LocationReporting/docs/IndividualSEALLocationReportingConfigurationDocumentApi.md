# \IndividualSEALLocationReportingConfigurationDocumentApi

All URIs are relative to *https://example.com/ss-lr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLocReportingConfig**](IndividualSEALLocationReportingConfigurationDocumentApi.md#DeleteLocReportingConfig) | **Delete** /trigger-configurations/{configurationId} | 
[**ModifyLocReportingConfig**](IndividualSEALLocationReportingConfigurationDocumentApi.md#ModifyLocReportingConfig) | **Patch** /trigger-configurations/{configurationId} | 
[**RetrieveLocReportingConfig**](IndividualSEALLocationReportingConfigurationDocumentApi.md#RetrieveLocReportingConfig) | **Get** /trigger-configurations/{configurationId} | 
[**UpdateLocReportingConfig**](IndividualSEALLocationReportingConfigurationDocumentApi.md#UpdateLocReportingConfig) | **Put** /trigger-configurations/{configurationId} | 



## DeleteLocReportingConfig

> DeleteLocReportingConfig(ctx, configurationId).Execute()





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
    configurationId := "configurationId_example" // string | String identifying an individual configuration resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSEALLocationReportingConfigurationDocumentApi.DeleteLocReportingConfig(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSEALLocationReportingConfigurationDocumentApi.DeleteLocReportingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | String identifying an individual configuration resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocReportingConfigRequest struct via the builder pattern


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


## ModifyLocReportingConfig

> LocationReportConfiguration ModifyLocReportingConfig(ctx, configurationId).LocationReportConfigurationPatch(locationReportConfigurationPatch).Execute()





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
    configurationId := "configurationId_example" // string | Identifier of an individual SEAL location reporting configuration.
    locationReportConfigurationPatch := *openapiclient.NewLocationReportConfigurationPatch() // LocationReportConfigurationPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSEALLocationReportingConfigurationDocumentApi.ModifyLocReportingConfig(context.Background(), configurationId).LocationReportConfigurationPatch(locationReportConfigurationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSEALLocationReportingConfigurationDocumentApi.ModifyLocReportingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyLocReportingConfig`: LocationReportConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualSEALLocationReportingConfigurationDocumentApi.ModifyLocReportingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | Identifier of an individual SEAL location reporting configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLocReportingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationReportConfigurationPatch** | [**LocationReportConfigurationPatch**](LocationReportConfigurationPatch.md) |  | 

### Return type

[**LocationReportConfiguration**](LocationReportConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveLocReportingConfig

> LocationReportConfiguration RetrieveLocReportingConfig(ctx, configurationId).Execute()





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
    configurationId := "configurationId_example" // string | String identifying an individual configuration resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSEALLocationReportingConfigurationDocumentApi.RetrieveLocReportingConfig(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSEALLocationReportingConfigurationDocumentApi.RetrieveLocReportingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveLocReportingConfig`: LocationReportConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualSEALLocationReportingConfigurationDocumentApi.RetrieveLocReportingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | String identifying an individual configuration resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveLocReportingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocationReportConfiguration**](LocationReportConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocReportingConfig

> LocationReportConfiguration UpdateLocReportingConfig(ctx, configurationId).LocationReportConfiguration(locationReportConfiguration).Execute()





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
    configurationId := "configurationId_example" // string | String identifying an individual configuration resource.
    locationReportConfiguration := *openapiclient.NewLocationReportConfiguration("ValServerId_example", openapiclient.ValTargetUe{Interface{}: new(interface{})}) // LocationReportConfiguration | Configuration information to be updated in location management server.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSEALLocationReportingConfigurationDocumentApi.UpdateLocReportingConfig(context.Background(), configurationId).LocationReportConfiguration(locationReportConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSEALLocationReportingConfigurationDocumentApi.UpdateLocReportingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLocReportingConfig`: LocationReportConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualSEALLocationReportingConfigurationDocumentApi.UpdateLocReportingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | String identifying an individual configuration resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocReportingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationReportConfiguration** | [**LocationReportConfiguration**](LocationReportConfiguration.md) | Configuration information to be updated in location management server. | 

### Return type

[**LocationReportConfiguration**](LocationReportConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

