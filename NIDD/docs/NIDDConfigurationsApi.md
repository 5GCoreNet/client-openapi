# \NIDDConfigurationsApi

All URIs are relative to *https://example.com/3gpp-nidd/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNIDDConfiguration**](NIDDConfigurationsApi.md#CreateNIDDConfiguration) | **Post** /{scsAsId}/configurations | Create a new NIDD configuration resource.
[**FetchAllNIDDConfigurations**](NIDDConfigurationsApi.md#FetchAllNIDDConfigurations) | **Get** /{scsAsId}/configurations | Read all NIDD configuration resources for a given SCS/AS.



## CreateNIDDConfiguration

> NiddConfiguration CreateNIDDConfiguration(ctx, scsAsId).NiddConfiguration(niddConfiguration).Execute()

Create a new NIDD configuration resource.

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
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    niddConfiguration := openapiclient.NiddConfiguration{Interface{}: new(interface{})} // NiddConfiguration | Contains the data to create a NIDD configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NIDDConfigurationsApi.CreateNIDDConfiguration(context.Background(), scsAsId).NiddConfiguration(niddConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NIDDConfigurationsApi.CreateNIDDConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNIDDConfiguration`: NiddConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NIDDConfigurationsApi.CreateNIDDConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNIDDConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **niddConfiguration** | [**NiddConfiguration**](NiddConfiguration.md) | Contains the data to create a NIDD configuration. | 

### Return type

[**NiddConfiguration**](NiddConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllNIDDConfigurations

> []NiddConfiguration FetchAllNIDDConfigurations(ctx, scsAsId).Execute()

Read all NIDD configuration resources for a given SCS/AS.

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
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NIDDConfigurationsApi.FetchAllNIDDConfigurations(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NIDDConfigurationsApi.FetchAllNIDDConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllNIDDConfigurations`: []NiddConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NIDDConfigurationsApi.FetchAllNIDDConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllNIDDConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NiddConfiguration**](NiddConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

