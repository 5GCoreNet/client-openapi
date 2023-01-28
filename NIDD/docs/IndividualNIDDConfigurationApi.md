# \IndividualNIDDConfigurationApi

All URIs are relative to *https://example.com/3gpp-nidd/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNIDDConfiguration**](IndividualNIDDConfigurationApi.md#DeleteNIDDConfiguration) | **Delete** /{scsAsId}/configurations/{configurationId} | Delete an existing NIDD configuration resource.
[**FetchIndNIDDConfiguration**](IndividualNIDDConfigurationApi.md#FetchIndNIDDConfiguration) | **Get** /{scsAsId}/configurations/{configurationId} | Read an NIDD configuration resource.
[**ModifyNIDDConfiguration**](IndividualNIDDConfigurationApi.md#ModifyNIDDConfiguration) | **Patch** /{scsAsId}/configurations/{configurationId} | Modify an existing NIDD configuration resource.



## DeleteNIDDConfiguration

> DeleteNIDDConfiguration(ctx, scsAsId, configurationId).Execute()

Delete an existing NIDD configuration resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/NIDD"
)

func main() {
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    configurationId := "configurationId_example" // string | String identifying the individual NIDD configuration resource in the SCEF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNIDDConfigurationApi.DeleteNIDDConfiguration(context.Background(), scsAsId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNIDDConfigurationApi.DeleteNIDDConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNIDDConfigurationRequest struct via the builder pattern


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


## FetchIndNIDDConfiguration

> NiddConfiguration FetchIndNIDDConfiguration(ctx, scsAsId, configurationId).Execute()

Read an NIDD configuration resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/NIDD"
)

func main() {
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    configurationId := "configurationId_example" // string | String identifying the individual NIDD configuration resource in the SCEF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNIDDConfigurationApi.FetchIndNIDDConfiguration(context.Background(), scsAsId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNIDDConfigurationApi.FetchIndNIDDConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndNIDDConfiguration`: NiddConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualNIDDConfigurationApi.FetchIndNIDDConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndNIDDConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NiddConfiguration**](NiddConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyNIDDConfiguration

> NiddConfiguration ModifyNIDDConfiguration(ctx, scsAsId, configurationId).NiddConfigurationPatch(niddConfigurationPatch).Execute()

Modify an existing NIDD configuration resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/NIDD"
)

func main() {
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    configurationId := "configurationId_example" // string | String identifying the individual NIDD configuration resource in the SCEF.
    niddConfigurationPatch := *openapiclient.NewNiddConfigurationPatch() // NiddConfigurationPatch | Contains information to be applied to the individual NIDD configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNIDDConfigurationApi.ModifyNIDDConfiguration(context.Background(), scsAsId, configurationId).NiddConfigurationPatch(niddConfigurationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNIDDConfigurationApi.ModifyNIDDConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyNIDDConfiguration`: NiddConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualNIDDConfigurationApi.ModifyNIDDConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyNIDDConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **niddConfigurationPatch** | [**NiddConfigurationPatch**](NiddConfigurationPatch.md) | Contains information to be applied to the individual NIDD configuration. | 

### Return type

[**NiddConfiguration**](NiddConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

