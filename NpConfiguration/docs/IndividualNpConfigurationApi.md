# \IndividualNpConfigurationApi

All URIs are relative to *https://example.com/3gpp-network-parameter-configuration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndNPConfiguration**](IndividualNpConfigurationApi.md#DeleteIndNPConfiguration) | **Delete** /{scsAsId}/configurations/{configurationId} | Deletes an already existing configuration.
[**FetchIndNPConfiguration**](IndividualNpConfigurationApi.md#FetchIndNPConfiguration) | **Get** /{scsAsId}/configurations/{configurationId} | Read an active configuration for the SCS/AS and the configuration Id.
[**ModifyIndNPConfiguration**](IndividualNpConfigurationApi.md#ModifyIndNPConfiguration) | **Patch** /{scsAsId}/configurations/{configurationId} | Updates/replaces an existing configuration resource.
[**UpdateIndNPConfiguration**](IndividualNpConfigurationApi.md#UpdateIndNPConfiguration) | **Put** /{scsAsId}/configurations/{configurationId} | Updates/replaces an existing configuration resource.



## DeleteIndNPConfiguration

> []ConfigResult DeleteIndNPConfiguration(ctx, scsAsId, configurationId).Execute()

Deletes an already existing configuration.

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
    configurationId := "configurationId_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNpConfigurationApi.DeleteIndNPConfiguration(context.Background(), scsAsId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNpConfigurationApi.DeleteIndNPConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIndNPConfiguration`: []ConfigResult
    fmt.Fprintf(os.Stdout, "Response from `IndividualNpConfigurationApi.DeleteIndNPConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndNPConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ConfigResult**](ConfigResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIndNPConfiguration

> NpConfiguration FetchIndNPConfiguration(ctx, scsAsId, configurationId).Execute()

Read an active configuration for the SCS/AS and the configuration Id.

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
    configurationId := "configurationId_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNpConfigurationApi.FetchIndNPConfiguration(context.Background(), scsAsId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNpConfigurationApi.FetchIndNPConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndNPConfiguration`: NpConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualNpConfigurationApi.FetchIndNPConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndNPConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NpConfiguration**](NpConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndNPConfiguration

> NpConfiguration ModifyIndNPConfiguration(ctx, scsAsId, configurationId).NpConfigurationPatch(npConfigurationPatch).Execute()

Updates/replaces an existing configuration resource.

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
    configurationId := "configurationId_example" // string | Identifier of the configuration resource
    npConfigurationPatch := *openapiclient.NewNpConfigurationPatch() // NpConfigurationPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNpConfigurationApi.ModifyIndNPConfiguration(context.Background(), scsAsId, configurationId).NpConfigurationPatch(npConfigurationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNpConfigurationApi.ModifyIndNPConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndNPConfiguration`: NpConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualNpConfigurationApi.ModifyIndNPConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndNPConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **npConfigurationPatch** | [**NpConfigurationPatch**](NpConfigurationPatch.md) |  | 

### Return type

[**NpConfiguration**](NpConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndNPConfiguration

> NpConfiguration UpdateIndNPConfiguration(ctx, scsAsId, configurationId).NpConfiguration(npConfiguration).Execute()

Updates/replaces an existing configuration resource.

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
    configurationId := "configurationId_example" // string | Identifier of the configuration resource
    npConfiguration := openapiclient.NpConfiguration{Interface{}: new(interface{})} // NpConfiguration | Parameters to update/replace the existing configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNpConfigurationApi.UpdateIndNPConfiguration(context.Background(), scsAsId, configurationId).NpConfiguration(npConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNpConfigurationApi.UpdateIndNPConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndNPConfiguration`: NpConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualNpConfigurationApi.UpdateIndNPConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndNPConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **npConfiguration** | [**NpConfiguration**](NpConfiguration.md) | Parameters to update/replace the existing configuration | 

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

