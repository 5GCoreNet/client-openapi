# \NpConfigurationsApi

All URIs are relative to *https://example.com/3gpp-network-parameter-configuration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNPConfiguration**](NpConfigurationsApi.md#CreateNPConfiguration) | **Post** /{scsAsId}/configurations | Creates a new configuration resource for network parameter configuration.
[**FetchAllNPConfigurations**](NpConfigurationsApi.md#FetchAllNPConfigurations) | **Get** /{scsAsId}/configurations | Read all of the active configurations for the SCS/AS.



## CreateNPConfiguration

> NpConfiguration CreateNPConfiguration(ctx, scsAsId).NpConfiguration(npConfiguration).Execute()

Creates a new configuration resource for network parameter configuration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/NpConfiguration"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    npConfiguration := openapiclient.NpConfiguration{Interface{}: new(interface{})} // NpConfiguration | new configuration creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NpConfigurationsApi.CreateNPConfiguration(context.Background(), scsAsId).NpConfiguration(npConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NpConfigurationsApi.CreateNPConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNPConfiguration`: NpConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NpConfigurationsApi.CreateNPConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNPConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **npConfiguration** | [**NpConfiguration**](NpConfiguration.md) | new configuration creation | 

### Return type

[**NpConfiguration**](NpConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllNPConfigurations

> []NpConfiguration FetchAllNPConfigurations(ctx, scsAsId).Execute()

Read all of the active configurations for the SCS/AS.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/NpConfiguration"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NpConfigurationsApi.FetchAllNPConfigurations(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NpConfigurationsApi.FetchAllNPConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllNPConfigurations`: []NpConfiguration
    fmt.Fprintf(os.Stdout, "Response from `NpConfigurationsApi.FetchAllNPConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllNPConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NpConfiguration**](NpConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

